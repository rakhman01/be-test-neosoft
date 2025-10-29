package handlers

import (
    "golang-invoice/models"
    "net/http"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type PatientHandler struct {
    DB *gorm.DB
}

func NewPatientHandler(db *gorm.DB) *PatientHandler {
    return &PatientHandler{DB: db}
}

func (h *PatientHandler) GetAll(c *gin.Context) {
    var patients []models.Patient
    if err := h.DB.Find(&patients).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, patients)
}

func (h *PatientHandler) GetByID(c *gin.Context) {
    id := c.Param("id")
    var patient models.Patient
    if err := h.DB.First(&patient, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Patient tidak ditemukan"})
        return
    }
    c.JSON(http.StatusOK, patient)
}

func (h *PatientHandler) Create(c *gin.Context) {
    var patient models.Patient
    if err := c.ShouldBindJSON(&patient); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.DB.Create(&patient).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, patient)
}

func (h *PatientHandler) Update(c *gin.Context) {
    id := c.Param("id")
    var patient models.Patient
    if err := h.DB.First(&patient, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Patient tidak ditemukan"})
        return
    }

    var input models.Patient
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    patient.Name = input.Name
    patient.Phone = input.Phone

    if err := h.DB.Save(&patient).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, patient)
}

func (h *PatientHandler) Delete(c *gin.Context) {
    id := c.Param("id")
    var patient models.Patient
    if err := h.DB.First(&patient, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Patient tidak ditemukan"})
        return
    }

    var invoiceCount int64
    if err := h.DB.Model(&models.Invoice{}).Where("patient_id = ?", id).Count(&invoiceCount).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if invoiceCount > 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Patient tidak dapat dihapus karena masih memiliki invoice"})
        return
    }

    if err := h.DB.Delete(&patient).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Patient berhasil dihapus"})
}
