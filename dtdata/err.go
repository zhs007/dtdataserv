package dtdata

import "errors"

var (
	// ErrNoServerAddress - no server address
	ErrNoServerAddress = errors.New("no server address")
	// ErrNoAnkaDBConfig - no ankadb config
	ErrNoAnkaDBConfig = errors.New("no ankadb config")
	// ErrInvliadDTDataType - invalid dtdata type
	ErrInvliadDTDataType = errors.New("invalid dtdata type")
	// ErrGameDayReport - GameDayReport error
	ErrGameDayReport = errors.New("GameDayReport error")
	// ErrNoHTTPServerAddr - no http server address
	ErrNoHTTPServerAddr = errors.New("no http server address")
	// ErrNoBusinessDayData - no BusinessDayData
	ErrNoBusinessDayData = errors.New("no BusinessDayData")
)
