package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DomainBlackList struct {
	Domain string `gorm:"unique;size:253"` // 唯一索引
	gorm.Model
}

func (DomainBlackList) TableName() string {
	return "domain_black_list"
}

func init() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/linkme?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	//AutoMigrate
	err = db.AutoMigrate(
		&DomainBlackList{},
	)
	if err != nil {
		panic(err)
	}
}

func main() {

}
