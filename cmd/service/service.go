package main

import (
    "context"
    "database/sql"
    "errors"
    "fmt"

    _ "github.com/jackc/pgx/v5/stdlib"
    "go-micro.dev/v5"
    "user-crud/internal/config"
)

// -------- Models ---------
type User struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

// -------- Requests / Responses ---------
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

// -------- Service ---------
type UserService struct {
    db *sql.DB
}

// -------- CRUD Methods ---------

func (s *UserService) CreateUser(ctx context.Context, req *CreateUserRequest, rsp *EmptyResponse) error {
    if req.User == nil || req.User.ID == "" {
        return errors.New("invalid user")
    }

    _, err := s.db.ExecContext(ctx,
        "INSERT INTO users (id, name, email) VALUES ($1, $2, $3)",
        req.User.ID, req.User.Name, req.User.Email,
    )
    return err
}

func (s *UserService) GetUser(ctx context.Context, req *GetUserRequest, rsp *GetUserResponse) error {
    row := s.db.QueryRowContext(ctx, "SELECT id, name, email FROM users WHERE id=$1", req.ID)

    user := &User{}
    if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return errors.New("user not found")
        }
        return err
    }
    rsp.User = user
    return nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *UpdateUserRequest, rsp *EmptyResponse) error {
    res, err := s.db.ExecContext(ctx,
        "UPDATE users SET name=$1, email=$2 WHERE id=$3",
        req.User.Name, req.User.Email, req.User.ID,
    )
    if err != nil {
        return err
    }

    count, _ := res.RowsAffected()
    if count == 0 {
        return errors.New("user not found")
    }
    return nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *DeleteUserRequest, rsp *EmptyResponse) error {
    res, err := s.db.ExecContext(ctx, "DELETE FROM users WHERE id=$1", req.ID)
    if err != nil {
        return err
    }
    count, _ := res.RowsAffected()
    if count == 0 {
        return errors.New("user not found")
    }
    return nil
}

// -------- Main ---------
func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	// Open database using config
	db, err := sql.Open(cfg.DBDriver, cfg.DatabaseDSN())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Create micro service
	service := micro.NewService(
		micro.Name(cfg.ServiceName),
	)
    service.Init()

    handler := &UserService{db: db}
    //Listens for incoming requests 
    //Receives RPC calls
    //Routes them to handler methods
    micro.RegisterHandler(service.Server(), handler)

    fmt.Println("User service running with PostgreSQL...")
    service.Run()
}

