package data
//
//import (
//	"github.com/jinzhu/gorm"
//	_ "github.com/jinzhu/gorm/dialects/postgres"
//	"fmt"
//	"log"
//)
//
//type Database interface {
//	InitDB()
//	InitSchema()
//}
//
//type User struct {
//	ID     string `json:"id" form:"-"`
//	Username string `json:"username" form:"username"`
//	Password string `json:"password" form:"password"`
//}
//
//type Impl struct {
//	DB *gorm.DB
//}
//
//func (i *Impl) InitDB() {
//	var err error
//	i.DB, err = gorm.Open("postgres", "host=52.59.245.228:5432 user=sqdron dbname=sqdron sslmode=disable password=sqdron")
//	if err != nil {
//		log.Fatalf("Got error when connect database, the error is '%v'", err)
//	}
//	i.DB.LogMode(true)
//}
//
//func (i *Impl) InitSchema() {
//	i.DB.AutoMigrate(&User{})
//}