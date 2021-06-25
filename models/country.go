package models

import "time"

type Detail struct {
	ID       uint   `gorm:"primary_key;auto_increment" json:"id"`
	Language string `json:"language" binding:"required" gorm:"type:varchar(20)"`
	Sport    string `json:"sport" binding:"required" gorm:"type:varchar(20)"`
	Year     uint   `json:"from_year" binding:"gte=1000,lte=2050"`
}

type Country struct {
	ID          uint      `json:"id" gorm:"primary_key;auto_increment"`
	Name        string    `json:"name" binding:"required" gorm:"type:varchar(30)"`
	Region      string    `json:"region" binding:"required" gorm:"type:varchar(30)"`
	Information Detail    `json:"details" binding:"required" gorm:"foreignkey:DetailsID"`
	DetailsID   uint      `json:"-"`
	CreatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// changes the name of the table
func (b *Country) TableName() string {
	return "countries"
}

func (b *Detail) TableName() string {
	return "details"
}
