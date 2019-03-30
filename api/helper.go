package api

import (
	"crypto/rand"
	"encoding/hex"
	fmt "fmt"
	"io"
)

// GenerateID generates random feasible `Station` ID.
func GenerateID() string {
	b := make([]byte, 32)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		panic(err)
	}
	id := hex.EncodeToString(b)

	return id
}

// TruncateID truncates given string.
func TruncateID(id string) string {
	if len(id) < 12 {
		return ""
	}
	return id[0:12]
}

// ToShort returns truncated token value.
func (token *Token) ToShort() string {
	return TruncateID(token.GetId())
}

// ToShort returns truncated `Station` ID.
func (station *Station) ToShort() string {
	return TruncateID(station.GetId())
}

// ToString returns serialized description of `Station`.
func (station *Station) ToString() string {
	return fmt.Sprint(station.GetImage(), "~", station.GetName())
}
