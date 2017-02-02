package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// DetailsMap to get the Movements Details
type DetailsMap map[string]interface{}

// Value = This one implements the driver.Value to JSONB
func (p DetailsMap) Value() (driver.Value, error) {
	j, err := json.Marshal(p)
	return j, err
}

// Scan = This one implements the sql.Scanner to JSONB
func (p *DetailsMap) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion .([]byte) failed")
	}

	var i interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	*p, ok = i.(map[string]interface{})
	if !ok {
		return errors.New("type assertion .(map[string]interface{}) failed")
	}

	return nil
}
