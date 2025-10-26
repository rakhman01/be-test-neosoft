package main

import (
    "golang-invoice/database"
    "golang-invoice/models"
    "golang-invoice/routes"
    "os"
)

func main() {
    db := database.Connect()

    // migrasi tabel otomatis
    db.AutoMigrate(&models.Product{}, &models.Patient{}, &models.Invoice{}, &models.InvoiceItem{})

    router := routes.SetupRouter(db)
    
    // ambil PORT dari environment variable, default 8080 jika tidak ada
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    router.Run(":" + port)
}
