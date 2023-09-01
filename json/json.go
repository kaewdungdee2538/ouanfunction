package json

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
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

func SetStructToBodyRequestForMiddleware(c *gin.Context, req interface{}) error {
	// ---------Convert obj to json string
	jsonData, err := json.Marshal(req)
	if err != nil {
		msgErr := fmt.Sprintf("Set model to body request Error : %s", err)
		return errors.New(msgErr)
	}
	// -----forward request body middleware to endpoint
	rdr2 := io.NopCloser(bytes.NewBuffer([]byte(fmt.Sprintf("%v", string(jsonData)))))
	c.Request.Body = rdr2
	return nil
}
