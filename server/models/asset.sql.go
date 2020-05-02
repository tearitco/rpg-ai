// Code generated by sqlc. DO NOT EDIT.
// source: asset.sql

package models

import (
	"context"
)

const getAssetsByOwnerID = `-- name: GetAssetsByOwnerID :many
SELECT id, owner_id, content_type, filename, filedata, created_at FROM assets WHERE owner_id = $1
`

func (q *Queries) GetAssetsByOwnerID(ctx context.Context, ownerID int32) ([]Asset, error) {
	rows, err := q.db.QueryContext(ctx, getAssetsByOwnerID, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Asset
	for rows.Next() {
		var i Asset
		if err := rows.Scan(
			&i.ID,
			&i.OwnerID,
			&i.ContentType,
			&i.Filename,
			&i.Filedata,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
