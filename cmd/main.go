package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Paramet02/multi-user-bookmark-api/internal/adapter/outbound/postgresql/connections"
	// "github.com/Paramet02/multi-user-bookmark-api/internal/domain"

	"github.com/Paramet02/multi-user-bookmark-api/internal/adapter/inbound/http/handlers"
	"github.com/Paramet02/multi-user-bookmark-api/internal/adapter/outbound/postgresql/repository"
	"github.com/Paramet02/multi-user-bookmark-api/internal/adapter/outbound/security"
	"github.com/Paramet02/multi-user-bookmark-api/internal/application/service"
)

func main() {
	app := fiber.New()

	// load file env
	err := godotenv.Load("./env/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// get value env
	dsn := os.Getenv("DATABASE_URL")
	db := connections.InitDatabasePostgares(dsn)
	if db == nil {
		panic("Failed to connect to the database")
	}

	Entropy := os.Getenv("MINENTROPY")
	intEntropy, err := strconv.Atoi(Entropy)
	if err != nil {
		fmt.Printf("Error converting environment variable to int: %v\n", err)
		return
	}

	// Initialize security & service layers
	hasher := security.NewBcryptHasher()
	policy := security.NewPolicy(float64(intEntropy))
	jwt := security.NewJWTManager("test", time.Hour*1)

	userRepository := repository.NewUserRepositoryDB(db)
	userService := service.NewUserService(userRepository, hasher, policy)
	userHandlers := handlers.NewUserHandler(userService, jwt)

	authService := service.NewAuthService(userRepository, hasher , jwt)
	authHandlers := handlers.NewAuthHandler(authService)
	middleware := handlers.NewAuthMiddleware(jwt)
	// collection := repository.NewCollectionRepositoryDB(db)

	// _ = collection // ป้องกัน unused variable error

	// input := &domain.Collection{
	// 	UserID: 1,
	// 	Name:   "My Collection",
	// 	CreatedAt: time.Now(),
	// 	UpdatedAt: time.Now(),
	// 	DeletedAt: nil,
	// }
	// //
	// err = collection.InsertCollection(context.Background() , input)
	// if err != nil {
	// 	fmt.Println("Error inserting collection:", err)
	// } else {
	// 	fmt.Println("Collection inserted successfully with ID:", input.ID)
	// }

	// ---------------- ROUTES ----------------

	api := app.Group("/api")

	// --- Auth routes ---
	auth := api.Group("/auth")
	auth.Post("/login", authHandlers.Login)
	auth.Post("/register", userHandlers.Register) 
	
	// --- User routes (ต้อง login ก่อน) ---
	users := api.Group("/users", middleware.Middleware() , middleware.RequestRole("user")) // middleware JWT
	users.Get("/id/:id" , userHandlers.GetUserID)
	
	log.Fatal(app.Listen(":8080"))
}

