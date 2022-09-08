# Todo List

## How to run
- Clone or fork this project on your local.
- Run project
```go
$ go run main.go

[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /v1/books                 --> pustaka-api/handler.(*bookHandler).FindAll-fm (3 handlers)
[GIN-debug] GET    /v1/books/:id             --> pustaka-api/handler.(*bookHandler).FindById-fm (3 handlers)
[GIN-debug] POST   /v1/books                 --> pustaka-api/handler.(*bookHandler).PostBooksHandler-fm (3 handlers)
[GIN-debug] PUT    /v1/books/:id             --> pustaka-api/handler.(*bookHandler).Update-fm (3 handlers)
[GIN-debug] DELETE /v1/books/:id             --> pustaka-api/handler.(*bookHandler).Delete-fm (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
[GIN-debug] Listening and serving HTTP on :8080

```
 
