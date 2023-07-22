package json

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

func ConvertInterfaceToJSON(obj map[string]interface{}) (bool, string) {
	obj_result, err := json.Marshal(obj)
	if err != nil {
		//--------create error log
		return true, ""
	}
	return false, string(obj_result)
}

func JsonBytesToMap(jsonBytes []byte) map[string]interface{} {
	result := make(map[string]interface{})
	json.Unmarshal(jsonBytes, &result)
	return result
}

// ---------------------------------------------------------------------------------------------------------------//
func PrettyStructJson(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}
