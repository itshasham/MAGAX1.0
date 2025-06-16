package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	mailerPkg "internship/contacts/infrastructure/mailer"

	// Database
	db "internship/pkg/database_connection"

	// Repositories
	contactRepoPkg "internship/contacts/infrastructure/persistence"
	faqRepoPkg "internship/faqs/infrastructure/persistence"
	refRepoPkg "internship/referral/infrastructure/persistence"
	teamRepoPkg "internship/team/infrastructure/persistence"
	userRepoPkg "internship/users/infrastructure/persistence"

	// Models for migration
	contactModels "internship/contacts/domain"
	faqModels "internship/faqs/domain/models"
	refModels "internship/referral/domain/models"
	teamModels "internship/team/domain"
	userModels "internship/users/domain/models"

	// Usecases & Handlers
	contactHandlerPkg "internship/contacts/presentation/http"
	faqHandlerPkg "internship/faqs/presentation/http"
	teamHandlerPkg "internship/team/presentation/http"
	userHandlerPkg "internship/users/presentation/http"

	// Routers
	contactRouter "internship/contacts/presentation/router"
	faqRouter "internship/faqs/presentation/router"
	teamRouter "internship/team/presentation/router"
	userRouter "internship/users/presentation/router"

	// Usecases
	contactUsecase "internship/contacts/usecase"
	faqUsecase "internship/faqs/usecase"
	teamUsecase "internship/team/usecase"
	userUsecase "internship/users/usecase"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		if err := godotenv.Load("../.env"); err != nil {
			log.Fatal("❌ Failed to load .env file")
		}
	}

	database := db.ConnectPostgres()

	log.Println("📦 Running DB migrations...")
	if err := database.AutoMigrate(
		&userModels.User{},
		&refModels.ReferralCodeModel{},
		&teamModels.Team{},
		&contactModels.Contact{},
		&faqModels.Faq{},
	); err != nil {
		log.Fatalf("❌ AutoMigration failed: %v", err)
	}
	log.Println("✅ DB migrations complete")

	userRepo := userRepoPkg.NewUserRepository(database)
	referralRepo := refRepoPkg.NewReferralCodeRepo(database)
	teamRepo := teamRepoPkg.NewTeamRepository(database)
	contactRepo := contactRepoPkg.NewContactRepository(database)
	faqRepo := faqRepoPkg.NewFaqRepository(database)

	userUseCase := &userUsecase.UserUseCase{
		UserRepo: userRepo,
		RefRepo:  referralRepo,
	}
	teamUseCase := teamUsecase.NewTeamUsecase(teamRepo)
	contactUseCase := contactUsecase.NewContactUsecase(contactRepo)
	faqUseCase := faqUsecase.NewFaqUsecase(faqRepo)

	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := getEnvAsInt("SMTP_PORT", 587)
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")
	smtpFrom := os.Getenv("SMTP_FROM")

	fmt.Println("🛠 SMTP config:")
	fmt.Println("Host:", smtpHost)
	fmt.Println("Port:", smtpPort)
	fmt.Println("User:", smtpUser)
	fmt.Println("From:", smtpFrom)

	contactMailer := mailerPkg.NewSMTPMailer(
		smtpHost,
		smtpPort,
		smtpUser,
		smtpPass,
		smtpFrom,
	)

	userHandler := &userHandlerPkg.UserHandler{UseCase: userUseCase}
	teamHandler := teamHandlerPkg.NewTeamHandler(teamUseCase)
	contactHandler := contactHandlerPkg.NewContactHandler(contactUseCase, contactMailer)
	faqHandler := faqHandlerPkg.NewFaqHandler(faqUseCase)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	userRouter.RegisterUserRoutes(r, userHandler)
	teamRouter.RegisterTeamRoutes(r, teamHandler)
	contactRouter.RegisterContactRoutes(r, contactHandler)
	faqRouter.RegisterFaqRoutes(r, faqHandler)

	log.Println("🚀 Server running at http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}

func getEnvAsInt(key string, defaultVal int) int {
	if val := os.Getenv(key); val != "" {
		var i int
		fmt.Sscanf(val, "%d", &i)
		return i
	}
	return defaultVal
}
