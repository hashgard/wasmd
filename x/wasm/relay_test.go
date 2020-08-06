package wasm_test

import (
	"testing"

	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/CosmWasm/wasmd/x/wasm/ibc_testing"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestFromAToB(t *testing.T) {
	t.Skip("not implemented")
	var (
		coordinator = ibc_testing.NewCoordinator(t, 2)
		chainA      = coordinator.GetChain(ibc_testing.GetChainID(0))
		chainB      = coordinator.GetChain(ibc_testing.GetChainID(1))

		testAddr1, _ = sdk.AccAddressFromBech32("cosmos1scqhwpgsmr6vmztaa7suurfl52my6nd2kmrudl")
	)
	//clientA, clientB, connA, connB, channelA, channelB := coordinator.Setup(chainA, chainB)
	_, _, _, _, channelA, _ := coordinator.Setup(chainB, chainA)

	handler := wasm.NewHandler(chainA.App.WasmKeeper)
	ctx := chainA.GetContext()
	msg := &wasm.MsgWasmIBCCall{
		SourcePort:       channelA.PortID,
		SourceChannel:    channelA.ID,
		Sender:           testAddr1,
		TimeoutHeight:    110,
		TimeoutTimestamp: 0,
		Msg:              []byte("{}"),
	}

	res, err := handler(ctx, msg)
	require.NoError(t, err)
	require.NotNil(t, res)

	//err := coordinator.SendPacket(chainA, chainB, packet, clientB)
	//require.NoError(t, err)
}
