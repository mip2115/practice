package main

import (
  "net/http"

  "github.com/gin-gonic/contrib/static"
  "github.com/gin-gonic/gin"
  "fmt"
)


func main() {
  // Set the router as the default one shipped with Gin
  router := gin.Default()

  router.Use(static.Serve("/", static.LocalFile("./index.html", true)))

  // Setup route group for the API
  api := router.Group("/api")
  {
    api.GET("/", func(c *gin.Context) {
      c.JSON(http.StatusOK, gin.H {
        "message": "pong",
      })
    })
  }

  // input form should hit this function when it posts
  api.POST("/postform", PostForm)
  router.Run(":5000")
}

// should be able to retrieve the value netered into the form and print it out.
func PostForm(c *gin.Context) {
  c.Header("Content-Type", "application/json")
  q := c.Request.FormValue("query")
  fmt.Println("VALUE IS: ", q)

}



