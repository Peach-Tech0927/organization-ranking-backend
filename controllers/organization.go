package controllers
import (
	"net/http"
	"organization-ranking-backend/models"
	"github.com/gin-gonic/gin"
)
type registerOrganization struct {
	Name string `json:"name"`

}


func RegisterOrganization(c *gin.Context){
	var input registerOrganization
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var organization models.Organization
	organization.Name = input.Name

	err = organization.SaveToDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"organization": organization.ToJSONResponse()})
}