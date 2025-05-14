package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lubualo/gambit-user/models"
	"github.com/lubualo/gambit-user/tools"
)

func SignUp(signUpModel models.SignUp) error {
	fmt.Println("Sign up starts")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	query := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('" + signUpModel.UserEmail + "','" + signUpModel.UserUUID + "','" + tools.DateMySQL() + "')"
	fmt.Println(query)

	_, err = Db.Exec(query)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Sign up ended successfully")
	return nil
}
