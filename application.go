// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package role

import (
	"net/http"

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
		FailureHandler func(*gin.Context, string)
		RoleHandler    func(*gin.Context, string) bool
		RoleHandler1   func(*gin.Context, string) bool
		RoleHandler2   func(*gin.Context, string) bool
		RoleHandler3   func(*gin.Context, string) bool
		RoleHandler4   func(*gin.Context, string) bool
		RoleHandler5   func(*gin.Context, string) bool
		RoleHandler6   func(*gin.Context, string) bool
	}
)

// default failure handlerss
var defaultFailureHandler = func(c *gin.Context, action string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "Internal Server Error",
		"stack":   "this resource is not authorized",
	})
	c.Abort()
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

// Can1 do what
func (role *Role) Can1(action string) gin.HandlerFunc {
	if role.FailureHandler == nil {
		role.FailureHandler = defaultFailureHandler
	}
	return func(c *gin.Context) {
		can := role.RoleHandler1(c, action)
		if !can {
			role.FailureHandler(c, action)
			c.Abort()
		} else {
			c.Next()
		}
	}
}

// Can2 do what
func (role *Role) Can2(action string) gin.HandlerFunc {
	if role.FailureHandler == nil {
		role.FailureHandler = defaultFailureHandler
	}
	return func(c *gin.Context) {
		can := role.RoleHandler2(c, action)
		if !can {
			role.FailureHandler(c, action)
			c.Abort()
		} else {
			c.Next()
		}
	}
}

// Can3 do what
func (role *Role) Can3(action string) gin.HandlerFunc {
	if role.FailureHandler == nil {
		role.FailureHandler = defaultFailureHandler
	}
	return func(c *gin.Context) {
		can := role.RoleHandler3(c, action)
		if !can {
			role.FailureHandler(c, action)
			c.Abort()
		} else {
			c.Next()
		}
	}
}

// Can4 do what
func (role *Role) Can4(action string) gin.HandlerFunc {
	if role.FailureHandler == nil {
		role.FailureHandler = defaultFailureHandler
	}
	return func(c *gin.Context) {
		can := role.RoleHandler4(c, action)
		if !can {
			role.FailureHandler(c, action)
			c.Abort()
		} else {
			c.Next()
		}
	}
}

// Can5 do what
func (role *Role) Can5(action string) gin.HandlerFunc {
	if role.FailureHandler == nil {
		role.FailureHandler = defaultFailureHandler
	}
	return func(c *gin.Context) {
		can := role.RoleHandler5(c, action)
		if !can {
			role.FailureHandler(c, action)
			c.Abort()
		} else {
			c.Next()
		}
	}
}

// Can6 do what
func (role *Role) Can6(action string) gin.HandlerFunc {
	if role.FailureHandler == nil {
		role.FailureHandler = defaultFailureHandler
	}
	return func(c *gin.Context) {
		can := role.RoleHandler6(c, action)
		if !can {
			role.FailureHandler(c, action)
			c.Abort()
		} else {
			c.Next()
		}
	}
}

// Plugin for role
func (role *Role) Plugin() *Role {
	return role
}
