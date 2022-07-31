package main

import (
	"os"

	"github.com/gotabit/gotabit/app"
	"github.com/ignite/cli/ignite/pkg/cosmoscmd"

	sdk "github.com/cosmos/cosmos-sdk/types"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	cmdOptions := GetWasmCmdOptions()

	// set CoinType
	config := sdk.GetConfig()
	config.SetCoinType(app.CoinType)

	rootCmd, _ := cosmoscmd.NewRootCmd(
		app.Name,
		app.AccountAddressPrefix,
		app.DefaultNodeHome,
		app.Name,
		app.ModuleBasics,
		app.New,
		// this line is used by starport scaffolding # root/arguments
		cmdOptions...,
	)
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
