package pkg

import (
	"encoding/json"
	"net/url"
	"sort"

	"github.com/pwdz/gurl/app"
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

func IsValidJSON(data string) bool {
	var marshaled map[string]interface{}
	return json.Unmarshal([]byte(data), &marshaled) == nil
}
func IsMethodValid(method string) bool{
	for _, m := range app.Methods{
		if m == method{
			return true
		}
	}
	return false
}