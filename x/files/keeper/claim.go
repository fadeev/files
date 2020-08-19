package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/fadeev/files/x/files/types"
)

func (k Keeper) CreateClaim(ctx sdk.Context, claim types.Claim) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.ClaimPrefix + claim.Proof)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(claim)
	store.Set(key, value)
}

func (k Keeper) DeleteClaim(ctx sdk.Context, claim types.Claim) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.ClaimPrefix + claim.Proof)
	store.Delete(key)
}

func listClaim(ctx sdk.Context, k Keeper) ([]byte, error) {
	var claimList []types.Claim
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.ClaimPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var claim types.Claim
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &claim)
		claimList = append(claimList, claim)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, claimList)
	return res, nil
}

func (k Keeper) GetClaim(ctx sdk.Context, claim types.Claim) (types.Claim, error) {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.ClaimPrefix + claim.Proof)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &claim)
	if err != nil {
		return claim, err
	}
	return claim, nil
}
