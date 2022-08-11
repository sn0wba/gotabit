package sendMsg

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gotabit/gotabit/x/sendMsg/keeper"
	"github.com/gotabit/gotabit/x/sendMsg/types"
)

func DefaultGenesisState() *types.GenesisState {
	return &types.GenesisState{}
}

// InitGenesis stores the genesis state
func InitGenesis(ctx sdk.Context, k keeper.Keeper, data types.GenesisState) {
}

// ExportGenesis outputs the genesis state
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{}
}
