// Code generated by sqlc. DO NOT EDIT.
// source: campaigns_encounters_and_characters.sql

package models

import (
	"context"
	"database/sql"
	"encoding/json"
)

const createCampaign = `-- name: CreateCampaign :one
INSERT INTO campaigns (owner_id, name, description) VALUES ($1, $2, $3) RETURNING id, owner_id, name, description, created_at
`

type CreateCampaignParams struct {
	OwnerID     int32
	Name        string
	Description sql.NullString
}

func (q *Queries) CreateCampaign(ctx context.Context, arg CreateCampaignParams) (Campaign, error) {
	row := q.db.QueryRowContext(ctx, createCampaign, arg.OwnerID, arg.Name, arg.Description)
	var i Campaign
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}

const createCharacter = `-- name: CreateCharacter :one
INSERT INTO characters (owner_id, name, definition) VALUES ($1, $2, $3) RETURNING id, owner_id, name, definition, sprite, created_at
`

type CreateCharacterParams struct {
	OwnerID    int32
	Name       string
	Definition json.RawMessage
}

func (q *Queries) CreateCharacter(ctx context.Context, arg CreateCharacterParams) (Character, error) {
	row := q.db.QueryRowContext(ctx, createCharacter, arg.OwnerID, arg.Name, arg.Definition)
	var i Character
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.Name,
		&i.Definition,
		&i.Sprite,
		&i.CreatedAt,
	)
	return i, err
}

const createEncounter = `-- name: CreateEncounter :one
INSERT INTO encounters (campaign_id, name)
SELECT $1 AS campaign_id, $2 AS name
FROM campaigns
WHERE EXISTS (SELECT id FROM campaigns WHERE id = $1 AND campaigns.owner_id = $3)
RETURNING id, campaign_id, name, description, tilemap_id, created_at
`

type CreateEncounterParams struct {
	CampaignID int32
	Name       string
	OwnerID    int32
}

func (q *Queries) CreateEncounter(ctx context.Context, arg CreateEncounterParams) (Encounter, error) {
	row := q.db.QueryRowContext(ctx, createEncounter, arg.CampaignID, arg.Name, arg.OwnerID)
	var i Encounter
	err := row.Scan(
		&i.ID,
		&i.CampaignID,
		&i.Name,
		&i.Description,
		&i.TilemapID,
		&i.CreatedAt,
	)
	return i, err
}

const deleteCampaign = `-- name: DeleteCampaign :exec
DELETE FROM campaigns WHERE id = $1 AND owner_id = $2
`

type DeleteCampaignParams struct {
	ID      int32
	OwnerID int32
}

func (q *Queries) DeleteCampaign(ctx context.Context, arg DeleteCampaignParams) error {
	_, err := q.db.ExecContext(ctx, deleteCampaign, arg.ID, arg.OwnerID)
	return err
}

const deleteCharacter = `-- name: DeleteCharacter :exec
DELETE FROM characters WHERE id = $1 AND owner_id = $2
`

type DeleteCharacterParams struct {
	ID      int32
	OwnerID int32
}

func (q *Queries) DeleteCharacter(ctx context.Context, arg DeleteCharacterParams) error {
	_, err := q.db.ExecContext(ctx, deleteCharacter, arg.ID, arg.OwnerID)
	return err
}

const deleteEncounter = `-- name: DeleteEncounter :exec
DELETE FROM encounters
WHERE encounters.id = $1 AND
EXISTS (SELECT id FROM campaigns WHERE campaigns.id = encounters.campaign_id AND owner_id = $2)
`

type DeleteEncounterParams struct {
	ID      int32
	OwnerID int32
}

func (q *Queries) DeleteEncounter(ctx context.Context, arg DeleteEncounterParams) error {
	_, err := q.db.ExecContext(ctx, deleteEncounter, arg.ID, arg.OwnerID)
	return err
}

