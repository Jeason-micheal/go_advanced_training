package main

import (
	"fmt"
	"sql/dao"
	"sql/dao/table"
)

func main() {
	var id int64 = 1
	user, err := dao.GetUserById(id)
	if err != nil {
		fmt.Printf("GetUserById(%d) err: %+v\n", id, err)
	} else {
		fmt.Printf("%+v\n", user)
	}

	PrintAllUsers()
	
	u := table.User{
		Username: "Amy",
		Password: "aa111",
		Status:   1,
		Phone:    "19876",
	}
	err = dao.AddUser(&u)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	PrintAllUsers()
}

func PrintAllUsers() {
	users, err := dao.GetAllUsers()
	if err != nil {
		fmt.Printf("")
	} else {
		for _, t := range users {
			fmt.Printf("%+v\n", t)
		}
	}
}
