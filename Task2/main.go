package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type User struct {
	Name string
	Email string
	Age int
}

func main(){

	file, err := os.Open("text.csv")
	if err != nil {
		fmt.Println("Error While opening file:", err)
		return
	}
	
	reader := csv.NewReader(file) //reader
	Users, err := reader.ReadAll() 
	if err != nil {
		fmt.Println("Error While reading CSV file:", err)
		return
	}

	
	var AllUsers []User
	for i, user := range Users {
		if i == 0 {
			continue
		}
		age, err := strconv.Atoi(user[2])
		if err != nil {
			fmt.Println("Error While parsing Age:", err)
			return
		}

		AllUsers = append(AllUsers, User{
			Name: user[0],
			Email: user[1],
			Age: age,
		})
	}
	fmt.Println("Name\tEmail\tAge")
	fmt.Println("-----------------------------------------")
	for _, user := range AllUsers {
		fmt.Printf("%s\t%s\t%d\n", user.Name, user.Email, user.Age)
	}
	
	defer file.Close()
}
