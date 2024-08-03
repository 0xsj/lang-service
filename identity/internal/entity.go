package internal

import "time"

type DIDEntity struct {
	DID string 
	DIDDocument string 
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ChallengeEntity struct {
	UserID string 
	Challenge string  // nonce
	IssuedAt time.Time
}