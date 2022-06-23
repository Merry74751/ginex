# ginex
Some extensions have been made based on [GIN](https://github.com/gin-gonic/gin)
- Automatic route registration
- The encapsulation of the return value

## GetStart
```
go get github.com/Merry74751/ginex
```

## Example
```go
import (
	"github.com/Merry74751/ginex/router"
	"github.com/gin-gonic/gin"
	"testing"
)

type UserApi struct {}

func (UserApi) GetById() (string, string, router.Handle) {
	return "GET", "getById", func(c *gin.Context) (any, error) {
		s := "hello world"
		return s, nil
	}
}

func main() {
	app := gin.Default()
	router.addRouter(UserApi{})
	router.RegisterRouter(app)
	
	app.Run()
}
```
