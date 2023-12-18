package helpers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SessionsSet(c *gin.Context, userID uint64) {
	session := sessions.Default(c)
	var idInterface interface{} = &userID
	session.Set("id", idInterface)
	session.Save()
}

func SessionsGet(c *gin.Context) uint64 {
	sessions := sessions.Default(c)
	return sessions.Get("id").(uint64)
}

func SessionsClear(c *gin.Context) {
	sessions := sessions.Default(c)
	sessions.Clear()
	sessions.Save()
}
