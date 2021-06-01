package models

import "gorm.io/gorm"

type Video struct {
	Id          int    	`json:"id"           mapstructure:"id"           gorm:"primary_key:true;column:id;auto_increment;not null"`
	Title 		string 	`json:"title"        mapstructure:"title"        gorm:"column:title;not null"`
	Name        string 	`json:"name"         mapstructure:"name"         gorm:"column:name;not null"`
	Format 		string 	`json:"format"       mapstructure:"format"       gorm:"column:format;not null"`
	URL 		string 	`json:"url"          mapstructure:"url"          gorm:"column:url;not null"`
	PictureURL  string 	`json:"picture_url"  mapstructure:"picture_url"  gorm:"column:picture_url;not null"`
	OwnerId 	int 	`json:"owner_id"     mapstructure:"owner_id"     gorm:"column:owner_id;not null"`
	Size        int 	`json:"size"         mapstructure:"size"         gorm:"column:size;not null"`
	Status 		int 	`json:"status"       mapstructure:"status"       gorm:"column:status;not null"`
	CreatedDate int64 	`json:"created_date" mapstructure:"created_date" gorm:"column:created_date;not null"`
	User        User    `gorm:"foreignKey:OwnerId"`
	Campaigns   []CampaignVideo	`gorm:"foreignKey:VideoId;references:Id"`
}

func (r Video) TableName() string {
	return "video"
}

func (r Video) BeforeCreate(tx *gorm.DB) (err error) {
	return nil
}

func (r Video) ToMap() map[string]interface{} {
	return map[string]interface{}{
		//"id":        r.Id,
		"title":        r.Title,
		"name":      	r.Name,
		"format":       r.Format,
		"url":          r.URL,
		"picture_url":  r.PictureURL,
		"size":         r.Size,
		"owner_id":    	r.OwnerId,
		"status": 		r.Status,
		"created_date": r.CreatedDate,
	}
}
