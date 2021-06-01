package models

import "gorm.io/gorm"

type Campaign struct {
	Id 			int 	`json:"id"           mapstructure:"id"           gorm:"primary_key:true;column:id;auto_increment;not null"`
	Name		string 	`json:"name"         mapstructure:"name"         gorm:"column:name;not null"`
	OwnerId     int     `json:"owner_id"     mapstructure:"owner_id"     gorm:"column:owner_id;not null"`
	Played      int     `json:"played"       mapstructure:"played"       gorm:"column:played;not null"`
	Status      int     `json:"status"       mapstructure:"status"       gorm:"column:status;not null"`
	StartDate   int64 	`json:"start_date"   mapstructure:"start_date"   gorm:"column:start_date;not null"`
	EndDate     int64   `json:"end_date"     mapstructure:"end_date"     gorm:"column:end_date;not null"`

	CreatedDate int64 	`json:"created_date" mapstructure:"created_date" gorm:"column:created_date;not null"`
	UpdatedDate int64   `json:"updated_date" mapstructure:"updated_date" gorm:"column:updated_date;not null"`
	Locations   []CampaignLocation
	Screens     []CampaignScreen
	Videos      []CampaignVideo
	User        User    `gorm:"foreignKey:OwnerId"`
}

func (r Campaign) TableName() string {
	return "campaign"
}

func (r Campaign) BeforeCreate(tx *gorm.DB) (err error) {
	return nil
}

func (r Campaign) ToMap() map[string]interface{} {
	return map[string]interface{}{
		//"id":        r.Id,
		"name":      	r.Name,
		"owner_id":    	r.OwnerId,
		"status": 		r.Status,
		"start_date":   r.StartDate,
		"end_date":     r.EndDate,
		"created_date": r.CreatedDate,
		"updated_date": r.UpdatedDate,
	}
}


