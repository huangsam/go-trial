package endpoint

import "github.com/gin-gonic/gin"

// HelloHandler returns a simple greeting as text.
func HelloHandler(c *gin.Context) {
	c.String(200, "Hello world")
}
