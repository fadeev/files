package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgDeleteClaim{}

type MsgDeleteClaim struct {
	ID      string
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	Proof   string         `json:"proof" yaml:"proof"`
}

func NewMsgDeleteClaim(creator sdk.AccAddress, proof string) MsgDeleteClaim {
	return MsgDeleteClaim{
		ID:      uuid.New().String(),
		Creator: creator,
		Proof:   proof,
	}
}

func (msg MsgDeleteClaim) Route() string {
	return RouterKey
}

func (msg MsgDeleteClaim) Type() string {
	return "DeleteClaim"
}

func (msg MsgDeleteClaim) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgDeleteClaim) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgDeleteClaim) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
