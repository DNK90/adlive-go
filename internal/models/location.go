package models

import "gorm.io/gorm"

type Location struct {
	Id 			int 	`json:"id"           mapstructure:"id"           gorm:"primary_key:true;column:id;auto_increment;not null"`
	Name		string 	`json:"name"         mapstructure:"name"         gorm:"column:name;not null;"`
	Address		string 	`json:"address"      mapstructure:"address"      gorm:"column:address;not null;"`
	OwnerId     int     `json:"ownerId"      mapstructure:"owner_id"     gorm:"column:owner_id;not null"`
	Status      int     `json:"status"       mapstructure:"status"       gorm:"column:status;not null"`
	CreatedDate int64 	`json:"created_date" mapstructure:"created_date" gorm:"column:created_date;not null"`
	UpdatedDate int64   `json:"updated_Date" mapstructure:"updated_date" gorm:"column:updated_date;not null"`
	User        User    `gorm:"foreignKey:OwnerId"`
	Campaigns   []CampaignLocation `gorm:"foreignKey:LocationId"`
	Areas 		[]Area
	Screens     []Screen
}

func (r Location) TableName() string {
	return "location"
}

func (r Location) BeforeCreate(tx *gorm.DB) (err error) {
	return nil
}

func (r Location) ToMap() map[string]interface{} {
	return map[string]interface{}{
		//"id":        r.Id,
		"name":      	r.Name,
		"address":      r.Address,
		"owner_id":    	r.OwnerId,
		"status": 		r.Status,
		"created_date": r.CreatedDate,
		"updated_date": r.UpdatedDate,
	}
}
