package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"myGorm/example"
)

func main() {
	db, err := gorm.Open(sqlite.Open("identifier.sqlite"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	example.Run(db)
	//db.Create()

	//user := &example.TestUser{}
	//user.Name = "wyc1"
	//user.MemberNumber.Valid = true
	//db.Create(user)

}
