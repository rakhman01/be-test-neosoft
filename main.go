package main

import (
    "golang-invoice/database"
    "golang-invoice/models"
    "golang-invoice/routes"
)

func main() {
    db := database.Connect()

    // migrasi tabel otomatis
    db.AutoMigrate(&models.Product{}, &models.Patient{}, &models.Invoice{}, &models.InvoiceItem{})

    router := routes.SetupRouter(db)
    router.Run(":8080") // buka di http://localhost:8080
}
