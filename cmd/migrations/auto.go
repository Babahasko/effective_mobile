package main

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(`.env`)
	if err != nil {
		panic(err)
	}
	_ , err = gorm.Open(postgres.Open(os.Getenv("DSN")),  &gorm.Config{
	})
	if err != nil {
		panic(err)
	}
	// err = db.AutoMigrate(&link.Link{}, &user.User{}, &stat.Stat{}) // Здесь добавляем структуры данных необходимые
	// if err != nil {
    //     log.Printf("Ошибка при выполнении миграций: %v", err)
	// 	return
    // }
	log.Println("Миграции успешно выполнены")
}