const getOwnedCampaignByID = `-- name: GetOwnedCampaignByID :one
SELECT id, owner_id, name, description, created_at FROM campaigns WHERE id = $1 AND owner_id = $2
`

type GetOwnedCampaignByIDParams struct {
	ID      int32
	OwnerID int32
}

func (q *Queries) GetOwnedCampaignByID(ctx context.Context, arg GetOwnedCampaignByIDParams) (Campaign, error) {
	row := q.db.QueryRowContext(ctx, getOwnedCampaignByID, arg.ID, arg.OwnerID)
	var i Campaign
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}

const listCampaignsByOwnerID = `-- name: ListCampaignsByOwnerID :many
SELECT id, owner_id, name, description, created_at FROM campaigns WHERE owner_id = $1
`

func (q *Queries) ListCampaignsByOwnerID(ctx context.Context, ownerID int32) ([]Campaign, error) {
	rows, err := q.db.QueryContext(ctx, listCampaignsByOwnerID, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Campaign
	for rows.Next() {
		var i Campaign
		if err := rows.Scan(
			&i.ID,
			&i.OwnerID,
			&i.Name,
			&i.Description,
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

const listCharactersByOwnerID = `-- name: ListCharactersByOwnerID :many
SELECT id, owner_id, name, definition, sprite, created_at FROM characters WHERE owner_id = $1
`

func (q *Queries) ListCharactersByOwnerID(ctx context.Context, ownerID int32) ([]Character, error) {
	rows, err := q.db.QueryContext(ctx, listCharactersByOwnerID, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Character
	for rows.Next() {
		var i Character
		if err := rows.Scan(
			&i.ID,
			&i.OwnerID,
			&i.Name,
			&i.Definition,
			&i.Sprite,
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

const listCharactersForEncounter = `-- name: ListCharactersForEncounter :many
SELECT characters.id, characters.owner_id, characters.name, characters.definition, characters.sprite, characters.created_at FROM encounter_characters
JOIN characters ON characters.id = character_id
WHERE encounter_id = $1
`

func (q *Queries) ListCharactersForEncounter(ctx context.Context, encounterID int32) ([]Character, error) {
	rows, err := q.db.QueryContext(ctx, listCharactersForEncounter, encounterID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Character
	for rows.Next() {
		var i Character
		if err := rows.Scan(
			&i.ID,
			&i.OwnerID,
			&i.Name,
			&i.Definition,
			&i.Sprite,
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

const listEncountersForCampaign = `-- name: ListEncountersForCampaign :many
SELECT id, campaign_id, name, description, tilemap_id, created_at FROM encounters WHERE campaign_id = $1
`

func (q *Queries) ListEncountersForCampaign(ctx context.Context, campaignID int32) ([]Encounter, error) {
	rows, err := q.db.QueryContext(ctx, listEncountersForCampaign, campaignID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Encounter
	for rows.Next() {
		var i Encounter
		if err := rows.Scan(
			&i.ID,
			&i.CampaignID,
			&i.Name,
			&i.Description,
			&i.TilemapID,
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

const updateCampaignDescription = `-- name: UpdateCampaignDescription :exec
UPDATE campaigns SET description = $3 WHERE id = $1 AND owner_id = $2
`

type UpdateCampaignDescriptionParams struct {
	ID          int32
	OwnerID     int32
	Description sql.NullString
}

func (q *Queries) UpdateCampaignDescription(ctx context.Context, arg UpdateCampaignDescriptionParams) error {
	_, err := q.db.ExecContext(ctx, updateCampaignDescription, arg.ID, arg.OwnerID, arg.Description)
	return err
}

const updateCampaignName = `-- name: UpdateCampaignName :exec
UPDATE campaigns SET name = $3 WHERE id = $1 AND owner_id = $2
`

type UpdateCampaignNameParams struct {
	ID      int32
	OwnerID int32
	Name    string
}

func (q *Queries) UpdateCampaignName(ctx context.Context, arg UpdateCampaignNameParams) error {
	_, err := q.db.ExecContext(ctx, updateCampaignName, arg.ID, arg.OwnerID, arg.Name)
	return err
}

const updateCharacterDefinition = `-- name: UpdateCharacterDefinition :exec
UPDATE characters SET definition = $3 WHERE id = $1 AND owner_id = $2
`

type UpdateCharacterDefinitionParams struct {
	ID         int32
	OwnerID    int32
	Definition json.RawMessage
}

func (q *Queries) UpdateCharacterDefinition(ctx context.Context, arg UpdateCharacterDefinitionParams) error {
	_, err := q.db.ExecContext(ctx, updateCharacterDefinition, arg.ID, arg.OwnerID, arg.Definition)
	return err
}

const updateCharacterName = `-- name: UpdateCharacterName :exec
UPDATE characters SET name = $3 WHERE id = $1 AND owner_id = $2
`

type UpdateCharacterNameParams struct {
	ID      int32
	OwnerID int32
	Name    string
}

func (q *Queries) UpdateCharacterName(ctx context.Context, arg UpdateCharacterNameParams) error {
	_, err := q.db.ExecContext(ctx, updateCharacterName, arg.ID, arg.OwnerID, arg.Name)
	return err
}

const updateCharacterSprite = `-- name: UpdateCharacterSprite :exec
UPDATE characters SET sprite = $3 WHERE id = $1 AND owner_id = $2
`

type UpdateCharacterSpriteParams struct {
	ID      int32
	OwnerID int32
	Sprite  []byte
}

func (q *Queries) UpdateCharacterSprite(ctx context.Context, arg UpdateCharacterSpriteParams) error {
	_, err := q.db.ExecContext(ctx, updateCharacterSprite, arg.ID, arg.OwnerID, arg.Sprite)
	return err
}

const updateEncounterDescription = `-- name: UpdateEncounterDescription :exec
UPDATE encounters SET description = $3 WHERE encounters.id = $1 AND
EXISTS (SELECT id FROM campaigns WHERE campaigns.id = encounters.campaign_id AND owner_id = $2)
`

type UpdateEncounterDescriptionParams struct {
	ID          int32
	OwnerID     int32
	Description sql.NullString
}

func (q *Queries) UpdateEncounterDescription(ctx context.Context, arg UpdateEncounterDescriptionParams) error {
	_, err := q.db.ExecContext(ctx, updateEncounterDescription, arg.ID, arg.OwnerID, arg.Description)
	return err
}

const updateEncounterName = `-- name: UpdateEncounterName :exec
UPDATE encounters SET name = $3 WHERE encounters.id = $1 AND
EXISTS (SELECT id FROM campaigns WHERE campaigns.id = encounters.campaign_id AND owner_id = $2)
`

type UpdateEncounterNameParams struct {
	ID      int32
	OwnerID int32
	Name    string
}

func (q *Queries) UpdateEncounterName(ctx context.Context, arg UpdateEncounterNameParams) error {
	_, err := q.db.ExecContext(ctx, updateEncounterName, arg.ID, arg.OwnerID, arg.Name)
	return err
}

const updateEncounterTilemap = `-- name: UpdateEncounterTilemap :exec
UPDATE encounters SET tilemap_id = $3 WHERE encounters.id = $1 AND
EXISTS (SELECT id FROM campaigns WHERE campaigns.id = encounters.campaign_id AND owner_id = $2)
`

type UpdateEncounterTilemapParams struct {
	ID        int32
	OwnerID   int32
	TilemapID sql.NullInt32
}

func (q *Queries) UpdateEncounterTilemap(ctx context.Context, arg UpdateEncounterTilemapParams) error {
	_, err := q.db.ExecContext(ctx, updateEncounterTilemap, arg.ID, arg.OwnerID, arg.TilemapID)
	return err
}
