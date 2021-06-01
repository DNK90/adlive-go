package models

import "gorm.io/gorm"

type Role struct {
	Id          int    `json:"id"        mapstructure:"id"          gorm:"primary_key:true;column:id;auto_increment;not null"`
	Name        string `json:"name"      mapstructure:"name"        gorm:"column:name;not null;unique"`
	ParentId    int    `json:"parent_id" mapstructure:"parent_id"   gorm:"column:parent_id;not null"`
	Status      int    `json:"status"    mapstructure:"status"      gorm:"column:status;not null"`
	CreatedDate int64  `json:"created_date" mapstructure:"created_date" gorm:"column:created_date;not null"`
	UpdatedDate int64  `json:"updated_date" mapstructure:"updated_date" gorm:"column:updated_date;not null"`
}

func (r Role) TableName() string {
	return "roles"
}

func (r Role) BeforeCreate(tx *gorm.DB) (err error) {
	return nil
}

func (r Role) ToMap() map[string]interface{} {
	return map[string]interface{}{
		//"id":        r.Id,
		"name":      r.Name,
		"status":    r.Status,
		"parent_id": r.ParentId,
	}
}

func (r Role) GetId() int {
	return r.Id
}

func (r Role) GetParentId() int {
	return r.ParentId
}

func (r Role) GetName() string {
	return r.Name
}
