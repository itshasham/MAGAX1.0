package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	// Mailer
	mailerPkg "internship/contacts/infrastructure/mailer"

	// Database
	db "internship/pkg/database_connection"

	// Repositories
	blogRepoPkg "internship/blog/infrastructure/persistence"
	contactRepoPkg "internship/contacts/infrastructure/persistence"
	couponRepoPkg "internship/coupon/infrastructure/persistence"
	faqRepoPkg "internship/faqs/infrastructure/persistence"
	refRepoPkg "internship/referral/infrastructure/persistence"
	teamRepoPkg "internship/team/infrastructure/persistence"
	userRepoPkg "internship/users/infrastructure/persistence"

	// Models for migration
	blogModels "internship/blog/domain/models"
	contactModels "internship/contacts/domain"
	couponModels "internship/coupon/domain/models"
	faqModels "internship/faqs/domain/models"
	refModels "internship/referral/domain/models"
	teamModels "internship/team/domain"
	userModels "internship/users/domain/models"

	// Usecases
	blogUsecase "internship/blog/usecase"
	contactUsecase "internship/contacts/usecase"
	couponUsecase "internship/coupon/usecase"
	faqUsecase "internship/faqs/usecase"
	teamUsecase "internship/team/usecase"
	userUsecase "internship/users/usecase"

	// Handlers
	blogHandlerPkg "internship/blog/presentation/http"
	contactHandlerPkg "internship/contacts/presentation/http"
	couponHandlerPkg "internship/coupon/presentation/http"
	faqHandlerPkg "internship/faqs/presentation/http"
	teamHandlerPkg "internship/team/presentation/http"
	userHandlerPkg "internship/users/presentation/http"

	// Routers
	blogRouter "internship/blog/presentation/router"
	contactRouter "internship/contacts/presentation/router"
	couponRouter "internship/coupon/presentation/router"
	faqRouter "internship/faqs/presentation/router"
	teamRouter "internship/team/presentation/router"
	userRouter "internship/users/presentation/router"
)

func main() {
	// Load .env
	if err := godotenv.Load(".env"); err != nil {
		if err := godotenv.Load("../.env"); err != nil {
			log.Fatal("❌ Failed to load .env file")
		}
	}

	database := db.ConnectPostgres()

	// DB Migrations
	log.Println("📦 Running DB migrations...")
	if err := database.AutoMigrate(
		&userModels.User{},
		&refModels.ReferralCodeModel{},
		&teamModels.Team{},
		&contactModels.Contact{},
		&faqModels.Faq{},
		&blogModels.Blog{},
		&couponModels.Coupon{},
	); err != nil {
		log.Fatalf("❌ AutoMigration failed: %v", err)
	}
	log.Println("✅ DB migrations complete")

	// Repositories & Usecases
	userRepo := userRepoPkg.NewUserRepository(database)
	referralRepo := refRepoPkg.NewReferralCodeRepo(database)
	teamRepo := teamRepoPkg.NewTeamRepository(database)
	contactRepo := contactRepoPkg.NewContactRepository(database)
	faqRepo := faqRepoPkg.NewFaqRepository(database)
	blogRepo := blogRepoPkg.NewBlogPGRepository(database)
	couponRepo := couponRepoPkg.NewCouponRepository(database) // ✅ Added

	userUseCase := &userUsecase.UserUseCase{UserRepo: userRepo, RefRepo: referralRepo}
	teamUseCase := teamUsecase.NewTeamUsecase(teamRepo)
	contactUseCase := contactUsecase.NewContactUsecase(contactRepo)
	faqUseCase := faqUsecase.NewFaqUsecase(faqRepo)
	blogUseCase := blogUsecase.NewBlogUsecase(blogRepo)
	couponUseCase := couponUsecase.NewCouponUsecase(couponRepo) // ✅ Added

	// Mailer
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := getEnvAsInt("SMTP_PORT", 587)
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")
	smtpFrom := os.Getenv("SMTP_FROM")
	contactMailer := mailerPkg.NewSMTPMailer(
		smtpHost, smtpPort, smtpUser, smtpPass, smtpFrom,
	)

	// Handlers
	userHandler := &userHandlerPkg.UserHandler{UseCase: userUseCase}
	teamHandler := teamHandlerPkg.NewTeamHandler(teamUseCase)
	contactHandler := contactHandlerPkg.NewContactHandler(contactUseCase, contactMailer)
	faqHandler := faqHandlerPkg.NewFaqHandler(faqUseCase)
	blogHandler := blogHandlerPkg.NewBlogHandler(blogUseCase)
	couponHandler := couponHandlerPkg.NewCouponHandler(couponUseCase) // ✅ Added

	// Router setup
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Register routes
	userRouter.RegisterUserRoutes(r, userHandler)
	teamRouter.RegisterTeamRoutes(r, teamHandler)
	contactRouter.RegisterContactRoutes(r, contactHandler)
	faqRouter.RegisterFaqRoutes(r, faqHandler)
	blogRouter.RegisterBlogRoutes(r, blogHandler)
	couponRouter.RegisterCouponRoutes(r, couponHandler) // ✅ Now correctly registered

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
