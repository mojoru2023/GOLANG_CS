package main


//gin-gonic/gin
//middleware的使用
//context的使用

import (
	_ "github.com/gin-gonic/gin"
	_ "go.uber.org/zap"
	"math/rand"
	"time"
)

const keyRequestId = "requestId"
func main()  {
	r := gin.Default()
	logger,err := zap.NewProduction()
	if err := nil{
		panic(err)
	}

	//middleware
	r.Use(
		func(c *gin.Context) {
			s := time.Now()
			c.Next()
			logger.Info("income request",
				zap.String("path", c.Request.URL.Path),
				zap.Int("status", c.Writer.Status()),
				zap.Duration("time", time.Now().Sub(s)),
				zap.Int(keyRequestId, c.GetInt(keyRequestId)),
			)
		},
		func(c *gin.Context) {
			c.Set(keyRequestId, rand.Int())
			c.Next()
		},
	)

	//router
	r.GET("/ping", func(c *gin.Context) {
		h := gin.H{
			"message":"pong",
		}

		if rid,exist := c.Get(keyRequestId);exist{
			h[keyRequestId] = rid
		}
		c.JSON(200,h)
	})

	r.GET("/hello",func(c *gin.Context){
		c.String(200,"hello")
	})
	r.Run()
}