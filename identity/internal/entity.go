package internal

import (
	"time"
)

// DID represents the database entity for a DID.
type DID struct {
    ID                uint      `json:"id"` // Unique identifier for the DID
    DID               string    `json:"did"`
    Method            string    `json:"method"`
    MethodSpecificID  string    `json:"method_specific_id"`
    CreatedAt         time.Time `json:"created_at"`
}

// Challenge represents the database entity for a Challenge.
type Challenge struct {
    ID          uint      `json:"id"` // Unique identifier for the Challenge
    ChallengeID string    `json:"challenge_id"`
    UserID      string    `json:"user_id"`
    Challenge   string    `json:"challenge"`
    IssuedAt    time.Time `json:"issued_at"`
    ExpiresAt   time.Time `json:"expires_at"`
}
