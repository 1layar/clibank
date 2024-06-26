package cmd

import (
	"context"
	"fmt"

	"github.com/1layar/clibank/app/users"
	"github.com/1layar/clibank/platform"
)

type CmdDeleteUser struct {
	PhoneNumber string
}

func (c *CmdDeleteUser) Execute(ctx context.Context) error {
	service := ctx.Value(platform.UserServiceKey)
	fmt.Println("Masukan nomor HP : ")
	fmt.Scan(&c.PhoneNumber)
	var userService users.UserService = service.(users.UserService)
	err := userService.DeleteUser(c.PhoneNumber)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
