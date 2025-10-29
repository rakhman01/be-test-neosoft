package handlers

import (
    "golang-invoice/models"
    "golang-invoice/utils"
    "net/http"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type InvoiceHandler struct {
    DB *gorm.DB
}

func NewInvoiceHandler(db *gorm.DB) *InvoiceHandler {
    return &InvoiceHandler{DB: db}
}

type CreateInvoiceRequest struct {
    PatientID uint `json:"patient_id" binding:"required"`
    Items     []struct {
        ProductID uint `json:"product_id" binding:"required"`
        Quantity  int  `json:"quantity" binding:"required,min=1"`
    } `json:"items" binding:"required,min=1"`
}

func (h *InvoiceHandler) Create(c *gin.Context) {
    var req CreateInvoiceRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var patient models.Patient
    if err := h.DB.First(&patient, req.PatientID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Patient tidak ditemukan"})
        return
    }

    invoice := models.Invoice{
        PatientID: req.PatientID,
    }

    var total int64 = 0
    var items []models.InvoiceItem

    for _, item := range req.Items {
        var product models.Product
        if err := h.DB.First(&product, item.ProductID).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Product ID " + err.Error()})
            return
        }

        subtotal := product.Price * int64(item.Quantity)
        total += subtotal

        items = append(items, models.InvoiceItem{
            ProductID: product.ID,
            Quantity:  item.Quantity,
            Price:     product.Price,
            Subtotal:  subtotal,
        })
    }

    invoice.Total = total
    invoice.Items = items

    if err := h.DB.Create(&invoice).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    h.DB.Preload("Patient").Preload("Items.Product").First(&invoice, invoice.ID)

    c.JSON(http.StatusCreated, invoice)
}

func (h *InvoiceHandler) GetAll(c *gin.Context) {
    var invoices []models.Invoice
    if err := h.DB.Preload("Patient").Preload("Items.Product").Find(&invoices).Error; err != nil {
      utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
		return

    }
  utils.SuccessResponse(c, "Data retrieved successfully", invoices)

}

func (h *InvoiceHandler) GetByID(c *gin.Context) {
    id := c.Param("id")
    var invoice models.Invoice
    if err := h.DB.Preload("Patient").Preload("Items.Product").First(&invoice, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Invoice tidak ditemukan"})
        return
    }
    c.JSON(http.StatusOK, invoice)
}

func (h *InvoiceHandler) GetByInvoiceNo(c *gin.Context) {
    invoiceNo := c.Param("invoiceNo")
    var invoice models.Invoice
    if err := h.DB.Preload("Patient").Preload("Items.Product").Where("invoice_no = ?", invoiceNo).First(&invoice).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Invoice tidak ditemukan"})
        return
    }
    c.JSON(http.StatusOK, invoice)
}

func (h *InvoiceHandler) Delete(c *gin.Context) {
    id := c.Param("id")
    var invoice models.Invoice
    if err := h.DB.First(&invoice, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Invoice tidak ditemukan"})
        return
    }

    if err := h.DB.Delete(&invoice).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Invoice berhasil dihapus"})
}
