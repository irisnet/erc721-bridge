package keeper

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/irisnet/erc721-bridge/x/converter/contracts"
	"github.com/irisnet/erc721-bridge/x/converter/types"
)

// SaveRegisteredClass saves the registered denom to the store
func (k Keeper) SaveRegisteredClass(ctx sdk.Context, sender sdk.AccAddress, classId string) (common.Address, error) {
	classInfo, found := k.nftKeeper.GetClass(ctx, classId)
	if !found {
		return common.Address{}, errorsmod.Wrapf(
			types.ErrClassPairNotFound,
			"denom metadata not registered %s", classId,
		)
	}
	deployer := common.BytesToAddress(sender.Bytes())

	// Deployed contract address is used as the key to save the class
	contractAddr, err := k.DeployERC721Contract(ctx,
		deployer,
		classInfo.GetName(),
		classInfo.GetSymbol(),
		classInfo.GetURI(),
		classInfo.GetData(),
		types.ModuleAddress)
	if err != nil {
		return common.Address{}, errorsmod.Wrap(
			err, "failed to create wrapped coin denom metadata for ERC721",
		)
	}
	pair := types.NewClassPair(contractAddr, classInfo.GetID(), types.OWNER_MODULE)
	k.SetClassPair(ctx, pair)
	k.SetClassMap(ctx, pair.ClassId, pair.GetID())
	k.SetERC721Map(ctx, common.HexToAddress(pair.ContractAddress), pair.GetID())
	return contractAddr, nil
}

// SaveRegisteredERC721 saves the registered ERC721 to the store
func (k Keeper) SaveRegisteredERC721(ctx sdk.Context, contract common.Address) (string, error) {
	classId := types.CreateClass(contract.String())
	if k.nftKeeper.HasClass(ctx, classId) {
		return "", errorsmod.Wrap(
			types.ErrInternalClassPair,
			"denom metadata already registered",
		)
	}

	if k.IsClassRegistered(ctx, classId) {
		return "", errorsmod.Wrap(
			types.ErrInternalClassPair,
			"denom metadata already registered",
		)
	}
	erc721Abi := contracts.ERC721PresetMinterPauserContract.ABI

	// Get ERC721 metadata
	erc721MetaData, err := k.QueryERC721(ctx, contract, erc721Abi, false)
	if err != nil {
		return "", err
	}
	// Create Data
	// Create Native Class
	if err := k.nftKeeper.SaveClass(ctx, classId, erc721MetaData.URI, erc721MetaData.Data); err != nil {
		return "", errorsmod.Wrapf(types.ErrSaveNativeClass,
			"failed to save class %s, contract address %s", classId, contract.String())
	}

	pair := types.NewClassPair(contract, classId, types.OWNER_EXTERNAL)
	k.SetClassPair(ctx, pair)
	k.SetClassMap(ctx, pair.ClassId, pair.GetID())
	k.SetERC721Map(ctx, common.HexToAddress(pair.ContractAddress), pair.GetID())
	return classId, nil
}
