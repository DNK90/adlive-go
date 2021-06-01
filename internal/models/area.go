package models

import "gorm.io/gorm"

type Area struct {
	Id 			int 	`json:"id"           mapstructure:"id"           gorm:"primary_key:true;column:id;auto_increment;not null"`
	Name		string 	`json:"name"         mapstructure:"name"         gorm:"column:name;size:50;not null;uniqueIndex:nameLocation"`
	OwnerId     int     `json:"owner_id"     mapstructure:"owner_id"     gorm:"column:owner_id;not null"`
	LocationId  int     `json:"location_id"  mapstructure:"location_id"  gorm:"column:location_id;not null;uniqueIndex:nameLocation"`
	Status      int     `json:"status"       mapstructure:"status"       gorm:"column:status;not null"`
	CreatedDate int64 	`json:"created_date" mapstructure:"created_date" gorm:"column:created_date;not null"`
	UpdatedDate int64   `json:"updated_Date" mapstructure:"updated_date" gorm:"column:updated_date;not null"`
	User        User    `gorm:"foreignKey:OwnerId"`
	Location    Location
}

func (r Area) TableName() string {
	return "area"
}

func (r Area) BeforeCreate(tx *gorm.DB) (err error) {
	return nil
}

func (r Area) ToMap() map[string]interface{} {
	return map[string]interface{}{
		//"id":        r.Id,
		"name":      	r.Name,
		"owner_id":    	r.OwnerId,
		"location_id":  r.LocationId,
		"status": 		r.Status,
		"created_date": r.CreatedDate,
		"updated_date": r.UpdatedDate,
	}
}
