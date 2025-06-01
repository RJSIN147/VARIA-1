package controllers

import (
	"backend/utils"
	// "database/sql"
	"encoding/json"
	"net/http"
	// "os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct{}

func NewAuthController() *AuthController {
	return &AuthController{}
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func (ac *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Basic validation
	req.Email = strings.TrimSpace(req.Email)
	req.Name = strings.TrimSpace(req.Name)
	req.Phone = strings.TrimSpace(req.Phone)
	if req.Email == "" || req.Name == "" || req.Phone == "" || req.Password == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	db, err := utils.InitDB()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Check if email already exists
	var exists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)", req.Email).Scan(&exists)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	if exists {
		http.Error(w, "Email already registered", http.StatusConflict)
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error processing password", http.StatusInternalServerError)
		return
	}

	// Insert user
	_, err = db.Exec(
		"INSERT INTO users (name, email, phone, password_hash) VALUES ($1, $2, $3, $4)",
		req.Name, req.Email, req.Phone, string(hashedPassword),
	)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, map[string]string{"message": "Registration successful"})
}

func (ac *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	// Handle login logic here
	render.JSON(w, r, map[string]string{"message": "Login successful"})
}

func (ac *AuthController) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/login", ac.Login)
	r.Post("/register", ac.Register)
	return r
}