package files

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/fadeev/files/x/files/keeper"
	"github.com/fadeev/files/x/files/types"
)

func handleMsgDeleteClaim(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeleteClaim) (*sdk.Result, error) {
	var claim = types.Claim{
		Creator: msg.Creator,
		ID:      msg.ID,
		Proof:   msg.Proof,
	}
	cl, _ := k.GetClaim(ctx, claim)
	if claim.Creator.Equals(cl.Creator) {
		k.DeleteClaim(ctx, claim)
	}
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
