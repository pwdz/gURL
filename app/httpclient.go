package app

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	// "time"

	v "github.com/pwdz/gurl/pkg/validation"
)

type Request struct{
	url		string

}

const(
	DefaultMethod = "GET"
	DefaultContentTypeValue = "application/x-www-form-urlencoded"
	JSONContentTypeValue = "application/json"
	contentTypeKey = "content-type"
)
var Methods = []string{"GET", "POST", "PATCH", "DELETE", "PUT"}

func Send(url, method string, rawHeaders, rawQuerries []string, data, json string, timeout int) error {
	// if !pkg.IsValidUrl(url){
	// 	return fmt.Errorf("URL is not valid.")
	// }
	if !v.IsMethodValid(method, Methods) {
		return fmt.Errorf("Method " + method + " is not valid. supported methods: GET, POST, PATCH, DELETE, PUT")
	}

	// client := http.Client{
	// 	Timeout: time.Duration(timeout * int(time.Second)),
	// };
	
	//init request
	log.Println("Init request")
	request, err := http.NewRequest(method, url, nil)

	if err != nil{
		return err
	}
	
	addCustomHeaders(request, rawHeaders)
	addQuerries(request, rawQuerries)

	if data != "" {
		addData(request, data)
	}else if json != "" {
		addJson(request, json)
	}

	

	fmt.Println(request)

	return nil
}
func addCustomHeaders(request *http.Request, rawHeaders []string){
	headersMap, warning := parseHeaders(rawHeaders)

	log.Println("Add custom headers...")
	for key, value := range headersMap{ 
		request.Header.Add(strings.ToLower(key),value)
	}
	if warning != ""{
		fmt.Println(warning)
	}
}
func addQuerries(request *http.Request, rawQuerries []string){
	querriesMap := parseQuerries(rawQuerries)

	log.Println("Add querries...")
	for key, value := range querriesMap{ 
		request.URL.Query().Add(key, value)
	}
}
func parseHeaders(rawHeaders []string) (map[string]string, string) {
	if rawHeaders == nil{
		return nil, ""
	}

	headersMap := make(map[string]string)
	warning := ""

	for _, rawHeader := range rawHeaders{ //rawHeader = "key1:value1,key2:value2"
		for _, header := range strings.Split(rawHeader,","){
			keyValue := strings.Split(header, ":")
			
			if v.InMapExists(headersMap, keyValue[0]){
				warning += "[#] Header " + keyValue[0] + " already exists. Replacing " + headersMap[keyValue[0]] + " with " + keyValue[1] +".\n"
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
func addData(request *http.Request, data string){
	request.Body = ioutil.NopCloser(bytes.NewBufferString(data))

	if request.Header.Get(contentTypeKey) == ""{	
		request.Header.Add(contentTypeKey, DefaultContentTypeValue)
	}

	if !v.IsDataValid(data){
		fmt.Println("Not a valid data body:\n", data)
	}
}
func addJson(request *http.Request, json string){
	request.Body = ioutil.NopCloser(bytes.NewBufferString(json))
	if request.Header.Get(contentTypeKey) == "" {
		request.Header.Add(contentTypeKey, JSONContentTypeValue)
	}

	if !v.IsJSONValid(json){
		fmt.Println("Not a valid json:\n ", json)
	} 
}