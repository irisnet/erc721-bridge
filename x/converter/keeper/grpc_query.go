package keeper

import (
	"context"
	"math/big"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/ethereum/go-ethereum/common"
	etherminttypes "github.com/evmos/ethermint/types"

	"github.com/irisnet/erc721-bridge/x/converter/types"
)

// TokenPairs returns all registered pairs
func (k Keeper) TokenPairs(
	c context.Context,
	req *types.QueryTokenPairsRequest,
) (*types.QueryTokenPairsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	var pairs []types.TokenPair
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixTokenPair)

	pageRes, err := query.Paginate(store, req.Pagination, func(_, value []byte) error {
		var pair types.TokenPair
		if err := k.cdc.Unmarshal(value, &pair); err != nil {
			return err
		}
		pairs = append(pairs, pair)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &types.QueryTokenPairsResponse{
		TokenPairs: pairs,
		Pagination: pageRes,
	}, nil
}

// TokenPair returns a given registered token pair
func (k Keeper) TokenPair(
	c context.Context,
	req *types.QueryTokenPairRequest,
) (*types.QueryTokenPairResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	// check if the token is a hex address, if not, check if it is a valid SDK
	// denom
	if err := etherminttypes.ValidateAddress(req.Token); err != nil {
		if err := sdk.ValidateDenom(req.Token); err != nil {
			return nil, status.Errorf(
				codes.InvalidArgument,
				"invalid format for token %s, should be either hex ('0x...') cosmos denom",
				req.Token,
			)
		}
	}

	id := k.GetTokenPairID(ctx, req.Token)

	if len(id) == 0 {
		return nil, status.Errorf(codes.NotFound, "token pair with token '%s'", req.Token)
	}

	pair, found := k.GetTokenPair(ctx, id)
	if !found {
		return nil, status.Errorf(codes.NotFound, "token pair with token '%s'", req.Token)
	}

	return &types.QueryTokenPairResponse{TokenPair: pair}, nil
}

// TokenTrace returns a cross-chain token trace
func (k Keeper) TokenTrace(
	c context.Context,
	req *types.QueryTokenTraceRequest,
) (*types.QueryTokenTraceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	erc721Keeper := k.ERC721Keeper()

	// If the classId has an "ibc/" prefix, it only needs to be handed over to nft-transfer for processing,
	// because nft-transfer has a trace to get the fullClassId of the original chain
	if strings.HasPrefix(req.ClassId, "ibc/") {
		contractAddr, ok := erc721Keeper.ClassToContract(ctx, req.ClassId)
		if !ok {
			return &types.QueryTokenTraceResponse{},
				status.Errorf(codes.NotFound, "token mapping with class_id '%s'", req.ClassId)
		}

		tokenId, err := erc721Keeper.nftToERC721(ctx, req.ClassId, req.TokenId)
		if err != nil {
			return &types.QueryTokenTraceResponse{},
				status.Errorf(codes.NotFound, "token mapping with token_id '%s'", req.TokenId)
		}
		return &types.QueryTokenTraceResponse{
			ClassId: contractAddr.String(),
			TokenId: tokenId.String(),
		}, nil
	}
	// If the classId does not contain the "ibc/" prefix and is not a legal contract address,
	// it means that the token is the nft of the native chain and can be directly handed over to nft-transfer for processing
	if !common.IsHexAddress(req.ClassId) {
		return &types.QueryTokenTraceResponse{},
			status.Errorf(
				codes.InvalidArgument,
				"class_id '%s' is not valid contract address",
				req.ClassId,
			)
	}

	contractAddr := common.HexToAddress(req.ClassId)
	ok := erc721Keeper.HasContract(ctx, contractAddr)
	// If classId is a contract address, but the contract does not exist,
	// it means that this is just an nft of the chain (classId and contract address have the same format)
	if !ok {
		return &types.QueryTokenTraceResponse{},
			status.Errorf(codes.NotFound, "class_id '%s' mapping is not exist", req.ClassId)
	}

	//If there is no mapping between the contract and other nft, it means that this is a local erc721 token
	ibcClassId, ok := erc721Keeper.ContractToClass(ctx, contractAddr)
	if !ok {
		return &types.QueryTokenTraceResponse{},
			status.Errorf(codes.NotFound, "class_id '%s' mapping is not exist", req.ClassId)
	}

	erc721TokenId, ok := new(big.Int).SetString(req.TokenId, 10)
	if !ok {
		return &types.QueryTokenTraceResponse{},
			status.Errorf(
				codes.InvalidArgument,
				"token_id '%s' is not valid erc721 token_id",
				req.TokenId,
			)
	}

	tokenId, ok := erc721Keeper.ERC721ToNFT(ctx, contractAddr, erc721TokenId)
	if !ok {
		return &types.QueryTokenTraceResponse{},
			status.Errorf(codes.NotFound, "token_id '%s' mapping is not exist", req.TokenId)
	}
	return &types.QueryTokenTraceResponse{
		ClassId: ibcClassId,
		TokenId: tokenId,
	}, nil
}
