package models

import (
    "fmt"
    "time"
    "gorm.io/gorm"
)

type Invoice struct {
    ID        uint          `gorm:"primaryKey" json:"id"`
    InvoiceNo string        `gorm:"uniqueIndex;size:30" json:"invoice_no"`
    Date      time.Time     `json:"date"`
    PatientID uint          `json:"patient_id"`
    Patient   Patient       `gorm:"foreignKey:PatientID" json:"patient"`
    Total     int64         `json:"total"`
    Items     []InvoiceItem `gorm:"foreignKey:InvoiceID" json:"items"`
}

func (i *Invoice) BeforeCreate(tx *gorm.DB) error {
    now := time.Now()
    i.Date = now
    year := now.Format("06")
    month := now.Format("01")
    day := now.Format("02")
    
    var count int64
    prefix := fmt.Sprintf("INV-%s%s%s", year, month, day)
    tx.Model(&Invoice{}).Where("invoice_no LIKE ?", prefix+"%").Count(&count)
    
    i.InvoiceNo = fmt.Sprintf("INV-%s%s%s%04d", year, month, day, count+1)
    return nil
}
