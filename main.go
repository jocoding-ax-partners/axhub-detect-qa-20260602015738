package main
import (
  "os"
  "github.com/gin-gonic/gin"
)
func main(){ port:=os.Getenv("PORT"); if port==""{port="8080"}; r:=gin.New(); r.GET("/", func(c *gin.Context){c.String(200,"axhub go gin ok")}); r.Run("0.0.0.0:"+port)}
