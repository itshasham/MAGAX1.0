package handler

import (
	model "internship/users/domain/models"
	"internship/users/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UseCase *usecase.UserUseCase
}

type ConnectWalletRequest struct {
	Wallet     string  `json:"wallet"`
	ReferralBy *string `json:"referral_by,omitempty"`
}

// 🔐 Wallet connect + referral sync + JWT issuance
func (h *UserHandler) ConnectWallet(c *gin.Context) {
	var req ConnectWalletRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}
	user, token, err := h.UseCase.ConnectWallet(req.Wallet, req.ReferralBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Wallet connect failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user, "token": token})
}

// 👤 Fetch user by ID
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User ID is required"})
		return
	}

	user, err := h.UseCase.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// ✏️ Update user profile fields
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Name           *string `json:"name"`
		Email          *string `json:"email"`
		PhoneNumber    *string `json:"phone_number"`
		SocialX        *string `json:"social_x"`
		SocialTelegram *string `json:"social_telegram"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid payload"})
		return
	}

	user := &model.User{
		ID:             model.UUIDFromString(id),
		Name:           body.Name,
		Email:          body.Email,
		PhoneNumber:    body.PhoneNumber,
		SocialX:        body.SocialX,
		SocialTelegram: body.SocialTelegram,
	}

	if err := h.UseCase.UpdateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":    user,
		"message": "Profile updated successfully",
		"status":  200,
	})
}

// 🏆 Leaderboard (Top 10 users by amount)
func (h *UserHandler) Leaderboard(c *gin.Context) {
	users, err := h.UseCase.GetTopUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get leaderboard"})
		return
	}
	c.JSON(http.StatusOK, users)
}
