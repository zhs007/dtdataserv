package dtdata

import "errors"

var (
	// ErrNoServerAddress - no server address
	ErrNoServerAddress = errors.New("no server address")
	// ErrNoAnkaDBConfig - no ankadb config
	ErrNoAnkaDBConfig = errors.New("no ankadb config")
	// ErrInvliadDTDataType - invalid dtdata type
	ErrInvliadDTDataType = errors.New("invalid dtdata type")
)
