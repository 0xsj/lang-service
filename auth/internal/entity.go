package internal

import "time"

// OTP represents the database entity for an OTP.
type OTP struct {
    ID         uint      `json:"id"` // Unique identifier for the OTP
    OTPToken   string    `json:"otp_token"`
    IssuedAt   time.Time `json:"issued_at"`
    ExpiresAt  time.Time `json:"expires_at"`
}

// Session represents the database entity for a Session.
type Session struct {
    ID          uint      `json:"id"` // Unique identifier for the Session
    SessionToken string   `json:"session_token"`
    IssuedAt    time.Time `json:"issued_at"`
    ExpiresAt   time.Time `json:"expires_at"`
}