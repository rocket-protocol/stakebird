package keeper

import (
	"strconv"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/public-awesome/stakebird/x/user/types"
)

// NewQuerier creates a new querier for user clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryParams:
			return queryParams(ctx, k)
		case types.QueryPost:
			return queryPost(ctx, path[1:], req, k)
		case types.QueryUpvotes:
			return queryUpvotes(ctx, path[1:], req, k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown user query endpoint")
		}
	}
}

func queryParams(ctx sdk.Context, k Keeper) ([]byte, error) {
	params := k.GetParams(ctx)

	res, err := codec.MarshalJSONIndent(types.ModuleCdc, params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryPost(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) ([]byte, error) {
	i64, err := strconv.ParseUint(path[0], 10, 32)
	if err != nil {
		return nil, err
	}
	vendorID := uint32(i64)
	postID := path[1]

	post, found, err := k.GetPost(ctx, vendorID, postID)
	if err != nil || !found {
		return nil, types.ErrPostNotFound
	}

	res, err := codec.MarshalJSONIndent(types.ModuleCdc, post)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryUpvotes(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) ([]byte, error) {
	i64, err := strconv.ParseUint(path[0], 10, 32)
	if err != nil {
		return nil, err
	}
	vendorID := uint32(i64)
	var upvotes []types.Upvote

	postID := path[1]
	post, found, err := k.GetPost(ctx, vendorID, postID)
	if err != nil || !found {
		return nil, types.ErrPostNotFound
	}
	k.IterateUpvotes(ctx, vendorID, post.PostIDHash, func(upvote types.Upvote) (stop bool) {
		upvotes = append(upvotes, upvote)
		return false
	})

	res, err := codec.MarshalJSONIndent(types.ModuleCdc, upvotes)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}