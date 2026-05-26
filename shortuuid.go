// Package shortuuid provides a library for generating concise, unambiguous,
// URL-safe UUIDs. It generates UUIDs using github.com/google/uuid and then
// translates them to base57 using a custom alphabet that removes similar-looking
// characters (l, 1, I, O, 0).
//
// The package is compatible with the Python library shortuuid and provides
// both a default encoder (base57) and support for custom alphabets and encoders.
package shortuuid

import (
	"github.com/google/uuid"
)

// DefaultEncoder is the default encoder used when generating new UUIDs, and is
// based on Base57.
var DefaultEncoder = b57Encoder{}

// Encoder is an interface for encoding/decoding UUIDs to strings.
type Encoder interface {
	Encode(uuid.UUID) string
	Decode(string) (uuid.UUID, error)
}

// New returns a new UUIDv4, encoded with base57.
func New() string { _ = "STUB: not implemented"; return "" }

// NewWithEncoder returns a new UUIDv4, encoded with enc.
func NewWithEncoder(enc Encoder) string { _ = "STUB: not implemented"; return "" }

// NewWithNamespace returns a new UUIDv5 (or v4 if name is empty), encoded with base57.
func NewWithNamespace(name string) string { _ = "STUB: not implemented"; return "" }

// NewWithAlphabet returns a new UUIDv4, encoded using the alternative
// alphabet abc.
//
// Panics if abc (after removing duplicates) has fewer than 2 characters.
// The alphabet will be automatically sorted and deduplicated to ensure
// consistency.
func NewWithAlphabet(abc string) string { _ = "STUB: not implemented"; return "" }

func hasPrefixCaseInsensitive(s, prefix string) bool { _ = "STUB: not implemented"; return false }

func hashedUUID(space uuid.UUID, data string) (u uuid.UUID) {
	_ = "STUB: not implemented"
	return *new(uuid.UUID)
}

// RFC 4122 variant
