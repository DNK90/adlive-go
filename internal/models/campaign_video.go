package models

import "gorm.io/gorm"

type CampaignVideo struct {
	Id 			int 	`json:"id"            mapstructure:"id"            gorm:"primary_key:true;column:id;auto_increment;not null"`
	OwnerId 	int 	`json:"owner_id"      mapstructure:"owner_id"      gorm:"column:owner_id;not null;uniqueIndex:idx_campaign_video"`
	CampaignId	int 	`json:"campaign_id"   mapstructure:"campaign_id"   gorm:"column:campaign_id;not null;uniqueIndex:idx_campaign_video"`
	VideoId 	int		`json:"video_id"      mapstructure:"video_id"      gorm:"column:video_id;not null;uniqueIndex:idx_campaign_video"`
	Status 		int 	`json:"status"        mapstructure:"status"        gorm:"column:status;not null"`
	CreatedDate int64 	`json:"created_date"  mapstructure:"updated_date"  gorm:"column:created_date;not null"`
	Campaign	Campaign
	Video 		Video
	User        User    `gorm:"foreignKey:OwnerId"`
}

func (r CampaignVideo) TableName() string {
	return "campaign_video"
}

func (r CampaignVideo) BeforeCreate(tx *gorm.DB) (err error) {
	return nil
}

func (r CampaignVideo) ToMap() map[string]interface{} {
	return map[string]interface{}{
		//"id":        r.Id,
		"campaign_id":  r.CampaignId,
		"owner_id":     r.OwnerId,
		"video_id":  	r.VideoId,
		"status": 		r.Status,
		"created_date": r.CreatedDate,
	}
}
