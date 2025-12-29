package main

import (
    "context"
    "fmt"
    "time"

    "go-micro.dev/v5"
    "go-micro.dev/v5/client"
)

type User struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

type CreateUserRequest struct {
    User *User `json:"user"`
}

type GetUserRequest struct {
    ID string `json:"id"`
}

type UpdateUserRequest struct {
    User *User `json:"user"`
}

type DeleteUserRequest struct {
    ID string `json:"id"`
}

type GetUserResponse struct {
    User *User `json:"user"`
}

func main() {
    service := micro.NewService(
        micro.Name("user.client"),
    )
    service.Init()

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // CREATE
    user := &User{ID: "1", Name: "Alice", Email: "alice@example.com"}
    createReq := client.NewRequest(
        "user.service",
        "UserService.CreateUser",
        &CreateUserRequest{User: user},
    )
    service.Client().Call(ctx, createReq, nil)

    // READ
    getRsp := &GetUserResponse{}
    getReq := client.NewRequest(
        "user.service",
        "UserService.GetUser",
        &GetUserRequest{ID: "1"},
    )
    service.Client().Call(ctx, getReq, getRsp)
    fmt.Println("GET:", getRsp.User)

    // UPDATE
    user.Name = "Alice Smith"
    updateReq := client.NewRequest(
        "user.service",
        "UserService.UpdateUser",
        &UpdateUserRequest{User: user},
    )
    service.Client().Call(ctx, updateReq, nil)

    // DELETE
    deleteReq := client.NewRequest(
        "user.service",
        "UserService.DeleteUser",
        &DeleteUserRequest{ID: "1"},
    )
    service.Client().Call(ctx, deleteReq, nil)

    fmt.Println("User deleted")
}
