package main

import (
	"effective_mobile/internal/sub"
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
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_URL")),  &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&sub.Subscription{}) // Здесь добавляем структуры данных необходимые
	if err != nil {
        log.Printf("Ошибка при выполнении миграций: %v", err)
		return
    }
	log.Println("Миграции успешно выполнены")
}