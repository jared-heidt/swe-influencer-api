package models

type Creator struct {
	ID         uint32 `gorm:"primary_key; auto_increment" json:"id"`
	Name       string `gorm:"size:255; not null" json:"name"`
	Profession string `gorm:"size:255" json:"profession"`
	Focus      string `gorm:"size:255" json:"focus"`
	Youtube    string `gorm:"size:255" json:"youtube"`
}
