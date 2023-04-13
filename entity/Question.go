package entity

type Question struct {
	Id              int    `gorm:"column:id; PRIMARY_KEY"`
	Text            string `gorm:"column:text"`
	A               string `gorm:"column:A"`
	B               string `gorm:"column:B"`
	C               string `gorm:"column:C"`
	D               string `gorm:"column:D"`
	Answer          string `gorm:"column:answer"`
	DifficultyLevel string `gorm:"column:difficulty_level"`
	TotalTrialNum   int    `gorm:"column:total_trial_num"`
	CorrectTrialNum int    `gorm:"column:correct_trial_num"`
	TotalTime       int    `gorm:"column:total_time"`
	DetailSolution  string `gorm:"column:detail_solution"`
}

func (q Question) TableName() string {
	return "question"
}
