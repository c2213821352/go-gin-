package model

//type User struct {
//	UserName string `gorm:"varchar(10);not null"`
//	Password string `gorm:"varchar(10);not null"`
//}

//type User struct {
//	status   uint   `gorm:"primary_key"`
//	Username string `form:"username"  gorm:"varchar(10);not null"`
//	Password string `form:"password"  gorm:"varchar(10);not null"`
//	//Id       uuid.UUID `form:"type:varchar(100);not null" gorm:"primary_key;AUTO_INCREMENT"`
//	//CreatedAt time.Time `gorm:"autoCreateTime"`
//	//CreatedAt time.Time `gorm:"autoCreateTime"`
//}

type User struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Static     bool   `json:"static"`
	Token      string `json:"token"`
	Money      int    `json:"money"`
	MaxScore   int    `json:"maxscore"`
	Coin       int    `json:"coin"`
}
