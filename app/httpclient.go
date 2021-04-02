package app

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	v "github.com/pwdz/gurl/pkg/validation"
)

type Request struct{
	url		string

}

const(
	DefaultMethod = "GET"
	DefaultContentTypeValue = "application/x-www-form-urlencoded"
	JSONContentTypeValue = "application/json"
	FileContentTypeValue = "application/octet-stream"
	contentTypeKey = "content-type"

)
var Methods = []string{"GET", "POST", "PATCH", "DELETE", "PUT"}

func Send(url, method string, rawHeaders, rawQuerries []string, data, json, filePath string, timeout int) error {
	// if !pkg.IsValidUrl(url){
	// 	return log.Errorf("URL is not valid.")
	// }

	if !v.IsMethodValid(method, Methods) {
		return fmt.Errorf("Method " + method + " is not valid. supported methods: GET, POST, PATCH, DELETE, PUT")
	}

	client := http.Client{};
	
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
	}else if filePath != "" {
		if err := addFile(request, filePath); err!= nil{
			return err
		}
	}

	log.Println("Sending request to: ", request.URL)

	if timeout > 0{
		timer := time.AfterFunc(time.Second*time.Duration(timeout), func() {
			
			client.CloseIdleConnections()
			log.Fatal("Time limit exceeded")
		})
		defer timer.Stop()
	}

	resp, err := client.Do(request)
	if err != nil{
		return err
	}

	parseResponse(resp)

	return nil
}
func addCustomHeaders(request *http.Request, rawHeaders []string){
	headersMap, warning := parseHeaders(rawHeaders)

	log.Println("Add custom headers...")
	for key, value := range headersMap{ 
		request.Header.Add(strings.ToLower(key),value)
	}
	if warning != ""{
		log.Println(warning)
	}
}
func addQuerries(request *http.Request, rawQuerries []string){
	querriesMap := parseQuerries(rawQuerries)

	log.Println("Add querries...")
	q := request.URL.Query()
	for key, value := range querriesMap{ 
		q.Add(key, value)
	}
	request.URL.RawQuery = q.Encode()
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
	log.Println("Adding data to body...")
	request.Body = ioutil.NopCloser(bytes.NewBufferString(data))

	if request.Header.Get(contentTypeKey) == ""{	
		request.Header.Add(contentTypeKey, DefaultContentTypeValue)
	}

	if !v.IsDataValid(data){
		log.Println("Warning: data body doesn't match " + DefaultContentTypeValue + ":\n", data)
	}
}
func addJson(request *http.Request, json string){
	log.Println("Adding json to body...")
	request.Body = ioutil.NopCloser(bytes.NewBufferString(json))
	if request.Header.Get(contentTypeKey) == "" {
		request.Header.Add(contentTypeKey, JSONContentTypeValue)
	}

	if !v.IsJSONValid(json){
		log.Println("Warning: data body doesn't match " + JSONContentTypeValue + ":\n ", json)
	} 
}
func addFile(request *http.Request, filePath string) error {
	log.Println("Adding file to body...")
	fileBytes, err := ioutil.ReadFile(filePath)
    if err != nil {
        return err
    }
	request.Body = ioutil.NopCloser(bytes.NewBuffer(fileBytes))

	if request.Header.Get(contentTypeKey) == "" {
		request.Header.Add(contentTypeKey, FileContentTypeValue)
	}
	return nil
}
func parseResponse(resp *http.Response){
	defer resp.Body.Close()
	log.Println("\n============================\nReponse:")
	fmt.Println(resp.Status)
	fmt.Println(resp.Request.Method)
	fmt.Println("===\nHeaders:")
	for key, value := range resp.Header{
		fmt.Println(key , ":" , value)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println("\nBody:\n" , string(body))
}