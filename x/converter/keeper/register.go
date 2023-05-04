package keeper

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/irisnet/erc721-bridge/x/converter/contracts"
	"github.com/irisnet/erc721-bridge/x/converter/types"
)

// SaveRegisteredClass saves the registered denom to the store
func (k Keeper) SaveRegisteredClass(ctx sdk.Context, classId string) (common.Address, error) {
	classInfo, found := k.nftKeeper.GetClass(ctx, classId)
	if !found {
		return common.Address{}, errorsmod.Wrapf(
			types.ErrClassNotFound,
			"denom metadata not registered %s", classId,
		)
	}

	// Deployed contract address is used as the key to save the class
	contractAddr, err := k.DeployERC721Contract(ctx,
		classInfo.GetName(), classInfo.GetSymbol(), classInfo.GetURI(), classInfo.GetData())
	if err != nil {
		return common.Address{}, errorsmod.Wrap(
			err, "failed to create wrapped coin denom metadata for ERC721",
		)
	}
	pair := types.NewTokenPair(contractAddr, classInfo.GetID(), types.OWNER_MODULE)
	k.SetTokenPair(ctx, pair)
	k.SetClassMap(ctx, pair.ClassId, pair.GetID())
	k.SetERC721Map(ctx, common.HexToAddress(pair.Erc721Address), pair.GetID())
	return contractAddr, nil
}

// SaveRegisteredERC721 saves the registered ERC721 to the store
func (k Keeper) SaveRegisteredERC721(ctx sdk.Context, contract common.Address) (string, error) {
	classId := types.CreateClass(contract.String())
	if k.nftKeeper.HasClass(ctx, classId) {
		return "", errorsmod.Wrap(
			types.ErrInternalTokenPair,
			"denom metadata already registered",
		)
	}

	if k.IsClassRegistered(ctx, classId) {
		return "", errorsmod.Wrap(
			types.ErrInternalTokenPair,
			"denom metadata already registered",
		)
	}
	erc721Abi := contracts.ERC721PresetMinterPauserContract.ABI

	// Get ERC721 metadata
	erc721MetaData, err := k.QueryERC721(ctx, contract, erc721Abi)
	if err != nil {
		return "", err
	}
	// Create Data
	// Create Native Class
	if err := k.nftKeeper.SaveClass(ctx, classId, erc721MetaData.Name, erc721MetaData.Symbol); err != nil {
		return "", errorsmod.Wrapf(types.ErrSaveClass,
			"failed to save class %s, contract address %s", classId, contract.String())
	}

	pair := types.NewTokenPair(contract, classId, types.OWNER_EXTERNAL)
	k.SetTokenPair(ctx, pair)
	k.SetClassMap(ctx, pair.ClassId, pair.GetID())
	k.SetERC721Map(ctx, common.HexToAddress(pair.Erc721Address), pair.GetID())
	return classId, nil
}
