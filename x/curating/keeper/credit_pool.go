package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/public-awesome/stakebird/x/curating/types"
)

// InitializeCreditPool sets up the credit pool from genesis
func (k Keeper) InitializeCreditPool(ctx sdk.Context, funds sdk.Coin) error {
	return k.bankKeeper.MintCoins(ctx, types.CreditPoolName, sdk.NewCoins(funds))
}

// GetCreditPoolBalance returns the curation credit pool balance
func (k Keeper) GetCreditPoolBalance(ctx sdk.Context) sdk.Coin {
	return k.bankKeeper.GetBalance(ctx, k.GetCreditPool(ctx).GetAddress(), k.GetParams(ctx).CreditDenom)
}

// GetCreditPool returns the credit pool account
func (k Keeper) GetCreditPool(ctx sdk.Context) authtypes.ModuleAccountI {
	return k.accountKeeper.GetModuleAccount(ctx, types.CreditPoolName)
}

// DistributeCredits distributes credits to all accounts
func (k Keeper) DistributeCredits(ctx sdk.Context) error {
	// blockInflationAddr := k.accountKeeper.GetModuleAccount(ctx, authtypes.FeeCollectorName).GetAddress()
	// blockInflation := k.bankKeeper.GetBalance(ctx, blockInflationAddr, k.GetParams(ctx).StakeDenom)
	// rewardPoolAllocation := k.GetParams(ctx).RewardPoolAllocation
	// blockInflationDec := sdk.NewDecFromInt(blockInflation.Amount)
	// rewardAmount := blockInflationDec.Mul(rewardPoolAllocation)
	// rewardCoin := sdk.NewCoin(k.GetParams(ctx).StakeDenom, rewardAmount.TruncateInt())
	// return k.bankKeeper.SendCoinsFromModuleToModule(
	// 	ctx, authtypes.FeeCollectorName, types.RewardPoolName, sdk.NewCoins(rewardCoin))

	return nil
}
