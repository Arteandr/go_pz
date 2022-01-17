package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
	"strings"
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
	text := string(stdout)
	textArr := strings.Split(text, "\n")
	fmt.Println(textArr)
	c.JSON(http.StatusOK, gin.H{
		"msg": string(stdout),
	})
}
