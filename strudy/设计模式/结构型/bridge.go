package main

import "fmt"

// 桥接模式：就是客户端类持有服务类的接口，controller和service直接就是一个桥接关系。形成一个抽象层和实施层的关系

type UserService interface {
	AddUser()
}

type UserController struct {
	service UserService
}

func NewUserController(service UserService) *UserController {
	return &UserController{service: service}
}

func (c UserController) AddUser() {
	c.service.AddUser()
}

type UserServiceImpl1 struct {
}

func (u UserServiceImpl1) AddUser() {
	fmt.Println("add user to window plat")
}

type UserServiceImpl2 struct {
}

func (u UserServiceImpl2) AddUser() {
	fmt.Println("add user to mac plat")
}

func main() {
	userController1 := NewUserController(new(UserServiceImpl1))
	userController2 := NewUserController(new(UserServiceImpl2))
	userController1.AddUser()
	userController2.AddUser()
}
