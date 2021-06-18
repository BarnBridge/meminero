package types

import (
	"database/sql/driver"
	"encoding/hex"
)

// ByteArray is a custom type that maps to a the database `bytea` fields
type ByteArray string

func (val *ByteArray) Scan(value interface{}) error {
	encoded := hex.EncodeToString(value.([]byte))
	*val = ByteArray(encoded)

	return nil
}

func (val ByteArray) Value() (driver.Value, error) {
	return hex.DecodeString(string(val))
}

func (val ByteArray) String() string {
	return string(val)
}
