package smartid

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

// A SmartID has size of 27 bytes which is encoded size of a 160 bit (20 byte) random identifier.
type SmartID [27]byte

var (
	rander  = rand.Reader
	enc     = base64.URLEncoding.WithPadding(base64.NoPadding)
	EmptyID = SmartID{}
)

// MustNew creates new SmartID or panics.
func MustNew() SmartID {
	return Must(New())
}

// MustNewString creates new SmartID and returns it as string or panics.
func MustNewString() string {
	return Must(New()).String()
}

// New creates new SmartID.
func New() (SmartID, error) {
	var unencodedID [20]byte

	_, err := io.ReadFull(rander, unencodedID[:])

	if err != nil {
		return EmptyID, err
	}

	encodedID := make([]byte, 27)
	enc.Encode(encodedID, unencodedID[:])

	var sID SmartID
	copy(sID[:], encodedID)

	return sID, nil
}

// Must returns SmartID if err is nil and panics otherwise.
func Must(ID SmartID, err error) SmartID {
	if err != nil {
		panic(err)
	}

	return ID
}

// String returns string from SmartID
func (s SmartID) String() string {
	return string(s[:])
}
