package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Store struct {
	Id              uint32   `json:"id"`
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	IsActive        int      `json:"is_active"`
	CreatedAt		time.Time 	`orm:"auto_now_add;type(datetime)" json:"created_at"`
	CreatedBy		string	  	`json:"created_by"`
	UpdatedAt		time.Time 	`orm:"auto_now;type(datetime)" json:"updated_at"`
	UpdatedBy		string 		`orm:"null" json:"updated_by"`
}

func init(){
	orm.RegisterModel(new(Store))
}
