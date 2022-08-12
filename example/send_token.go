package example

import (
	// "github.com/cosmos/cosmos-sdk/simapp"

	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"

	chainclient "github.com/gotabit/gotabit/client/chain"
)

func NewTxFactory(clientCtx client.Context) tx.Factory {
	return new(tx.Factory).
		WithKeybase(clientCtx.Keyring).
		WithTxConfig(clientCtx.TxConfig).
		WithAccountRetriever(clientCtx.AccountRetriever).
		WithSimulateAndExecute(true).
		WithGasAdjustment(1.5).
		WithChainID(clientCtx.ChainID).
		WithSignMode(signing.SignMode_SIGN_MODE_DIRECT)
}

func main() {
	tmRPC, err := rpchttp.New("https://rpc-testnet.gotabit.dev:443", "")
	if err != nil {
		panic(err)
	}

	kb, err := keyring.New("gotabit", "file", os.Getenv("HOME")+"/.gotabit", os.Stdin)
	if err != nil {
		panic(err)
	}
	kb.NewAccount("sender", "dinner proud piano mention silk plunge forest fold trial duck electric define", "", "m/44'/118'/0'/0/0", hd.Secp256k1)

	list, err := kb.List()
	if err != nil {
		panic(err)
	}

	clientCtx, err := chainclient.NewClientContext(
		"gotabit-test-1",
		list[0].GetAddress().String(),
		kb,
	)

	if err != nil {
		fmt.Println(err)
	}

	clientCtx = clientCtx.WithNodeURI("https://rest-testnet.gotabit.dev:443").WithClient(tmRPC)

	_, _, addr := testdata.KeyTestPubAddr()

	msg := &banktypes.MsgSend{
		FromAddress: list[0].GetAddress().String(),
		ToAddress:   addr.String(),
		Amount: []sdktypes.Coin{{
			Denom:  "ugtb",
			Amount: sdktypes.NewInt(600000000),
		}},
	}

	txFactory := NewTxFactory(clientCtx)

	tx.BroadcastTx(clientCtx, txFactory, msg)
}
