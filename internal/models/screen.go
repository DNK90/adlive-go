package models

import "gorm.io/gorm"

type Screen struct {
	Id 			int 	`json:"id"           mapstructure:"id"           gorm:"primary_key:true;column:id;auto_increment;not null"`
	Name        string  `json:"name"         mapstructure:"name"         gorm:"column:name;not null;"`
	LocationId	int 	`json:"location_id"  mapstructure:"location_id"  gorm:"column:location_id;not null;"`
	AreaId		int 	`json:"area_id"      mapstructure:"area_id"      gorm:"column:area_id;not null;"`
	OwnerId     int     `json:"owner_id"     mapstructure:"owner_id"     gorm:"column:owner_id;not null"`
	DeviceToken string  `json:"device_token" mapstructure:"device_token" gorm:"column:device_token;not null"`
	DeviceType  string  `json:"device_type"  mapstructure:"device_type"  gorm:"column:device_type;not null"`
	MAC         string  `json:"mac"          mapstructure:"mac"          gorm:"column:mac;not null;unique"`
	OS          string  `json:"os"           mapstructure:"os"           gorm:"column:os;not null"`
	AppVersion  string  `json:"app_version"  mapstructure:"app_version"  gorm:"column:app_version;not null"`
	IPAddress   string  `json:"ip_address"   mapstructure:"ip_address"   gorm:"column:ip_address;not null"`
	Status      int     `json:"status"       mapstructure:"status"       gorm:"column:status;not null"`
	CreatedDate int64 	`json:"created_date" mapstructure:"created_date" gorm:"column:created_date;not null"`
	UpdatedDate int64   `json:"updated_Date" mapstructure:"updated_date" gorm:"column:updated_date;not null"`

	User        *User    `gorm:"foreignKey:OwnerId"`
	Location    *Location
	Area        *Area
	Campaigns   []CampaignScreen
}

func (r Screen) TableName() string {
	return "screen"
}

func (r Screen) BeforeCreate(tx *gorm.DB) (err error) {
	return nil
}

func (r Screen) ToMap() map[string]interface{} {
	return map[string]interface{}{
		//"id":        r.Id,
		"name":      	r.Name,
		"owner_id":    	r.OwnerId,
		"location_id":  r.LocationId,
		"area_id":      r.AreaId,
		"device_type":  r.DeviceType,
		"mac":          r.MAC,
		"os":           r.OS,
		"app_version":  r.AppVersion,
		"ip_address":   r.IPAddress,
		"status": 		r.Status,
		"created_date": r.CreatedDate,
		"updated_date": r.UpdatedDate,
	}
}


