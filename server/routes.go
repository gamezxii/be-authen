package server

import (
	"be-authen/di"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, container *di.Container) {
	v1 := router.Group("/v1")
	{
		// Health check
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "OKaa",
			})
		})

		// User routes
		v1.GET("/users", container.UserHandler.GetUsers)
		v1.GET("/users/:id", container.UserHandler.GetUserDetail)
		v1.POST("/users", container.UserHandler.CreateUser)
		v1.PATCH("/users/:id/suspend", container.UserHandler.SuspendUser)
		v1.DELETE("/users/:id", container.UserHandler.SoftDeleteUser)

		// Operator routes
		v1.GET("/operators", container.OperatorHandler.GetAllOperatos)
		v1.POST("/operators", container.OperatorHandler.CreateOperator)
		v1.GET("/operators/:id", container.OperatorHandler.GetOperatorByID)
		v1.PUT("/operators/:id", container.OperatorHandler.UpdateOperator)
		v1.DELETE("/operators/:id", container.OperatorHandler.DeleteOperator)

		// OTP
		v1.GET("/otp/requests", container.OtpHandler.GetOtpRequests)
		// v1.GET("/otp/confirms", container.OtpHandler.GetOTPConfirms)
		v1.POST("/otp/request", container.OtpHandler.RequestOTP)
		v1.POST("/otp/verify", container.OtpHandler.VerifyOTP)

	}
}
