package dao

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/x/dao/keeper"
	"github.com/public-awesome/stargaze/x/dao/types"
)

// InitGenesis initializes the dao module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, dk types.DistKeeper, genState types.GenesisState) {
	amount := sdk.NewCoins(genState.Params.GenesisAllocation)
	sender, err := sdk.AccAddressFromBech32(genState.Params.GenesisSender)
	if err != nil {
		panic(err)
	}

	err = dk.FundCommunityPool(ctx, amount, sender)
	if err != nil {
		panic(err)
	}

	// this line is used by starport scaffolding # genesis/module/init

	// this line is used by starport scaffolding # ibc/genesis/init
}

// ExportGenesis returns the dao module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export

	// this line is used by starport scaffolding # ibc/genesis/export

	return genesis
}
