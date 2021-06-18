package types

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// DatetimeToJSONUnix binds a time.Time to a `timestamp` database field
// when marshaled to JSON, outputs a unix timestamp
type DatetimeToJSONUnix time.Time

func (t DatetimeToJSONUnix) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).UTC().Unix())
}

func (t *DatetimeToJSONUnix) UnmarshalJSON(data []byte) error {
	var unix int64

	err := json.Unmarshal(data, &unix)
	if err != nil {
		return err
	}

	*t = DatetimeToJSONUnix(time.Unix(unix, 0).UTC())

	return nil
}

func (t *DatetimeToJSONUnix) Scan(value interface{}) error {
	*t = DatetimeToJSONUnix(value.(time.Time))

	return nil
}

func (t DatetimeToJSONUnix) Value() (driver.Value, error) {
	return time.Time(t), nil
}

func (t DatetimeToJSONUnix) String() string {
	return time.Time(t).UTC().String()
}
