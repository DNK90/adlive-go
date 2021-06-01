package models

import "gorm.io/gorm"

type CampaignScreen struct {
	Id 			int 	`json:"id"            mapstructure:"id"            gorm:"primary_key:true;column:id;auto_increment;not null"`
	OwnerId 	int 	`json:"owner_id"      mapstructure:"owner_id"      gorm:"column:owner_id;not null;uniqueIndex:idx_campaign_screen"`
	CampaignId	int 	`json:"campaign_id"   mapstructure:"campaign_id"   gorm:"column:campaign_id;not null;uniqueIndex:idx_campaign_screen"`
	ScreenId 	int		`json:"screen_id"     mapstructure:"screen_id"     gorm:"column:screen_id;not null;uniqueIndex:idx_campaign_screen"`
	Status 		int 	`json:"status"        mapstructure:"status"        gorm:"column:status;not null"`
	CreatedDate int64 	`json:"created_date"  mapstructure:"updated_date"  gorm:"column:created_date;not null"`
	Campaign	Campaign
	Screen 		Screen
	User        User    `gorm:"foreignKey:OwnerId"`
}

func (r CampaignScreen) TableName() string {
	return "campaign_screen"
}

func (r CampaignScreen) BeforeCreate(tx *gorm.DB) (err error) {
	return nil
}

func (r CampaignScreen) ToMap() map[string]interface{} {
	return map[string]interface{}{
		//"id":        r.Id,
		"campaign_id":  r.CampaignId,
		"screen_id":  	r.ScreenId,
		"status": 		r.Status,
		"created_date": r.CreatedDate,
	}
}

