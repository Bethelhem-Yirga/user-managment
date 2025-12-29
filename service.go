package main

import (
    "context"
    "errors"
    "fmt"
    "sync"

    "go-micro.dev/v5"
)

// --------- Models ---------

type User struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

// --------- Requests / Responses ---------

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

// --------- Service ---------

type UserService struct {
    mu    sync.Mutex
    users map[string]*User
}

// --------- CRUD METHODS ---------

func (s *UserService) CreateUser(ctx context.Context, req *CreateUserRequest, rsp *EmptyResponse) error {
    s.mu.Lock()
    defer s.mu.Unlock()

    if req.User == nil || req.User.ID == "" {
        return errors.New("invalid user")
    }

    s.users[req.User.ID] = req.User
    return nil
}

func (s *UserService) GetUser(ctx context.Context, req *GetUserRequest, rsp *GetUserResponse) error {
    s.mu.Lock()
    defer s.mu.Unlock()

    user, ok := s.users[req.ID]
    if !ok {
        return errors.New("user not found")
    }

    rsp.User = user
    return nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *UpdateUserRequest, rsp *EmptyResponse) error {
    s.mu.Lock()
    defer s.mu.Unlock()

    if _, ok := s.users[req.User.ID]; !ok {
        return errors.New("user not found")
    }

    s.users[req.User.ID] = req.User
    return nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *DeleteUserRequest, rsp *EmptyResponse) error {
    s.mu.Lock()
    defer s.mu.Unlock()

    delete(s.users, req.ID)
    return nil
}

// --------- MAIN ---------

func main() {
    service := micro.NewService(
        micro.Name("user.service"),
    )
    service.Init()

    handler := &UserService{
        users: make(map[string]*User),
    }

    micro.RegisterHandler(service.Server(), handler)

    fmt.Println("User service running...")
    service.Run()
}
