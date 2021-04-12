package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	UserID    string `json:"userid"`
	StartDate string `json:"startdate"`
	EndDate   string `json:"enddate"`
	Title     string `json:"title"`
	Status    string `json:"status"`
}

type User struct {
	gorm.Model
	UserID   string `json:"userid"`
	Password string `json:"password"`
}

func main() {
	e := echo.New()

	err := e.PUT("/todo", insertTodo)
	if err != nil {
		fmt.Println(err)
	}

	err = e.PATCH("/todo", updateTodo)
	if err != nil {
		fmt.Println(err)
	}

	err = e.DELETE("/todo", deleteTodo)
	if err != nil {
		fmt.Println(err)
	}

	err = e.GET("/todo", selectTodo)
	if err != nil {
		fmt.Println(err)
	}

	err = e.PUT("/user", insertUser)
	if err != nil {
		fmt.Println(err)
	}

	err = e.PATCH("/user", updateUser)
	if err != nil {
		fmt.Println(err)
	}

	err = e.DELETE("/user", deleteUser)
	if err != nil {
		fmt.Println(err)
	}

	err = e.GET("/user", selectUser)
	if err != nil {
		fmt.Println(err)
	}

	e.Logger.Fatal(e.Start(":1323"))
}

func insertTodo(c echo.Context) error {

	db, err := gorm.Open(mysql.Open("root:1234@tcp(localhost:3306)/todo?parseTime=true"), &gorm.Config{})
	if err != nil {
		return err
	}

	todo := Todo{
		UserID:    "1",
		StartDate: "2",
		EndDate:   "3",
		Title:     "4",
		Status:    "5",
	}
	db.AutoMigrate(&todo)

	db.Create(&todo)

	return c.String(http.StatusOK, "inserted")
}

func insertUser(c echo.Context) error {

	db, err := gorm.Open(mysql.Open("root:1234@tcp(localhost:3306)/todo?parseTime=true"), &gorm.Config{})
	if err != nil {
		return err
	}

	user := User{
		UserID:   "1",
		Password: "2",
	}
	db.AutoMigrate(&user)

	db.Create(&user)

	return c.String(http.StatusOK, "inserted")
}

func updateTodo(c echo.Context) error {
	db, err := gorm.Open(mysql.Open("root:1234@tcp(localhost:3306)/todo?parseTime=true"), &gorm.Config{})
	if err != nil {
		return err
	}

	var todo Todo

	db.Model(&todo).Where("user_id=?", "1").Update("title", "Updated").Update("status", "Updated")

	return c.String(http.StatusOK, "updated")
}

func updateUser(c echo.Context) error {
	db, err := gorm.Open(mysql.Open("root:1234@tcp(localhost:3306)/todo?parseTime=true"), &gorm.Config{})
	if err != nil {
		return err
	}

	var user User

	db.Model(&user).Where("user_id=?", "1").Update("password", "Updated")

	return c.String(http.StatusOK, "updated")
}

func deleteTodo(c echo.Context) error {
	db, err := gorm.Open(mysql.Open("root:1234@tcp(localhost:3306)/todo?parseTime=true"), &gorm.Config{})
	if err != nil {
		return err
	}

	var todo Todo

	db.Where("user_id=?", "1").Unscoped().Delete(&todo)

	return c.String(http.StatusOK, "deleted")
}

func deleteUser(c echo.Context) error {
	db, err := gorm.Open(mysql.Open("root:1234@tcp(localhost:3306)/todo?parseTime=true"), &gorm.Config{})
	if err != nil {
		return err
	}

	var user User

	db.Where("user_id=?", "1").Unscoped().Delete(&user)

	return c.String(http.StatusOK, "deleted")
}

func selectTodo(c echo.Context) error {
	db, err := gorm.Open(mysql.Open("root:1234@tcp(localhost:3306)/todo?parseTime=true"), &gorm.Config{})
	if err != nil {
		return err
	}

	var todos []Todo

	err = db.Table("todos").Where("user_id=?", "1").Scan(&todos).Error
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, todos)
}

func selectUser(c echo.Context) error {
	db, err := gorm.Open(mysql.Open("root:1234@tcp(localhost:3306)/todo?parseTime=true"), &gorm.Config{})
	if err != nil {
		return err
	}

	var users []User

	err = db.Table("users").Where("user_id=?", "1").Scan(&users).Error
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}
