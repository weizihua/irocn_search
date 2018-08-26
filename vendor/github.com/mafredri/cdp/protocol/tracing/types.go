// Code generated by cdpgen. DO NOT EDIT.

package tracing

import (
	"encoding/json"
	"errors"
)

// MemoryDumpConfig Configuration for memory dump. Used only when "memory-infra" category is enabled.
type MemoryDumpConfig []byte

// MarshalJSON copies behavior of json.RawMessage.
func (m MemoryDumpConfig) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON copies behavior of json.RawMessage.
func (m *MemoryDumpConfig) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("tracing.MemoryDumpConfig: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[0:0], data...)
	return nil
}

var _ json.Marshaler = (*MemoryDumpConfig)(nil)
var _ json.Unmarshaler = (*MemoryDumpConfig)(nil)

// TraceConfig
type TraceConfig struct {
	// RecordMode Controls how the trace buffer stores data.
	//
	// Values: "recordUntilFull", "recordContinuously", "recordAsMuchAsPossible", "echoToConsole".
	RecordMode           *string          `json:"recordMode,omitempty"`
	EnableSampling       *bool            `json:"enableSampling,omitempty"`       // Turns on JavaScript stack sampling.
	EnableSystrace       *bool            `json:"enableSystrace,omitempty"`       // Turns on system tracing.
	EnableArgumentFilter *bool            `json:"enableArgumentFilter,omitempty"` // Turns on argument filter.
	IncludedCategories   []string         `json:"includedCategories,omitempty"`   // Included category filters.
	ExcludedCategories   []string         `json:"excludedCategories,omitempty"`   // Excluded category filters.
	SyntheticDelays      []string         `json:"syntheticDelays,omitempty"`      // Configuration to synthesize the delays in tracing.
	MemoryDumpConfig     MemoryDumpConfig `json:"memoryDumpConfig,omitempty"`     // Configuration for memory dump triggers. Used only when "memory-infra" category is enabled.
}
