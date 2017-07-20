package Database

import (
	"../Config"
	"../Models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Connect(config Config.Configuration) {
	db, err := gorm.Open("postgres", "user="+config.Database.UserName+" password="+config.Database.Password+" dbname="+config.Database.DBName+" sslmode=disable")
	if err != nil {
		panic(err)
	}
	// defer db.Close()
	var user Models.User
	db.AutoMigrate(&Models.User{})
	fmt.Println(db.Find(&user))
}

func Insert(u interface{}) {
	// Insert into table
	// TODO
	fmt.Println(u)
}
