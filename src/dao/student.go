package dao

import (
	"fmt"
	"gorm_demo1/src/model"
	"time"
)

func CreateStudent() {
	stu := &model.Student{
		Name:     "xiaote33",
		Age:      18,
		Birthday: time.Now(),
	}
	DB.Create(stu)
}
func QueryFirstStudent() *model.Student {
	var stu model.Student
	DB.First(&stu)
	return &stu
}

func Find() {
	var stu model.Student
	DB.First(&stu, "name = ?", "xiaote32")
	fmt.Println(stu)
	var stus []model.Student
	DB.Where("name <> ?", "xiaote33").Find(&stus)
	for i, _ := range stus {
		fmt.Println(stus[i])
	}
}

func CreatTable() {
	DB.AutoMigrate(&model.Message{})
	DB.Create(&model.Message{})
	DB.AutoMigrate(&model.Community{})
	DB.Create(&model.Community{})
	DB.AutoMigrate(&model.Contact{})
	DB.Create(&model.Contact{})
	DB.AutoMigrate(&model.UserBasic{})
	DB.Create(&model.UserBasic{})
	DB.AutoMigrate(&model.GroupBasic{})
	DB.Create(&model.GroupBasic{})
}
