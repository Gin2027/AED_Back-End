package routers

import (
	. "aed/apis"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc  {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func InitRouter() *gin.Engine  {
	router := gin.Default()

	router.Use(Cors())

	router.GET("/",IndexApi)

	router.GET("/md5/:str",MD5)

	router.GET("/GetAllStudents",GetAll)

	router.GET("/GetAllTeachers",TeachersList)

	router.POST("/Student/Login",Login)

	router.POST("/Student/Register",Register)

	router.POST("/Student/Modify",ModifyStudent)

	router.POST("/Student/Choose",ChooseTeacher)

	router.POST("/Student/GiveAdvice",GiveAdvice)

	router.POST("/Student/GetReplies",GetReplies)

	router.POST("/Teacher/Login",TLogin)

	router.POST("/Teacher/Register",TRegister)

	router.POST("/Teacher/Get",GetTeacher)

	router.POST("/Teacher/Modify",ModifyTeacher)

	router.POST("/Teacher/GetAdvices",AdviceList)

	router.POST("/Teacher/Reply",Reply)

	return router
}