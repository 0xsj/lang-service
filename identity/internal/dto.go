package internal

import "time"

// DIDDTO represents the data transfer object for a DID.
type DIDDTO struct {
    DID                string    `json:"did"`
    Method             string    `json:"method"`
    MethodSpecificID   string    `json:"method_specific_id"`
    CreatedAt          time.Time `json:"created_at"`
}

// ChallengeDTO represents the data transfer object for a Challenge.
type ChallengeDTO struct {
    ChallengeID   string    `json:"challenge_id"`
    UserID        string    `json:"user_id"`
    Challenge     string    `json:"challenge"`
    IssuedAt      time.Time `json:"issued_at"`
    ExpiresAt     time.Time `json:"expires_at"`
}
