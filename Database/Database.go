package Database

import (
	"../Models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Connect() {
	db, err := gorm.Open("postgres", "user=postgres password=password dbname=test sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var user Models.User
	db.AutoMigrate(&Models.User{})
	fmt.Println(db.Find(&user))
}

func Insert(u interface{}) {
	// Insert into table
	// TODO
	fmt.Println(u)
}
