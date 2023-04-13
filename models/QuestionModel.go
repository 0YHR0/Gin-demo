package models

import (
	//"GoReadConfig/config"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"logic-app-backend/entity"
	"strconv"
	"strings"
)

var db *gorm.DB

type TimeJson struct {
	Id                string `json:"id"`
	ConsumingDuration string `json:"consumingDuration"`
}

type ReturnType struct {
	Id                 int     `json:"id"`
	AverageCorrectness float32 `json:"avg_correctness"`
	AverageTime        float32 `json:"avg_time"`
}

func init() {
	//导入配置文件
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
	fmt.Println(config["DB"]["Username"])
	//username := _config.DB.Username
	//配置MySQL连接参数
	username := fmt.Sprintf("%v", config["DB"]["Username"]) //账号
	password := fmt.Sprintf("%v", config["DB"]["Password"]) //密码
	host := fmt.Sprintf("%v", config["DB"]["Host"])         //数据库地址，可以是Ip或者域名
	port := config["DB"]["Port"].(int)                      //数据库端口
	Dbname := fmt.Sprintf("%v", config["DB"]["Dbname"])     //数据库名
	timeout := "10s"                                        //连接超时，10秒
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	dbMysql, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	} else {
		db = dbMysql
	}

}

// GetAllQuestions from the database/**
func GetAllQuestions() []entity.QuestionDTO {
	var questions []entity.Question
	//text := "text_test"
	//db.Debug().Find(&questions, "text", text)
	db.Debug().Find(&questions)
	var questionDTOs []entity.QuestionDTO
	for i := 0; i < len(questions); i++ {
		if questions[i].TotalTrialNum == 0 {
			questionDTOs = append(questionDTOs, entity.QuestionDTO{questions[i],
				0.25,
				0,
			})
		} else {
			questionDTOs = append(questionDTOs, entity.QuestionDTO{questions[i],
				float32(questions[i].CorrectTrialNum) / float32(questions[i].TotalTrialNum),
				float32(questions[i].TotalTime) / float32(questions[i].TotalTrialNum),
			})
		}

	}
	return questionDTOs
}

//lowest time, highest time, lowest correctness, highest correctness
//func GetStatistics() string {
//	var questions []entity.Question
//	db.Debug().Find(&questions)
//	for i := 0; i < len(questions); i++ {
//
//
//	}
//}

// Get all questions with difficulty level
func GetAllQuestionsWithDifficultyLevel(difficulty string) []entity.Question {
	var questions []entity.Question
	//text := "text_test"
	//db.Debug().Find(&questions, "text", text)
	db.Debug().Find(&questions, "difficulty_level", difficulty)
	return questions
}

// store the result
func StoreResult(correct string, wrong string, time string) string {
	//fmt.Println(time[:len(time)-2] + "]")
	temp := time[:len(time)-2] + "]"
	var timeJson []TimeJson
	json.Unmarshal([]byte(temp), &timeJson)
	//fmt.Println(timeJson)
	for _, v := range timeJson {
		id, _ := strconv.Atoi(v.Id)
		time, _ := strconv.Atoi(v.ConsumingDuration)
		db.Model(&entity.Question{}).Where("id = ?", id).Update("total_time", gorm.Expr("total_time + ?", time))

	}
	tempCorrect := correct[:len(correct)-1]
	tempWrong := wrong[:len(wrong)-1]
	var correctArray []string
	if correct == "none" {
		correctArray = []string{}
	} else {
		correctArray = strings.Split(tempCorrect, ",")
	}
	var wrongArray []string
	if wrong == "none" {
		wrongArray = []string{}
	} else {
		wrongArray = strings.Split(tempWrong, ",")
	}

	//wrongArray = strings.Split(tempWrong, ",")
	var allArray []string = append(correctArray, wrongArray...)
	for _, v := range correctArray {
		id, _ := strconv.Atoi(v)
		db.Model(&entity.Question{}).Where("id = ?", id).Update("correct_trial_num", gorm.Expr("correct_trial_num + ?", 1))
	}
	var result []entity.Question
	var resultJsons []ReturnType
	fmt.Println(allArray)
	for _, v := range allArray {
		fmt.Println(allArray)
		//fmt.Println(v + "ppp")
		id, _ := strconv.Atoi(v)
		//fmt.Println(id)
		db.Model(&entity.Question{}).Where("id = ?", id).Update("total_trial_num", gorm.Expr("total_trial_num + ?", 1))
		fmt.Println("---------------------------------------")
		var question entity.Question
		//fmt.Println("id=" + string(id))
		db.First(&question, id)
		result = append(result, question)
		//fmt.Println(result)
		var resultJson ReturnType
		resultJson.Id = question.Id
		fmt.Println(string(question.TotalTrialNum) + "-------------------")
		if question.TotalTrialNum == 0 {
			//fmt.Println("jinqulemei")
			resultJson.AverageCorrectness = 0.25
			resultJson.AverageTime = 0
		} else {
			resultJson.AverageCorrectness = float32(question.CorrectTrialNum) / float32(question.TotalTrialNum)
			resultJson.AverageTime = float32(question.TotalTime) / float32(question.TotalTrialNum)
		}

		resultJsons = append(resultJsons, resultJson)
		//fmt.Println(resultJsons)
		//question = entity.Question{}
	}

	finalResult, _ := json.Marshal(resultJsons)
	//fmt.Println(result)

	return string(finalResult)
}

// getResult
func GetResult(ids string) []entity.Question {
	var strIdArray []string = strings.Split(ids, ",")
	var intIdArray []int
	for _, v := range strIdArray {
		id, _ := strconv.Atoi(v)
		intIdArray = append(intIdArray, id)
	}
	var questions []entity.Question
	db.Where("id IN ?", intIdArray).Find(&questions)
	return questions
}

// delete Question
func DeleteQuestion(id string) string {
	db.Delete(&entity.Question{}, id)
	return "The question with id=" + id + " success deleted!"
}

func CreateQuestion(question entity.Question) string {
	result := db.Create(&question)
	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return result.Error.Error()
	} else {
		return "created"
	}
}
