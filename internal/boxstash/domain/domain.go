package domain

import (
	"time"
)

//-----------------------------------------------------------------------------
// Helpers
//-----------------------------------------------------------------------------

// Timestamper defines methods to set timestamps on types that need to do so
type Timestamper interface {
	CreatedTimestamps()
	UpdatedTimestamps()
}

// Parameterizer turns structs into maps for use in db queries
type Parameterizer interface {
	ToParams() map[string]interface{}
}

// ToInternetTime converts a UNIX timestamp to an RFC-3339 timestamp
func ToInternetTime(t int64) string {
	if t == 0 {
		return ""
	}
	return time.Unix(t, 0).Format(time.RFC3339)
}

// ToUnixTime converts an RFC 3339 timestamp to a UNIX timestamp
func ToUnixTime(t string) int64 {
	ts, err := time.Parse(time.RFC3339, t)
	if err != nil {
		return 0
	}
	return ts.Unix()
}
