package internal

import "time"

// CreateDIDRequest is the DTO used for requesting the creation of a DID.
type CreateDIDRequest struct {
    UserID string // The ID of the user for whom the DID is being created
}

// CreateDIDResponse is the DTO returned after creating a DID.
type CreateDIDResponse struct {
    DID      string    // The newly created DID
    Message  string    // Additional information, e.g., success or error messages
    IssuedAt time.Time // Timestamp when the DID was created
}

// ResolveDIDRequest is the DTO used for requesting the resolution of a DID.
type ResolveDIDRequest struct {
    DID string // The DID to be resolved
}

// ResolveDIDResponse is the DTO returned after resolving a DID.
type ResolveDIDResponse struct {
    DIDDocument string    // The DID document associated with the resolved DID
    Message     string    // Additional information, e.g., success or error messages
    ResolvedAt  time.Time // Timestamp when the DID was resolved
}

// IssueChallengeRequest is the DTO used for requesting the issuance of a challenge.
type IssueChallengeRequest struct {
    UserID string // The ID of the user for whom the challenge is issued
}

// IssueChallengeResponse is the DTO returned after issuing a challenge.
type IssueChallengeResponse struct {
    Challenge string    // The cryptographic challenge (nonce) to be signed
    IssuedAt  time.Time // Timestamp when the challenge was issued
    Message   string    // Additional information, e.g., success or error messages
}

// VerifySignedChallengeRequest is the DTO used for verifying a signed challenge.
type VerifySignedChallengeRequest struct {
    UserID          string // The ID of the user associated with the signed challenge
    SignedChallenge string // The signed challenge provided by the user
}

// VerifySignedChallengeResponse is the DTO returned after verifying a signed challenge.
type VerifySignedChallengeResponse struct {
    Valid   bool   // Indicates if the signed challenge is valid
    Message string // Additional information, e.g., error messages
}
