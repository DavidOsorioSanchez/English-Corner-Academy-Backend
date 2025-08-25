package main

import (
	"English-Corner-Academy-Gim/internal/database"

	"github.com/gin-gonic/gin"
)

// este es un archivo para una capa extra de seguridad,
// comprueba que sea igual en la base de datos que en el servicio que se use

// probablemente tenga que cambiar la database ya que no me parece tan escalable el sqlite3
// asi que probablemente tenga que reacer el archivo entero ಥ‿ಥ

func (app *application) GetUserFromContext(c *gin.Context) *database.User {
	contextUser, exist := c.Get("user")

	if !exist {
		return &database.User{}
	}

	user, ok := contextUser.(*database.User)

	if !ok {
		return &database.User{}
	}

	return user
}
