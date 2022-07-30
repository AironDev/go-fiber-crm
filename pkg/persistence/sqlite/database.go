package sqlite

import (
	"fmt"
	"github.com/airondev/go-fiber-crm/pkg/application"
	"github.com/jinzhu/gorm"
)

var (
	DbConn *gorm.DB
	err    error
)

func InitDatabase() {
	DbConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Connection opened to database")
	DbConn.AutoMigrate(&application.Lead{})
	defer DbConn.Close()
}
