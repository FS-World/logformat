package logformat

import (
	"encoding/json"

	"github.com/tinylib/msgp/msgp"
)

// Serializable is an interface that all data to be persisted as logdata must satisfy.
type Serializable interface {
	msgp.Marshaler
	msgp.Unmarshaler
	msgp.Encodable
	msgp.Decodable
	msgp.Sizer
	json.Marshaler
	json.Unmarshaler
}
