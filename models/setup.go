package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() {
	// mysql://doadmin:AVNS_om32aAj1SIl1t_Z3CYQ@db-mysql-blr1-10201-do-user-12928252-0.b.db.ondigitalocean.com:25060/defaultdb?ssl-mode=REQUIRED
	// doadmin:AVNS_om32aAj1SIl1t_Z3CYQ@(db-mysql-blr1-10201-do-user-12928252-0.b.db.ondigitalocean.com:25060)/notes?parseTime=true
	database, err := gorm.Open(mysql.Open("root:admin123@(127.0.0.1:3306)/notes?parseTime=true"), &gorm.Config{})
	if err != nil {
		// log.Fatal(err)
		panic("Failed to connect to database")
	}
	DB = database
}

func DBMigrate() {
	DB.AutoMigrate(&Note{})
	DB.AutoMigrate(&User{})
}
