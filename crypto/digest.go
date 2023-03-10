package crypto

import (
	"crypto/sha512"
	"encoding/base32"
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
)

// DigestSize is the number of bytes in the preferred hash Digest used here.
const DigestSize = sha512.Size256

// Digest represents a 32-byte value holding the 256-bit Hash digest.
type Digest [DigestSize]byte

// ToSlice converts Digest to slice, is used by bookkeeping.PaysetCommit
func (d Digest) ToSlice() []byte {
	return d[:]
}

// String returns the digest in a human-readable Base32 string
func (d Digest) String() string {
	return base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(d[:])
}

// TrimUint64 returns the top 64 bits of the digest and converts to uint64
func (d Digest) TrimUint64() uint64 {
	return binary.LittleEndian.Uint64(d[:8])
}

// IsZero return true if the digest contains only zeros, false otherwise
func (d Digest) IsZero() bool {
	return d == Digest{}
}

// DigestFromString converts a string to a Digest
func DigestFromString(str string) (d Digest, err error) {
	decoded, err := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(str)
	if err != nil {
		return d, err
	}
	if len(decoded) != len(d) {
		msg := fmt.Sprintf(`Attempted to decode a string which was not a Digest: "%v"`, str)
		return d, errors.New(msg)
	}
	copy(d[:], decoded[:])
	return d, err
}

// NewHash returns a sha512-256 object to do the same operation as Hash()
func NewHash() hash.Hash {
	return sha512.New512_256()
}

// Hash computes the SHASum512_256 hash of an array of bytes
func Hash(data []byte) Digest {
	return sha512.Sum512_256(data)
}
