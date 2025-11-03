package main

import (
	"backend-base/internal/database"
	"backend-base/internal/model"
	"fmt"
)

func main() {
	db, err := database.ConnectPostgres(database.Config{
		Host:     "localhost",
		User:     "postgres",
		Password: "password",
		DBName:   "app_db",
		Port:     "5432",
	})
	if err != nil {
		panic(err)
	}

	admin := model.User{
		Name:     "Admin",
		Email:    "admin@example.com",
		Password: "$2a$10$Nw5qKe....", // pre-hashed bcrypt password
	}
	db.FirstOrCreate(&admin, model.User{Email: admin.Email})
	fmt.Println("✅ Seed completed")
}
