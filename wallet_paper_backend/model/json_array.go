package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type StringArray []string

func (a StringArray) Value() (driver.Value, error) {
    return json.Marshal(a)
}

func (a *StringArray) Scan(value interface{}) error {
    bytes, ok := value.([]byte)
    if !ok {
        return fmt.Errorf("failed to unmarshal JSONB value: %v", value)
    }
    return json.Unmarshal(bytes, a)
}
