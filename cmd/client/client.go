package main

import (
	"context"
	"fmt"

	"go-micro.dev/v5"
	"go-micro.dev/v5/client"
	"user-crud/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

	
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

	cfg, err := config.LoadConfig()
if err != nil {
    fmt.Println("Error loading config:", err)
    return
}

	// Initialize Go Micro client service
	service := micro.NewService(
		micro.Name("user.client"),
	)
	service.Init()

	c := service.Client()
	r := gin.Default()
	r.Use(cors.New(cors.Config{
    AllowOrigins:     []string{cfg.FrontendURL},
    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
    AllowCredentials: true,
}))

	// ---------------- Create User ----------------
		r.POST("/users", func(ctx *gin.Context) {
		var user map[string]interface{}
		ctx.BindJSON(&user)

		req := client.NewRequest(
			cfg.ServiceName,
			"UserService.CreateUser",
			map[string]interface{}{
				"user": user,
			},
		)
		var resp interface{}
		if err := c.Call(context.Background(), req, &resp); err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, resp)
	})


	// ---------------- Read User ----------------
		r.GET("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		req := client.NewRequest(
			cfg.ServiceName,
			"UserService.GetUser",
			map[string]string{"id": id},
		)

		var resp interface{}
		if err := c.Call(context.Background(), req, &resp); err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, resp)
	})

	// ---------------- Update User ----------------
		r.PUT("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		var user User
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		user.ID = id

		req := client.NewRequest(
			cfg.ServiceName,
			"UserService.UpdateUser",
			map[string]interface{}{
				"user": user,
			},
		)

		var resp interface{}
		if err := c.Call(context.Background(), req, &resp); err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, resp)
	})


	// ---------------- Delete User ----------------
		r.DELETE("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		
		req := client.NewRequest(
			cfg.ServiceName,
			"UserService.DeleteUser",
			map[string]string{"id": id},
		)
		
		var resp interface{}
		if err := c.Call(context.Background(), req, &resp); err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, resp)
	})

		r.Run(":8080")
}
