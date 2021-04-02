package app

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	// "time"

	"github.com/pwdz/gurl/pkg/validation"
)

type Request struct{
	url		string

}

const DefaultMethod = "GET"
var Methods = [...]string{"GET", "POST", "PATCH", "DELETE", "PUT"}

func Send(url, method string, rawHeaders, rawQuerries []string, data, json string, timeout int) error {
	// if !pkg.IsValidUrl(url){
	// 	return fmt.Errorf("URL is not valid.")
	// }
	if !pkg.IsMethodValid(method) {
		return fmt.Errorf("Method is not valid. [supported methods: GET, POST, PATCH, DELETE, PUT")
	}

	// client := http.Client{
	// 	Timeout: time.Duration(timeout * int(time.Second)),
	// };
	
	//init request
	request, err := http.NewRequest(method, url, bytes.NewBufferString(data))

	if err != nil{
		return err
	}

	addHeaders(request, rawHeaders)
	addQuerries(request, rawQuerries)



	return nil
}
func addHeaders(request *http.Request, rawHeaders []string){
	headersMap, warning := parseHeaders(rawHeaders)
	for key, value := range headersMap{ 
		request.Header.Add(key, value)
	}
	if warning != nil{
		fmt.Println(warning)
	}
}
func addQuerries(request *http.Request, rawQuerries []string){
	querriesMap := parseQuerries(rawQuerries)
	for key, value := range querriesMap{ 
		request.URL.Query().Add(key, value)
	}
}
func parseHeaders(rawHeaders []string) (map[string]string, string) {
	if rawHeaders == nil{
		return nil, nil
	}

	headersMap := make(map[string]string)
	warning := ""

	for _, rawHeader := range rawHeaders{ //rawHeader = "key1:value1,key2:value2"
		for _, header := range strings.Split(rawHeader,","){
			keyValue := strings.Split(header, ":")
			
			if pkg.InMapExists(headersMap, keyValue[0]){
				warning += "Header " + keyValue[0] " already exists. Replacing " + headersMap[kekeyValue[0]] + " with " + keykeyValue[1] +"\n"
			}
			headersMap[keyValue[0]] = keyValue[1]
		}
	}

	return headersMap, warning
}
func parseQuerries(rawQuerries []string) map[string]string{ //rawQuerry = "key1=value1&key2=value2"
	if rawQuerries == nil{
		return nil
	}
	querriesMap := make(map[string]string)

	for _, rawQuery := range rawQuerries{ //rawHeader = "key1:value1,key2:value2"
		for _, query := range strings.Split(rawQuery,"&"){
			keyValue := strings.Split(query, "=")
			querriesMap[keyValue[0]] = keyValue[1]
		}
	}

	return querriesMap
}