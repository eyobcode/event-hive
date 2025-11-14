package utils

import "github.com/gin-gonic/gin"

func GetCurrentUserID(c *gin.Context) (int, bool) {
	userId, exists := c.Get("user_id")
	if !exists {
		return 0, false
	}
	return userId.(int), false
}
