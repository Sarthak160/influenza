package routes

import (
	"github.com/gin-gonic/gin"
	get "github.com/shivamsouravjha/influenza/controllers/GET"
	post "github.com/shivamsouravjha/influenza/controllers/POST"
	"github.com/shivamsouravjha/influenza/helpers/mongo"
	"github.com/shivamsouravjha/influenza/helpers/redis"
)

func NewRouter() *gin.Engine {
	redis.RedisInit()
	mongo.InitDb()
	router := gin.New()
	v1 := router.Group("/api")

	emailRoutes := v1.Group("/email")
	{
		emailRoutes.GET("/getVerificationCode", get.VerificationCode)
		emailRoutes.POST("/verifyCode", post.VerifyCode)
	}
	userRoutes := v1.Group("/post")
	{
		userRoutes.POST("/postFeedback")
		userRoutes.GET("/getFeedback", get.GetUserFeedback)
	}
	return router
}
