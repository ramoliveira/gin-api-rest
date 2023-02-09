package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/ramoliveira/api-go-gin/controllers"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/:name", controller.Greetings)

	r.GET("/students", controller.ShowAllStudents)
	r.POST("/students", controller.CreateNewStudent)
	r.GET("/students/:id", controller.ShowStudentById)
	r.GET("/students/cpf/:cpf", controller.ShowStudentByDocument)
	r.DELETE("/students/:id", controller.DeleteStudentById)
	r.PATCH("/students/:id", controller.UpdateStudentById)

	r.Run()
}
