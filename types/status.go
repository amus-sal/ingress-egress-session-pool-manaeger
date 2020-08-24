package types

// STATUS ...
type STATUS string

const (
	// DELETED STATUS .... is no longer used  from session pool
	DELETED STATUS = "DELETED"
	// UP STATUS .... is the session up and running
	UP STATUS = "UP"
	// DOWN .... is the session DOWN but recoverable
	DOWN STATUS = "DOWN"
)
