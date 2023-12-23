package model

import "gorm.io/gorm"

type ClinicLocation struct {
    gorm.Model
    LocationID   uint   `gorm:"primaryKey"`
    Name         string
    Address      string `gorm:"type:text"`
    ContactInfo  string
}

