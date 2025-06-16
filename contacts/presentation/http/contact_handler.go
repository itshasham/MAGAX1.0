package http

import (
	"errors"
	"internship/contacts/domain"
	"internship/contacts/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ContactDTO struct {
	Name    string `json:"name" binding:"required,min=1,max=50"`
	Email   string `json:"email" binding:"required,email,min=5,max=100"`
	Subject string `json:"subject" binding:"required,min=1,max=150"`
	Message string `json:"message" binding:"required,min=1,max=1000"`
}

type ContactHandler struct {
	UseCase usecase.ContactUsecase
	Mailer  usecase.Mailer
}

func NewContactHandler(useCase usecase.ContactUsecase, mailer usecase.Mailer) *ContactHandler {
	return &ContactHandler{
		UseCase: useCase,
		Mailer:  mailer,
	}
}

func (h *ContactHandler) CreateContact(c *gin.Context) {
	var input ContactDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errorsMap := make(map[string]string)
			for _, fe := range ve {
				errorsMap[fe.Field()] = fe.Tag() + " validation failed"
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Validation failed",
				"details": errorsMap,
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contact := domain.Contact{
		Name:    input.Name,
		Email:   input.Email,
		Subject: input.Subject,
		Message: input.Message,
	}

	result, err := h.UseCase.CreateContact(contact)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := h.Mailer.SendContactConfirmation(contact.Email, contact.Name, contact.Subject, contact.Message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Contact saved, but failed to send confirmation email",
		})
		return
	}

	c.JSON(http.StatusCreated, result)
}
