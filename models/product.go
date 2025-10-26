package models

import (
    "fmt"
    "time"
    "gorm.io/gorm"
)

type Product struct {
    ID    uint   `gorm:"primaryKey" json:"id"`
    Code  string `gorm:"uniqueIndex;size:20" json:"code"`
    Name  string `json:"name"`
    Price int64  `json:"price"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) error {
    now := time.Now()
    year := now.Format("06")
    month := now.Format("01")
    
    var count int64
    prefix := fmt.Sprintf("P-%s%s", year, month)
    tx.Model(&Product{}).Where("code LIKE ?", prefix+"%").Count(&count)
    
    p.Code = fmt.Sprintf("P-%s%s%04d", year, month, count+1)
    return nil
}

