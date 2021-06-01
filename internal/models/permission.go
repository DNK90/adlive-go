package models

import "gorm.io/gorm"

type Permission struct {
	Id          int    `json:"id"       mapstructure:"id"          gorm:"primary_key:true;column:id;auto_increment;not null"`
	Role        int    `json:"role"     mapstructure:"role"        gorm:"column:role;not null;"`
	Resource    int    `json:"resource" mapstructure:"resource"    gorm:"column:resource;not null;"`
	Allowed     int    `json:"allowed"  mapstructure:"allowed"     gorm:"column:allowed;not null"`
	Status      int    `json:"status"   mapstructure:"status"      gorm:"column:status;not null"`
	CreatedDate int64  `json:"created_date" mapstructure:"created_date" gorm:"column:created_date;not null"`
	UpdatedDate int64  `json:"updated_date" mapstructure:"updated_date" gorm:"column:updated_date;not null"`
}

func (r Permission) TableName() string {
	return "permissions"
}

func (r Permission) BeforeCreate(tx *gorm.DB) (err error) {
	return nil
}

func (r Permission) ToMap() map[string]interface{} {
	return map[string]interface{}{
		//"id":          r.Id,
		"role":     r.Role,
		"resource": r.Resource,
		"allowed":     r.Allowed,
		"status":      r.Status,
	}
}
