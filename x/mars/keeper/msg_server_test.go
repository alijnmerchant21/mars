package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/alijnmerchant21/mars/testutil/keeper"
	"github.com/alijnmerchant21/mars/x/mars/keeper"
	"github.com/alijnmerchant21/mars/x/mars/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.MarsKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
