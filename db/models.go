// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type SecretMessage struct {
	ID      pgtype.UUID `json:"id"`
	Message string      `json:"message"`
}
