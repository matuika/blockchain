package v043

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
)


func pruneZeroBalancesJSON(oldBalances []types.Balance) []types.Balance {
	var balances []types.Balance

	for _, b := range oldBalances {
		if !b.Coins.IsZero() {
			b.Coins = sdk.NewCoins(b.Coins...) 
			balances = append(balances, b)
		}
	}

	return balances
}




func MigrateJSON(oldState *types.GenesisState) *types.GenesisState {
	return &types.GenesisState{
		Params:        oldState.Params,
		Balances:      pruneZeroBalancesJSON(oldState.Balances),
		Supply:        sdk.NewCoins(oldState.Supply...), 
		DenomMetadata: oldState.DenomMetadata,
	}
}
