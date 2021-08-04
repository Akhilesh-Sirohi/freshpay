package routes

import (
	"github.com/freshpay/internal/controllers/payments"
	"github.com/gin-gonic/gin"
)
//SetupRouter ... Configure routes

func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/payments")
	{
		grp1.POST("", payments.AddPayment)
	}
	return r
}
