package models

import (
	"bitbucket.org/klokinnovations/webapp/models"
)

type Building struct {
	models.BaseModel `valid:"-"`
	Code             string `gorm:"type:varchar(16);unique;not null"`
	Name             string `gorm:"not null"`
	ShortName        string `gorm:"type:varchar(32);not null"`
}

type Floor struct {
	models.BaseModel `valid:"-"`

	Building Building `gorm:"foreignKey:BuildingCode;references:Code"`

	Code          string `gorm:"type:varchar(16);unique;not null"`
	BuildingCode  string `gorm:"type:varchar(16);not null;"`
	Name          string `gorm:"not null"`
	ShortName     string `gorm:"type:varchar(32);not null"`
	PhysicalIndex int    `gorm:"not null"`
}

type Section struct {
	models.BaseModel `valid:"-"`
	Floor            Floor `gorm:"foreignKey:FloorCode;references:Code"`

	Code      string `gorm:"type:varchar(16);unique;not null"`
	Name      string `gorm:"not null"`
	ShortName string `gorm:"type:varchar(32);not null"`
	FloorCode string `gorm:"type:varchar(16);not null"`
}
