package shortuuid

// DefaultAlphabet is the default alphabet used for base57 encoding.
// It excludes similar-looking characters (0, 1, I, O, l) to avoid confusion.
const (
	DefaultAlphabet = "23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
)

// alphabet represents a character set for base-N encoding. It stores the
// sorted, deduplicated characters along with precomputed values for efficient
// encoding and decoding.
type alphabet struct {
	chars    []rune // sorted, deduplicated characters
	len      int64  // number of characters in the alphabet
	encLen   uint8  // maximum encoded length for a 128-bit value
	maxBytes uint8  // maximum UTF-8 bytes needed for any character
}

// newAlphabet creates a new alphabet from the given string. Removes
// duplicates and sorts the characters to ensure reproducibility.
//
// Panics if the alphabet (after removing duplicates) has fewer than 2
// characters. An alphabet must have at least 2 characters to be usable for
// base-N encoding.
func newAlphabet(s string) alphabet { _ = "STUB: not implemented"; return *new(alphabet) }

func (a *alphabet) Length() int64 {
	_ = "STUB: not implemented"

	// Index returns the index of the first instance of t in the alphabet, or an
	// error if t is not present.
	return 0
}

func (a *alphabet) Index(t rune) (int64, error) { _ = "STUB: not implemented"; return 0, nil }
