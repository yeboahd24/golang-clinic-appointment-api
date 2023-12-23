package model

import "gorm.io/gorm"

type Service struct {
    gorm.Model
    ServiceID   uint   `gorm:"primaryKey"`
    Name        string
    Description string `gorm:"type:text"`
    Duration    int
}

