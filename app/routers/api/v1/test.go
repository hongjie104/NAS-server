package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjie104/NAS-server/app/pkg/e"
	"github.com/hongjie104/NAS-server/app/pkg/log"
	"github.com/hongjie104/NAS-server/app/pkg/utils"
	"github.com/hongjie104/NAS-server/app/routers/api"
)

// Test a
func Test(c *gin.Context) {
	// token, _ := utils.GenerateToken("tom", "123")
	// c.JSON(http.StatusOK, gin.H{"test": true, "data": token})
	// id := com.StrTo(c.Param("id")).MustInt()
	id := 3630
	response := api.Gin{C: c}
	t, err := utils.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRvbSIsInBhc3N3b3JkIjoiMTIzIiwiZXhwIjoxNTY0MjI3NzQ1LCJpc3MiOiJsZWFybi1nbyJ9.67I3ccNeCPRheTc-YoUnPeZwTXNb6u2d8dLIqktAQT0")
	if err == nil {
		response.Success(gin.H{"test": true, "data": t.ID})
	} else {
		log.LogError(err)
		// c.JSON(http.StatusOK, gin.H{"success": false, "data": nil})
		log.LogInfo(id)
		response.Fail(e.Error)
	}
}
