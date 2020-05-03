// Code generated by sqlc. DO NOT EDIT.

package models

import (
	"context"
)

type Querier interface {
	CreateAsset(ctx context.Context, arg CreateAssetParams) (Asset, error)
	GetAssetsByOwnerID(ctx context.Context, ownerID int32) ([]Asset, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
}

var _ Querier = (*Queries)(nil)
