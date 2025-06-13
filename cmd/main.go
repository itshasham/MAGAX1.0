package main

import (
	"log"

	"github.com/gin-gonic/gin"

	// Database
	db "internship/pkg/database_connection"

	// Repositories
	refRepoPkg "internship/referral/infrastructure/persistence"
	teamRepoPkg "internship/team/infrastructure/persistence"
	userRepoPkg "internship/users/infrastructure/persistence"

	// Models for migration
	refModels "internship/referral/domain/models"
	teamModels "internship/team/domain"
	userModels "internship/users/domain/models"

	// Usecases & Handlers
	teamHandlerPkg "internship/team/presentation/http"
	userHandlerPkg "internship/users/presentation/http"

	teamRouter "internship/team/presentation/router"
	userRouter "internship/users/presentation/router"

	teamUsecase "internship/team/usecase"
	userUsecase "internship/users/usecase"
)

func main() {
	// Step 1: Connect to PostgreSQL
	database := db.ConnectPostgres()

	// Step 2: Auto Migrate Models
	log.Println("📦 Running DB migrations...")
	if err := database.AutoMigrate(
		&userModels.User{},
		&refModels.ReferralCodeModel{},
		&teamModels.Team{}, // Team model
	); err != nil {
		log.Fatalf("❌ AutoMigration failed: %v", err)
	}
	log.Println("✅ DB migrations complete")

	// Step 3: Initialize repositories
	userRepo := userRepoPkg.NewUserRepository(database)
	referralRepo := refRepoPkg.NewReferralCodeRepo(database)
	teamRepo := teamRepoPkg.NewTeamRepository(database)

	// Step 4: Initialize use cases
	userUseCase := &userUsecase.UserUseCase{
		UserRepo: userRepo,
		RefRepo:  referralRepo,
	}
	teamUseCase := teamUsecase.NewTeamUsecase(teamRepo)

	// Step 5: Initialize handlers
	userHandler := &userHandlerPkg.UserHandler{
		UseCase: userUseCase,
	}
	teamHandler := teamHandlerPkg.NewTeamHandler(teamUseCase)

	// Step 6: Initialize Gin router
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Health check
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Step 7: Register routes
	userRouter.RegisterUserRoutes(r, userHandler)
	teamRouter.RegisterTeamRoutes(r, teamHandler)

	// Step 8: Start HTTP server
	log.Println("🚀 Server running at http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
