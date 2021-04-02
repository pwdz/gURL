# gurl
go http client (a simpler version of curl)
### Run
```
go run main.go URL [flags]
```
Flags:
```
  -D, --data string      (OPTIONAL) pass body data
  -F, --file string      (OPTIONAL) pass location of the file
  -H, --header strings   (OPTIONAL) pass headers
  -h, --help             help for gurl
  -J, --json string      (OPTIONAL) pass body in json format
  -M, --method string    (OPTIONAL) pass method(GET/POST/PATCH/DELETE/PUT). default value is GET. (default "GET")
  -Q, --query strings    (OPTIONAL) pass querries
  -T, --timeout int      (OPTIONAL) request timeout. default is infinite
  ```