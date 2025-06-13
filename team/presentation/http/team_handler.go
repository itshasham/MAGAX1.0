package http

import (
	"net/http"
	"strconv"

	"internship/team/domain"
	"internship/team/usecase"

	"github.com/gin-gonic/gin"
)

type TeamHandler struct {
	UseCase usecase.TeamUsecase
}

func NewTeamHandler(useCase usecase.TeamUsecase) *TeamHandler {
	return &TeamHandler{UseCase: useCase}
}

func (h *TeamHandler) GetTeams(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	teams, count, err := h.UseCase.GetEnabledTeams(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": teams, "count": count})
}

func (h *TeamHandler) GetTeamByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	team, err := h.UseCase.GetTeamByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Team not found"})
		return
	}
	c.JSON(http.StatusOK, team)
}

func (h *TeamHandler) CreateTeam(c *gin.Context) {
	var team domain.Team
	if err := c.ShouldBindJSON(&team); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTeam, err := h.UseCase.CreateTeam(team)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newTeam)
}

func (h *TeamHandler) UpdateTeam(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var team domain.Team
	if err := c.ShouldBindJSON(&team); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedTeam, err := h.UseCase.UpdateTeam(id, team)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedTeam)
}

func (h *TeamHandler) DeleteTeam(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.UseCase.DeleteTeam(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}
