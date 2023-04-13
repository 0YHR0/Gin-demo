package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"logic-app-backend/entity"
	"logic-app-backend/models"
	"net/http"
	"strconv"
)

func GetAllQuestions(ginServer *gin.Engine) {

	ginServer.GET("/admin", func(context *gin.Context) {
		questions := models.GetAllQuestions()
		context.HTML(http.StatusOK, "index.html", gin.H{
			"questions": questions,
		})
	})

	ginServer.GET("/getAllQuestions", func(context *gin.Context) {
		difficultyLevel := context.Query("difficulty")
		questions := models.GetAllQuestionsWithDifficultyLevel(difficultyLevel)
		context.JSON(http.StatusOK, questions)
	})

}

func StoreResults(ginServer *gin.Engine) {
	ginServer.GET("/storeResults", func(context *gin.Context) {
		correct := context.Query("correct")
		wrong := context.Query("wrong")
		time := context.Query("time")
		result := models.StoreResult(correct, wrong, time)
		context.JSON(http.StatusOK, gin.H{"result": result})
	})
}

func GetResults(ginServer *gin.Engine) {
	ginServer.GET("/getResults", func(context *gin.Context) {
		ids := context.Query("ids")
		questions := models.GetResult(ids)
		context.JSON(http.StatusOK, questions)
	})
}

func DeleteQuestion(ginServer *gin.Engine) {
	ginServer.GET("/deleteQuestion/:id", func(context *gin.Context) {
		id := context.Param("id")
		fmt.Println(models.DeleteQuestion(id))
		context.JSON(http.StatusOK, gin.H{"result": "deleted"})
	})
}

func CreateQuestion(ginServer *gin.Engine) {
	ginServer.POST("/createQuestion", func(context *gin.Context) {
		id, _ := strconv.Atoi(context.PostForm("ID"))
		questionText := context.PostForm("questionText")
		answerA := context.PostForm("answerA")
		answerB := context.PostForm("answerB")
		answerC := context.PostForm("answerC")
		answerD := context.PostForm("answerD")
		correctAnswer := context.PostForm("correctAnswer")
		difficultyLevel := context.PostForm("difficultyLevel")
		totalTryNum, _ := strconv.Atoi(context.PostForm("totalTryNum"))
		correctTryNum, _ := strconv.Atoi(context.PostForm("correctTryNum"))
		totalTime, _ := strconv.Atoi(context.PostForm("totalTime"))
		detailSolution := context.PostForm("detailSolution")
		question := entity.Question{
			Id:              id,
			Text:            questionText,
			A:               answerA,
			B:               answerB,
			C:               answerC,
			D:               answerD,
			Answer:          correctAnswer,
			DifficultyLevel: difficultyLevel,
			TotalTrialNum:   totalTryNum,
			CorrectTrialNum: correctTryNum,
			TotalTime:       totalTime,
			DetailSolution:  detailSolution,
		}
		result := models.CreateQuestion(question)
		context.JSON(http.StatusOK, gin.H{"result": result})

	})

}
