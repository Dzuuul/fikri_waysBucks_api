package models

type User struct {
	ID       int    `json:"id"`
	Role     string `json:"-"`
	Fullname string `json:"fullname" gorm:"type: varchar(255)"`
	Email    string `json:"email" gorm:"type: varchar(255)"`
	Password string `json:"-" gorm:"type: varchar(255)"`
	Image    string `json:"image" gorm:"type: varchar(255)"`
}

type UsersProfileResponse struct {
	ID       int    `json:"id"`
	Role     string `json:"role"`
	Fullname string `json:"fullname"`
}

func (UsersProfileResponse) TableName() string {
	return "users"
}
