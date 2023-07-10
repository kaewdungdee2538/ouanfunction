package jsonUtil

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type JSON map[string]interface{}

func (a JSON) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *JSON) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return nil
	}

	raw := value
	j, err := json.Marshal(raw)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(j))

	var c map[string]interface{}
	json.Unmarshal(j, &c)
	fmt.Println(string(j))

	return json.Unmarshal(b, &a)
}
