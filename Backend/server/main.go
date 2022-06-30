package main

import (
	"net/http"

	"bidding/server/controllers"
	"bidding/server/utils"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"admin": "123", // user:admin password:123
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("admin:123")
		curl -X POST \
	  	http://localhost:8080/admin/start \
	  	-H 'authorization: Basic YWRtaW46MTIz \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'

		this will be changed in signing way by owner key, on live version.
	*/
	authorized.POST("admin/start", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}
		controllers.StartGame()
		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

func main() {
	r := setupRouter()
	utils.TotalBidAmount()

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
