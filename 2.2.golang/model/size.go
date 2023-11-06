package model

type Size struct {
    Id        int    `gorm:"type:int;primary_key"`
    Name      string `gorm:"type:varchar(255)"`
    Value     string `gorm:"type:varchar(255)"`
    CreatedAt string `gorm:"type:varchar(19);default:to_char(now(), 'YYYY-MM-DD HH24:MI:SS')"`
}