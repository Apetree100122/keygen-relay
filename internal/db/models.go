// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql"
)

type AuditLog struct {
	ID         int64          `json:"id"`
	Action     string         `json:"action"`
	EntityType string         `json:"entity_type"`
	EntityID   string         `json:"entity_id"`
	Timestamp  sql.NullString `json:"timestamp"`
}

type License struct {
	ID             string         `json:"id"`
	File           []byte         `json:"file"`
	Key            string         `json:"key"`
	Claims         sql.NullInt64  `json:"claims"`
	LastClaimedAt  sql.NullString `json:"last_claimed_at"`
	LastReleasedAt sql.NullString `json:"last_released_at"`
	NodeID         sql.NullInt64  `json:"node_id"`
}

type Node struct {
	ID              int64          `json:"id"`
	Fingerprint     string         `json:"fingerprint"`
	ClaimedAt       sql.NullString `json:"claimed_at"`
	LastHeartbeatAt sql.NullString `json:"last_heartbeat_at"`
}
