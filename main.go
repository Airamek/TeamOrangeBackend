package main

import (
	"fmt"
	"main/dataProvider"
)

func main() {
	var provider = dataProvider.CreateProvider()
	provider.Init()

	users := provider.GetUsers()
	for _, user := range users {
		fmt.Print(user.GetDisplayName() + " " + user.GetMainEmail())
		for _, mail := range user.GetAliasEmails() {
			fmt.Print(" " + mail + " ")
		}
		fmt.Println()
	}
}
