package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
	"time"
)

type Product struct {
	Id              uint32      `json:"id"`
	Store           *Store      `orm:"rel(fk);column(store_id)" json:"store"`
	Title           string      `orm:"size(100)" valid:"Required" json:"title"`
	Description     string      `orm:"size(100)" json:"description"`
	Size            int         `orm:"size(11)"  valid:"Required"   json:"size"`
	CreatedAt		time.Time 	`orm:"auto_now_add;type(datetime)" json:"created_at"`
	CreatedBy		string	  	`json:"created_by"`
	UpdatedAt		time.Time 	`orm:"auto_now;type(datetime)" json:"updated_at"`
	UpdatedBy		string 		`orm:"null" json:"updated_by"`
}

func init(){
	orm.RegisterModel(new(Product))
}

func GetProducts()(products []Product,err error){
	
	o := orm.NewOrm()
	
	_,err =o.QueryTable("Product").
		All(&products)
	
	if err != nil{
		return products,errors.New("Internal Error!")
	}
	
	return products,nil
}

