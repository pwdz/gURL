package pkg

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)
func IsValidUrl(data string) bool {
	_, err := url.ParseRequestURI(data)
	if err != nil {
		return false
	}

	u, err := url.Parse(data)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

func IsJSONValid(data string) bool {
	var marshaled map[string]interface{}
	return json.Unmarshal([]byte(data), &marshaled) == nil
}
func IsMethodValid(method string, methods []string) bool{
	fmt.Println("Validating method...")
	for _, m := range methods{
		if m == method{
			return true
		}
	}
	return false
}
func InMapExists(m map[string]string, value string) bool {
	_, ok := m[value]
	return ok
}
func IsDataValid(data string) bool{
	parts := strings.Split(data, "&")
	for _, part := range parts{
		keyValue := strings.Split(part, "=")
		if len(keyValue) > 1{
		}else{
			return false
		}

	}
	return true
}