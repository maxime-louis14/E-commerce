package helper

import (
	"errors"

	"github.com/gin-gonic/gin"
)

//CheckUserType renews the user tokens when they login
func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")
	err = nil
	if userType != role {
		err = errors.New("Unauthorized to access this resource")
		return err
	}

	return err
}

//MatchUserTypeToUid ne permet à l'utilisateur d'accéder qu'à ses données et à aucune autre donnée. Seul l'administrateur peut accéder à toutes les données de l'utilisateur
func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")
	err = nil

	if userType == "USER" && uid != userId {
		err = errors.New("Unauthorized to access this resource")
		return err
	}
	err = CheckUserType(c, userType)

	return err
}