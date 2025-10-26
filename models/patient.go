package models

import (
    "fmt"
    "time"
    "gorm.io/gorm"
)

type Patient struct {
    ID    uint   `gorm:"primaryKey" json:"id"`
    Code  string `gorm:"uniqueIndex;size:20" json:"code"`
    Name  string `json:"name"`
    Phone string `json:"phone"`
}

func (p *Patient) BeforeCreate(tx *gorm.DB) error {
    now := time.Now()
    year := now.Format("06")
    month := now.Format("01")
    
    var count int64
    prefix := fmt.Sprintf("EM-%s%s", year, month)
    tx.Model(&Patient{}).Where("code LIKE ?", prefix+"%").Count(&count)
    
    p.Code = fmt.Sprintf("EM-%s%s%04d", year, month, count+1)
    return nil
}
