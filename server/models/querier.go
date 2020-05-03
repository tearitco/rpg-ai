// Code generated by sqlc. DO NOT EDIT.

package models

import (
	"context"
)

type Querier interface {
	CreateAsset(ctx context.Context, arg CreateAssetParams) (Asset, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	ListAssetMetadataByOwnerID(ctx context.Context, ownerID int32) ([]ListAssetMetadataByOwnerIDRow, error)
}

var _ Querier = (*Queries)(nil)
