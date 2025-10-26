package handlers

import (
    "golang-invoice/models"
    "net/http"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type ProductHandler struct {
    DB *gorm.DB
}

func NewProductHandler(db *gorm.DB) *ProductHandler {
    return &ProductHandler{DB: db}
}

func (h *ProductHandler) GetAll(c *gin.Context) {
    var products []models.Product
    if err := h.DB.Find(&products).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetByID(c *gin.Context) {
    id := c.Param("id")
    var product models.Product
    if err := h.DB.First(&product, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product tidak ditemukan"})
        return
    }
    c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) Create(c *gin.Context) {
    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.DB.Create(&product).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) Update(c *gin.Context) {
    id := c.Param("id")
    var product models.Product
    if err := h.DB.First(&product, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product tidak ditemukan"})
        return
    }

    var input models.Product
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    product.Name = input.Name
    product.Price = input.Price

    if err := h.DB.Save(&product).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) Delete(c *gin.Context) {
    id := c.Param("id")
    if err := h.DB.Delete(&models.Product{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Product berhasil dihapus"})
}
