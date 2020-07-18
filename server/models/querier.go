// Code generated by sqlc. DO NOT EDIT.

package models

import (
	"context"
)

type Querier interface {
	CreateCampaign(ctx context.Context, arg CreateCampaignParams) (Campaign, error)
	CreateCharacter(ctx context.Context, arg CreateCharacterParams) (Character, error)
	CreateEncounter(ctx context.Context, arg CreateEncounterParams) (Encounter, error)
	CreateSpritesheet(ctx context.Context, arg CreateSpritesheetParams) (Spritesheet, error)
	CreateTilemap(ctx context.Context, arg CreateTilemapParams) (Tilemap, error)
	CreateTilemapReference(ctx context.Context, arg CreateTilemapReferenceParams) error
	DeleteCampaign(ctx context.Context, arg DeleteCampaignParams) error
	DeleteCharacter(ctx context.Context, arg DeleteCharacterParams) error
	DeleteEncounter(ctx context.Context, arg DeleteEncounterParams) error
	DeleteSpritesheet(ctx context.Context, arg DeleteSpritesheetParams) error
	DeleteTilemap(ctx context.Context, arg DeleteTilemapParams) error
	GetOwnedCampaignByID(ctx context.Context, arg GetOwnedCampaignByIDParams) (Campaign, error)
	GetSpritesheetByHash(ctx context.Context, hash []byte) (Spritesheet, error)
	GetThumbnailByHash(ctx context.Context, hash []byte) (GetThumbnailByHashRow, error)
	GetTilemapByHash(ctx context.Context, hash []byte) (Tilemap, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	InsertTilemapThumbnail(ctx context.Context, arg InsertTilemapThumbnailParams) (int64, error)
	ListCampaignsByOwnerID(ctx context.Context, ownerID int32) ([]Campaign, error)
	ListCharactersByOwnerID(ctx context.Context, ownerID int32) ([]Character, error)
	ListCharactersForEncounter(ctx context.Context, encounterID int32) ([]Character, error)
	ListEncountersForCampaign(ctx context.Context, campaignID int32) ([]Encounter, error)
	ListRecentSpritesheets(ctx context.Context, limit int32) ([]Spritesheet, error)
	ListRecentTilemaps(ctx context.Context, limit int32) ([]Tilemap, error)
	ListSpritesheetsByOwnerID(ctx context.Context, ownerID int32) ([]ListSpritesheetsByOwnerIDRow, error)
	ListSpritesheetsForTilemap(ctx context.Context, tilemapID int32) ([]ListSpritesheetsForTilemapRow, error)
	ListThumbnailsForSpritesheets(ctx context.Context, dollar_1 []int32) ([]ListThumbnailsForSpritesheetsRow, error)
	ListThumbnailsForTilemaps(ctx context.Context, dollar_1 []int32) ([]ListThumbnailsForTilemapsRow, error)
	ListTilemapsByOwnerID(ctx context.Context, ownerID int32) ([]ListTilemapsByOwnerIDRow, error)
	UpdateCampaignDescription(ctx context.Context, arg UpdateCampaignDescriptionParams) error
	UpdateCampaignName(ctx context.Context, arg UpdateCampaignNameParams) error
	UpdateCharacterDefinition(ctx context.Context, arg UpdateCharacterDefinitionParams) error
	UpdateCharacterName(ctx context.Context, arg UpdateCharacterNameParams) error
	UpdateCharacterSprite(ctx context.Context, arg UpdateCharacterSpriteParams) error
	UpdateEncounterDescription(ctx context.Context, arg UpdateEncounterDescriptionParams) error
	UpdateEncounterName(ctx context.Context, arg UpdateEncounterNameParams) error
	UpdateEncounterTilemap(ctx context.Context, arg UpdateEncounterTilemapParams) error
	UpdateTilemapThumbnail(ctx context.Context, arg UpdateTilemapThumbnailParams) error
}

var _ Querier = (*Queries)(nil)
