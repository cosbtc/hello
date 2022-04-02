package hello_test

import (
	"testing"

	keepertest "github.com/cosbtc/hello/testutil/keeper"
	"github.com/cosbtc/hello/testutil/nullify"
	"github.com/cosbtc/hello/x/hello"
	"github.com/cosbtc/hello/x/hello/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.HelloKeeper(t)
	hello.InitGenesis(ctx, *k, genesisState)
	got := hello.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
