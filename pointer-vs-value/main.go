package main

import (
	"fmt"
)

type User struct {
	Name         string
	Email        string
	TempPassword string
	PasswordHash string
	PasswordSalt []byte
	Bio          *string
	ImageLink    *string
}

type UserOut struct {
	User userOut0
}

type userOut0 struct {
	Email    string
	Token    string
	Username string
	Bio      string
	Image    *string
}

func (s *UserOut) FromModelP(user User, token string) UserOut {
	s.User = userOut0{
		Email:    user.Email,
		Token:    token,
		Username: user.Name,
		Bio:      *user.Bio,
		Image:    user.ImageLink,
	}

	return *s
}

func (s UserOut) FromModelV(user User, token string) UserOut {
	s.User = userOut0{
		Email:    user.Email,
		Token:    token,
		Username: user.Name,
		Bio:      *user.Bio,
		Image:    user.ImageLink,
	}

	return s
}

type IUserOutP interface {
	FromModelP(user User, token string) UserOut
}

func IFromModelP(s IUserOutP, user User, token string) UserOut {
	return s.FromModelP(user, token)
}

type IUserOutV interface {
	FromModelV(user User, token string) UserOut
}

func IFromModelV(s IUserOutV, user User, token string) UserOut {
	return s.FromModelV(user, token)
}

func main() {
	bio := "My bio."
	token := "JWT"
	user := User{
		Name:         "Paulo",
		Email:        "foo@bar.com",
		TempPassword: "tempxxx",
		PasswordHash: "9b8c34abcd",
		PasswordSalt: nil,
		Bio:          &bio,
		ImageLink:    nil,
	}

	printResults := func(exampleName string, base, result interface{}) {
		fmt.Println("***", exampleName, "***")
		fmt.Println("base:  ", base)
		fmt.Println("result:", result)
	}

	func() {
		exampleName := "Value receiver, value method"
		base := UserOut{}
		result := base.FromModelV(user, token)
		printResults(exampleName, base, result)
	}()

	func() {
		exampleName := "Pointer receiver, value method"
		base := UserOut{}
		result := (&base).FromModelV(user, token)
		printResults(exampleName, base, result)
	}()

	func() {
		exampleName := "Value receiver, pointer method"
		base := UserOut{}
		result := base.FromModelP(user, token)
		printResults(exampleName, base, result)
	}()

	func() {
		exampleName := "Pointer receiver, pointer method"
		base := UserOut{}
		result := (&base).FromModelP(user, token)
		printResults(exampleName, base, result)
	}()

	func() {
		exampleName := "Via interface: value receiver, value method"
		base := UserOut{}
		result := IFromModelV(base, user, token)
		printResults(exampleName, base, result)
	}()

	func() {
		exampleName := "Via interface: pointer receiver, value method"
		base := UserOut{}
		result := IFromModelV(&base, user, token)
		printResults(exampleName, base, result)
	}()

	func() {
		exampleName := "Via interface: value receiver, pointer method"
		base := UserOut{}
		//result := IFromModelP(base, user, token)  // doesn't compile
		result := "doesn't compile"
		printResults(exampleName, base, result)
	}()

	func() {
		exampleName := "Via interface: pointer receiver, pointer method"
		base := UserOut{}
		result := IFromModelP(&base, user, token)
		printResults(exampleName, base, result)
	}()
}
