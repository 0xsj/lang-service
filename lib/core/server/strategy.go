package server

type TransportStrategy interface {
	Setup() error 
	Serve() error
}
