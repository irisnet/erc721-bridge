package keeper

import (
	"context"
	"math/big"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"

	nfttransferkeeper "github.com/bianjieai/nft-transfer/keeper"
	nfttransfertypes "github.com/bianjieai/nft-transfer/types"

	"github.com/irisnet/erc721-bridge/x/nft-transfer/types"
)

var _ nfttransfertypes.MsgServer = Keeper{}

// Transfer defines a rpc handler method for MsgTransfer.
func (k Keeper) Transfer(goCtx context.Context, msg *nfttransfertypes.MsgTransfer) (*nfttransfertypes.MsgTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// If the classId has an "ibc/" prefix, it only needs to be handed over to nft-transfer for processing,
	// because nft-transfer has a trace to get the fullClassId of the original chain
	if strings.HasPrefix(msg.ClassId, nfttransfertypes.ClassPrefix+"/") {
		return k.ics721Keeper.Transfer(goCtx, msg)
	}

	// If the classId does not contain the "ibc/" prefix and is not a legal contract address,
	// it means that the token is the nft of the native chain and can be directly handed over to nft-transfer for processing
	if !common.IsHexAddress(msg.ClassId) {
		return k.ics721Keeper.Transfer(goCtx, msg)
	}

	contractAddr := common.HexToAddress(msg.ClassId)
	ok := k.converterKeeper.HasContract(ctx, contractAddr)
	// If classId is a contract address, but the contract does not exist,
	// it means that this is just an nft of the chain (classId and contract address have the same format)
	if !ok {
		return k.ics721Keeper.Transfer(goCtx, msg)
	}

	//If there is no mapping between the contract and other nft, it means that this is a local erc721 token
	ibcClassId, ok := k.converterKeeper.ContractToClass(ctx, contractAddr)
	if !ok {
		return k.ics721Keeper.Transfer(ctx, msg)
	}

	var tokenIds []string
	for _, erc721TokenId := range msg.TokenIds {
		erc721TokenIdInt, ok := new(big.Int).SetString(erc721TokenId, 10)
		if !ok {
			return nil, types.ErrInvalidErc721TokenId
		}

		tokenId, ok := k.converterKeeper.ERC721ToNFT(ctx, contractAddr, erc721TokenIdInt)
		if !ok {
			return nil, types.ErrNotExistErc721TokenId
		}
		tokenIds = append(tokenIds, tokenId)
	}
	msg.ClassId = ibcClassId
	msg.TokenIds = tokenIds
	return k.ics721Keeper.Transfer(goCtx, msg)
}

func (k Keeper) ISC721Keeper() nfttransferkeeper.Keeper {
	return k.ics721Keeper
}
