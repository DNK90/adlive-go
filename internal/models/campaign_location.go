package models

import "gorm.io/gorm"

type CampaignLocation struct {
	Id 			int 	`json:"id"                  mapstructure:"id"              gorm:"primary_key:true;column:id;auto_increment;not null"`
	OwnerId 	int 	`json:"owner_id"            mapstructure:"owner_id"        gorm:"column:owner_id;not null;uniqueIndex:idx_campaign_location"`
	CampaignId	int 	`json:"campaign_id"         mapstructure:"campaign_id"     gorm:"column:campaign_id;not null;uniqueIndex:idx_campaign_location"`
	LocationId  int     `json:"location_id"         mapstructure:"location_id"     gorm:"column:location_id;not null;uniqueIndex:idx_campaign_location"`
	AreaId  	int 	`json:"area_id"             mapstructure:"area_id"         gorm:"column:area_id;not null;uniqueIndex:idx_campaign_location"`
	Status      int     `json:"status"              mapstructure:"status"          gorm:"column:status;not null"`
	CreatedDate int64	`json:"created_date"        mapstructure:"created_date"    gorm:"column:created_date;not null"`
	Campaign 	Campaign
	Location 	Location
	Area        Area
	User        User    `gorm:"foreignKey:OwnerId"`
}

func (r CampaignLocation) TableName() string {
	return "campaign_location"
}

func (r CampaignLocation) BeforeCreate(tx *gorm.DB) (err error) {
	return nil
}

func (r CampaignLocation) ToMap() map[string]interface{} {
	return map[string]interface{}{
		//"id":        r.Id,
		"owner_id":		r.OwnerId,
		"campaign_id":  r.CampaignId,
		"location_id":  r.LocationId,
		"area_id":		r.AreaId,
		"status": 		r.Status,
		"created_date": r.CreatedDate,
	}
}

func (r CampaignLocation) Fields() []string {
	fields := make([]string, 0)
	fields = append(fields, "id")
	for field, _ := range r.ToMap() {
		fields = append(fields, field)
	}
	return fields
}


