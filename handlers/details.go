package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
)

func Details(c *gin.Context) {
	localCmd := exec.Command(CMD, "details")
	stdout, err := localCmd.Output()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println(string(stdout))
	c.JSON(http.StatusOK, gin.H{
		"msg": string(stdout),
	})
}
