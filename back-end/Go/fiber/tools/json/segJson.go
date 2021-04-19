// +build !stdjson

package json

import segJson "github.com/segmentio/encoding/json"

var (
	// Marshal is exported by json package.
	Marshal = segJson.Marshal
	// Unmarshal is exported by json package.
	Unmarshal = segJson.Unmarshal
	// MarshalIndent is exported by json package.
	MarshalIndent = segJson.MarshalIndent
	// NewDecoder is exported by json package.
	NewDecoder = segJson.NewDecoder
	// NewEncoder is exported by json package.
	NewEncoder = segJson.NewEncoder
)
