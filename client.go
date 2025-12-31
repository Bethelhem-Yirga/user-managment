package main

import (
	"context"
	"fmt"

	"go-micro.dev/v5"
	"go-micro.dev/v5/client"
)

// -------------------- Models --------------------
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// -------------------- Request structs --------------------
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

type EmptyResponse struct{}

// -------------------- Main --------------------
func main() {
	// Initialize Go Micro client service
	service := micro.NewService(
		micro.Name("user.client"),
	)
	service.Init()

	c := service.Client()

	// ---------------- Create User ----------------
	user := &User{ID: "1", Name: "Alice", Email: "alice@example.com"}
	createReq := &CreateUserRequest{User: user}
	createRPC := client.NewRequest("user.service", "UserService.CreateUser", createReq)

	if err := c.Call(context.Background(), createRPC, &EmptyResponse{}); err != nil {
		fmt.Println("Error creating user:", err)
		return
	}
	fmt.Println("User created successfully!")

	// ---------------- Read User ----------------
	getRPC := client.NewRequest("user.service", "UserService.GetUser", &GetUserRequest{ID: "1"})
	var getResp GetUserResponse
	if err := c.Call(context.Background(), getRPC, &getResp); err != nil {
		fmt.Println("Error fetching user:", err)
		return
	}
	fmt.Printf("Retrieved User: %+v\n", getResp.User)

	// ---------------- Update User ----------------
	user.Name = "Alice Smith"
	updateRPC := client.NewRequest("user.service", "UserService.UpdateUser", &UpdateUserRequest{User: user})
	if err := c.Call(context.Background(), updateRPC, &EmptyResponse{}); err != nil {
		fmt.Println("Error updating user:", err)
		return
	}
	fmt.Println("User updated successfully!")

	// ---------------- Verify Update ----------------
	if err := c.Call(context.Background(), getRPC, &getResp); err != nil {
		fmt.Println("Error fetching user:", err)
		return
	}
	fmt.Printf("Updated User: %+v\n", getResp.User)

	// ---------------- Delete User ----------------
	deleteRPC := client.NewRequest("user.service", "UserService.DeleteUser", &DeleteUserRequest{ID: "1"})
	if err := c.Call(context.Background(), deleteRPC, &EmptyResponse{}); err != nil {
		fmt.Println("Error deleting user:", err)
		return
	}
	fmt.Println("User deleted successfully!")

	// ---------------- Verify Deletion ----------------
	if err := c.Call(context.Background(), getRPC, &getResp); err != nil {
		fmt.Println("User not found (as expected):", err)
	} else {
		fmt.Println("Error: user still exists:", getResp.User)
	}
}
