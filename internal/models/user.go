package models

import "gorm.io/gorm"

type User struct {
	Id          int    	`json:"id"           mapstructure:"id"           gorm:"primary_key:true;column:id;auto_increment;not null"`
	Name        string 	`json:"name"         mapstructure:"name"         gorm:"column:name;not null"`
	Email       string 	`json:"email"        mapstructure:"email"        gorm:"column:email;not null;unique"`
	Phone       string 	`json:"phone"        mapstructure:"phone"        gorm:"column:phone;not null;unique"`
	Company     string 	`json:"company"      mapstructure:"company"      gorm:"column:company;not null"`
	Role        int    	`json:"role"         mapstructure:"role"         gorm:"column:role;not null"`
	Username    string 	`json:"username"     mapstructure:"username"     gorm:"column:user_name;not null;unique"`
	Password    string 	`json:"password"     mapstructure:"password"     gorm:"column:password;not null"`
	Status      int 	`json:"status"       mapstructure:"status"       gorm:"column:status;not null"`
	CreatedDate int64  	`json:"created_date" mapstructure:"created_date" gorm:"column:created_date;not null"`
	UpdatedDate int64  	`json:"updated_date" mapstructure:"updated_date" gorm:"column:updated_date;not null"`
}

func (u *User) ToMap() map[string]interface{} {
	return map[string]interface{}{
		//"id":           u.Id,
		"name":         u.Name,
		"email":        u.Email,
		"phone":        u.Phone,
		"company":      u.Company,
		"role":         u.Role,
		"created_date": u.CreatedDate,
		"username":     u.Username,
		"password":     u.Password,
		"status":       u.Status,
	}
}

func (u User) BeforeCreate(tx *gorm.DB) (err error) {
	return nil
}

func (u User) TableName() string {
	return "user"
}
