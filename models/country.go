package models

type Country struct {
	Id       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Region   string `json:"region"`
	Language string `json:"language"`
}

// changes the name of the table
func (b *Country) TableName() string {
	return "countries"
}
