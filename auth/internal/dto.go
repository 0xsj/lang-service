package internal

import "time"

type CreateOtpRequest struct {
	UserID   string    `json:"user_id"`   // ID of the user
	Token    string    `json:"token"`     // The OTP token
	IssuedAt time.Time `json:"issued_at"` // Timestamp when the OTP was issued
}

type CreateOtpResponse struct {
	Success bool   `json:"success"` // Indicates if the OTP creation was successful
	Message string `json:"message"` // Additional information (e.g., success or error messages)
}

// DTO for creating or updating a session token
type CreateSessionTokenRequest struct {
	UserID   string    `json:"user_id"`   // ID of the user
	Token    string    `json:"token"`     // The session token value
	IssuedAt time.Time `json:"issued_at"` // Timestamp when the token was issued
}

type CreateSessionTokenResponse struct {
	Success bool   `json:"success"` // Indicates if the session token creation was successful
	Message string `json:"message"` // Additional information (e.g., success or error messages)
}

// DTO for issuing or verifying a challenge
type CreateChallengeRequest struct {
	UserID    string    `json:"user_id"`    // ID of the user
	Challenge string    `json:"challenge"`  // The cryptographic challenge (nonce)
	IssuedAt  time.Time `json:"issued_at"`  // Timestamp when the challenge was issued
}

type CreateChallengeResponse struct {
	Success bool   `json:"success"` // Indicates if the challenge creation was successful
	Message string `json:"message"` // Additional information (e.g., success or error messages)
}

type VerifyChallengeRequest struct {
	UserID          string `json:"user_id"`          // ID of the user
	SignedChallenge string `json:"signed_challenge"` // The signed challenge from the user
}

type VerifyChallengeResponse struct {
	Valid   bool   `json:"valid"`   // Indicates if the signed challenge is valid
	Message string `json:"message"` // Additional information (e.g., error messages)
}
