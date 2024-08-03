package internal

import "time"

// OTPDTO represents the data transfer object for an OTP.
type OTPDTO struct {
    OTPToken  string    `json:"otp_token"`
    IssuedAt  time.Time `json:"issued_at"`
    ExpiresAt time.Time `json:"expires_at"`
}

// SessionDTO represents the data transfer object for a Session.
type SessionDTO struct {
    SessionToken string    `json:"session_token"`
    IssuedAt     time.Time `json:"issued_at"`
    ExpiresAt    time.Time `json:"expires_at"`
}

// ChallengeDTO is used to transfer challenge data within the auth service.
type ChallengeDTO struct {
    ChallengeID   string    `json:"challenge_id"`
    UserID        string    `json:"user_id"`
    Challenge     string    `json:"challenge"`
    IssuedAt      time.Time `json:"issued_at"`
    ExpiresAt     time.Time `json:"expires_at"`
}