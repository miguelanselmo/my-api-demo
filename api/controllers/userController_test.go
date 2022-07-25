package controllers

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestController_CreateUser(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		ctrl *Controller
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ctrl.CreateUser(tt.args.c)
		})
	}
}

func TestController_GetUserById(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		ctrl *Controller
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ctrl.GetUserById(tt.args.c)
		})
	}
}

func TestController_GetUserAll(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		ctrl *Controller
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ctrl.GetUserAll(tt.args.c)
		})
	}
}
