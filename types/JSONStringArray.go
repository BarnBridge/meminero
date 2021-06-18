package types

import (
	"database/sql/driver"
	"encoding/json"
)

// JSONStringArray binds a slice of strings to a `jsonb` database field
type JSONStringArray []string

func (j *JSONStringArray) Scan(value interface{}) error {
	err := json.Unmarshal(value.([]byte), j)

	return err
}

func (j JSONStringArray) Value() (driver.Value, error) {
	return json.Marshal(j)
}
