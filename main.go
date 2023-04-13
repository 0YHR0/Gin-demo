package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"logic-app-backend/controller"
)

func main() {
	//create a server
	ginServer := gin.Default()

	//set CORS
	ginServer.Use(cors.Default())

	//load the html
	ginServer.LoadHTMLGlob("templates/*")

	//load the static
	ginServer.Static("/static", "./static")

	//404
	ginServer.NoRoute(func(context *gin.Context) {
		context.HTML(404, "404.html", gin.H{})
	})

	//create question
	controller.GetAllQuestions(ginServer)

	//store the result
	controller.StoreResults(ginServer)

	//get the result
	controller.GetResults(ginServer)

	//delete question
	controller.DeleteQuestion(ginServer)

	//create question
	controller.CreateQuestion(ginServer)

	//read from config
	// ref: https://www.golinuxcloud.com/golang-parse-yaml-file/#Example-1_Parse_YAML_variable_into_map
	yamlFile, err := ioutil.ReadFile("env.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}

	var config map[string]map[string]interface{} //var _config *config.Config
	//将配置文件读取到结构体中
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Println(err.Error())
	}

	port := fmt.Sprintf("%v", config["Server"]["Port"])
	//run
	ginServer.Run(":" + port)

}

//func Cors() gin.HandlerFunc {
//	return func(context *gin.Context) {
//		method := context.Request.Method
//
//		context.Header("Access-Control-Allow-Origin", "http://127.0.0.1:8080")
//		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, x-token")
//		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")
//		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
//		context.Header("Access-Control-Allow-Credentials", "true")
//
//		if method == "OPTIONS" {
//			context.AbortWithStatus(http.StatusNoContent)
//		}
//	}
//}
