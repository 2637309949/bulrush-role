/**
 * @author [Double]
 * @email [2637309949@qq.com.com]
 * @create date 2019-01-12 22:46:31
 * @modify date 2019-01-12 22:46:31
 * @desc [bulrush role plugin]
 */

package role

import (
	"net/http"
	"strings"

	"github.com/2637309949/bulrush"
	"github.com/gin-gonic/gin"
)

type (
	// Action for role
	Action struct {
		Roles   []string
		permits []string
	}
	// Role for role
	Role struct {
		bulrush.PNBase
		FailureHandler func(*gin.Context, string)
		RoleHandler    func(*gin.Context, string) bool
	}
)

// default failure handlerss
var defaultFailureHandler = func(c *gin.Context, action string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "This resource is not authorized",
	})
	c.Abort()
}

// TransformAction action tran
func TransformAction(action string) []Action {
	actStrs := strings.Split(action, ";")
	actions := make([]Action, 0)
	for _, act := range actStrs {
		rpStr := strings.Split(act, "@")
		if len(rpStr) == 1 {
			actions = append(actions, Action{Roles: strings.Split(rpStr[0], ",")})
		} else if len(rpStr) == 2 {
			actions = append(actions, Action{Roles: strings.Split(rpStr[0], ","), permits: strings.Split(rpStr[1], ",")})
		}
	}
	return actions
}

// Can do what
func (role *Role) Can(action string) gin.HandlerFunc {
	if role.FailureHandler == nil {
		role.FailureHandler = defaultFailureHandler
	}
	return func(c *gin.Context) {
		can := role.RoleHandler(c, action)
		if !can {
			role.FailureHandler(c, action)
			c.Abort()
		} else {
			c.Next()
		}
	}
}

// Plugin for role
func (role *Role) Plugin() bulrush.PNRet {
	return func() *Role {
		return role
	}
}
