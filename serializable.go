package logformat

import (
	"encoding/json"

	"github.com/tinylib/msgp/msgp"
)

// Serializable is an interface that all data to be persisted as log data must satisfy
type Serializable interface {
	msgp.Marshaler
	msgp.Unmarshaler
	msgp.Encodable
	msgp.Decodable
	msgp.Sizer
	json.Marshaler
	json.Unmarshaler
}

// LogReader is an interface for log readers
type LogReader interface {
	// ReadAll reads all log data from the file
	ReadAll() ([]Serializable, error)

	// Close closes the logfile
	Close() error
}

// LogWriter is an interface for log writers
type LogWriter interface {
	// Write writes a Serializable msg to the logfile
	Write(msg Serializable) error

	// Sync flushes the buffers and syncs the file to disk
	Sync() error

	// Close closes the logfile
	Close() error
}
