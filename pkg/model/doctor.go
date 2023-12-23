package model

import "gorm.io/gorm"

type Doctor struct {
    gorm.Model
    DoctorID           uint   `gorm:"primaryKey"`
    UserID             uint   `gorm:"unique;index"`
    Specialization     string
    Qualifications     string `gorm:"type:text"`
    YearsOfExperience  int
}
