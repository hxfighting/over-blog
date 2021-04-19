// +build stdjson

package json

import "encoding/json"

var (
	// Marshal is exported by json package.
	Marshal = json.Marshal
	// Unmarshal is exported by json package.
	Unmarshal = json.Unmarshal
	// MarshalIndent is exported by json package.
	MarshalIndent = json.MarshalIndent
	// NewDecoder is exported by json package.
	NewDecoder = json.NewDecoder
	// NewEncoder is exported by json package.
	NewEncoder = json.NewEncoder
)
