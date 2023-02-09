package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ramoliveira/api-go-gin/database"
	models "github.com/ramoliveira/api-go-gin/models"
)

func Greetings(ctx *gin.Context) {
	name := ctx.Params.ByName("name")
	ctx.JSON(200, gin.H{
		"API says:": "Wassup, " + name + "?",
	})
}

func CreateNewStudent(ctx *gin.Context) {
	var student models.Student
	if err := ctx.ShouldBindJSON(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	database.DB.Create(&student)
	ctx.JSON(http.StatusOK, student)
}

func ShowAllStudents(ctx *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	ctx.JSON(200, students)
}

func ShowStudentById(ctx *gin.Context) {
	var student models.Student
	id := ctx.Params.ByName("id")
	database.DB.First(&student, id)

	if student.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Aluno não encontrado",
		})
		return
	}

	ctx.JSON(http.StatusOK, student)
}

func ShowStudentByDocument(ctx *gin.Context) {
	var student models.Student
	document := ctx.Param("cpf")
	database.DB.Where(&models.Student{CPF: document}).First(&student)

	if student.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Aluno não encontrado",
		})
		return
	}

	ctx.JSON(http.StatusOK, student)
}

func DeleteStudentById(ctx *gin.Context) {
	var student models.Student
	id := ctx.Params.ByName("id")

	database.DB.Delete(&student, id)
	ctx.JSON(http.StatusOK, gin.H{
		"data": "Aluno deletado com sucesso!",
	})
}

func UpdateStudentById(ctx *gin.Context) {
	var student models.Student
	id := ctx.Params.ByName("id")
	database.DB.First(&student, id)

	if err := ctx.ShouldBindJSON(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Model(&student).UpdateColumns(student)
	ctx.JSON(http.StatusOK, student)
}
