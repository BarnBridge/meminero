package types

import (
	"database/sql/driver"
	"encoding/json"
)

// JSONObject binds a map[string]interface{} to a `jsonb` database field
type JSONObject map[string]interface{}

func (obj *JSONObject) Scan(value interface{}) error {
	err := json.Unmarshal(value.([]byte), obj)

	return err
}

func (obj JSONObject) Value() (driver.Value, error) {
	return json.Marshal(obj)
}
