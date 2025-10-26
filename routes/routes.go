package routes

import (
    "golang-invoice/handlers"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "github.com/gin-contrib/cors"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
    r := gin.Default()

    //handle cors
        r.Use(cors.New(cors.Config{
        AllowOrigins: []string{"http://localhost:5173"},
        AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
    }))


    productHandler := handlers.NewProductHandler(db)
    patientHandler := handlers.NewPatientHandler(db)
    invoiceHandler := handlers.NewInvoiceHandler(db)

    api := r.Group("/api")
    {
        products := api.Group("/products")
        {
            products.GET("", productHandler.GetAll)
            products.GET("/:id", productHandler.GetByID)
            products.POST("", productHandler.Create)
            products.PUT("/:id", productHandler.Update)
            products.DELETE("/:id", productHandler.Delete)
        }

        patients := api.Group("/patients")
        {
            patients.GET("", patientHandler.GetAll)
            patients.GET("/:id", patientHandler.GetByID)
            patients.POST("", patientHandler.Create)
            patients.PUT("/:id", patientHandler.Update)
            patients.DELETE("/:id", patientHandler.Delete)
        }

        invoices := api.Group("/invoices")
        {
            invoices.GET("", invoiceHandler.GetAll)
            invoices.GET("/:id", invoiceHandler.GetByID)
            invoices.GET("/no/:invoiceNo", invoiceHandler.GetByInvoiceNo)
            invoices.POST("", invoiceHandler.Create)
        }
    }

    return r
}
