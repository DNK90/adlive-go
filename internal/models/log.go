package models

import "gorm.io/gorm"

type Log struct {
	Id 			int 	`json:"id"           mapstructure:"id"           gorm:"primary_key:true;column:id;auto_increment;not null"`
	Activity	string 	`json:"activity"     mapstructure:"activity"     gorm:"column:activity;not null;"`
	Type	    string 	`json:"type"         mapstructure:"type"         gorm:"column:type;not null;"`
	ScreenId    int     `json:"screen_id"    mapstructure:"screen_id"    gorm:"column:screen_id;not null;"`
	CreatedDate int64 	`json:"created_date" mapstructure:"created_date" gorm:"column:created_date;not null"`
	UserId      int 	`json:"user_id"      mapstructure:"user_id"      gorm:"column:user_id;not null"`
}

func (r Log) TableName() string {
	return "log"
}

func (r Log) BeforeCreate(tx *gorm.DB) (err error) {
	return nil
}

func (r Log) ToMap() map[string]interface{} {
	return map[string]interface{}{
		//"id":        r.Id,
		"name":      	r.Activity,
		"type":    	    r.Type,
		"screen_id":    r.ScreenId,
		"user_id":      r.UserId,
		"created_date": r.CreatedDate,
	}
}
