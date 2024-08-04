package core

import "context"

// returns a new context value, propogate across API boundaries and go routines
func NewContext() context.Context {
	return context.Background()
}