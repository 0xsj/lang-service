package internal

import "time"

type OtpEntity struct {
	UserID string 
	Token string 
	IssuedAt time.Time
}

type SessionTokenEntity struct {
	UserID string 
	Token string 
	IssuedAt time.Time
}

type ChallengeEntity struct {
	UserID string 
	Challenge string 
	IssuedAt time.Ticker
}