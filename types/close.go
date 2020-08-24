package types

// CloseType ...
type CloseType string

const (
	// INGRESS STATUS .... is no longer used  from session pool
	INGRESS CloseType = "INGRESS"
	// EGRESS STATUS .... is the session up and running
	EGRESS CloseType = "EGRESS"
	// ALL .... is the session DOWN but recoverable
	ALL CloseType = "ALL"
)
