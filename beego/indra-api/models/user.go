package models

import (
	"github.com/astaxie/beego/orm"
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"
	"github.com/vjeantet/jodaTime"
)


func init() {
	orm.RegisterModel(new(Users))
}

type Users struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Status int `json:"status"`
	CreatedDate string `json:"created_date"`
	UpdatedDate string `json:"updated_date"`

}

type UserFilter struct {
	Username string `sql:"username"`
	Status interface{} `sql:"status"`
}


type ResponseUser struct {
	Count int
	Page int
	Perpage int
	Data []Users
}



func AddUser(users Users) int64 {
	oRM := orm.NewOrm()
	hasher := md5.New()
	hasher.Write([]byte(users.Password))
	oRM.Begin()

	users.Password = hex.EncodeToString(hasher.Sum(nil))
	users.Status = 1 //active
	users.CreatedDate =  jodaTime.Format("YYYY-MM-dd HH:mm:ss", time.Now()) //time now
	users.UpdatedDate =  jodaTime.Format("YYYY-MM-dd HH:mm:ss", time.Now()) //time now

	id_user,err := oRM.Insert(&users)
	if(err != nil){
		println(users.CreatedDate)
		println(err.Error())
		oRM.Rollback()
		return 0
	}
	oRM.Commit()
	return id_user
}

func GetUser(username string) ResponseUser {
	oRM := orm.NewOrm()
	var mapsUser []orm.Params
	var users Users
	var dataUser []Users
	var resUser ResponseUser
	var condAll *orm.Condition
	queryString := oRM.QueryTable("users")
	cond := orm.NewCondition()

	if(username != ""){
		condAll = cond.And("username",username)
	}

	num,_ :=queryString.SetCond(condAll).Limit(1).Offset(0).Values(&mapsUser)
	resUser.Count = int(num)

	if(num > 0) {
		for _, mUser := range mapsUser {
			users.Status = int(mUser["Status"].(int64))
			users.Username = mUser["Username"].(string)
			if(mUser["UpdatedDate"] != nil) {
				users.UpdatedDate = mUser["UpdatedDate"].(string)
			}
			if(mUser["CreatedDate"] != nil) {
				users.CreatedDate = mUser["CreatedDate"].(string)
			}
			users.Id= int(mUser["Id"].(int64))
			dataUser = append(dataUser,users)
		}
	}
	resUser.Data = dataUser
	return resUser

}

func GetAllUsers(page interface{},perpage interface{}, filter UserFilter) ResponseUser {
	oRM := orm.NewOrm()
	var users Users
	var dataUser []Users
	var mapsUser []orm.Params
	var resUser ResponseUser
	var condUsername *orm.Condition
	var condStatus *orm.Condition
	var offset int
	perpage_int,_ := strconv.Atoi(perpage.(string))
	page_int,_ := strconv.Atoi(page.(string))
	queryString := oRM.QueryTable("users")
	cond := orm.NewCondition()

	if(filter.Username != ""){
		condUsername = cond.And("username__contains",filter.Username)
	}

	if(filter.Status != ""){
		condStatus = cond.And("status",filter.Status)
	}

	condAll := cond.AndCond(condUsername).AndCond(condStatus)

	if(page.(string) !="" && perpage.(string) !="") {

		offset = (page_int-1) * perpage_int
		resUser.Page = page_int
		resUser.Perpage = perpage_int
	}else{
		offset = 0
		perpage_int = 25
		resUser.Page = offset
		resUser.Perpage = perpage_int
	}

	queryString.SetCond(condAll).Limit(perpage_int).Offset(offset).Values(&mapsUser)
	num, _ := queryString.SetCond(condAll).Count()
	resUser.Count = int(num)

	if(num > 0) {
		for _, mUser := range mapsUser {
			users.Status = int(mUser["Status"].(int64))
			users.Username = mUser["Username"].(string)
			if(mUser["UpdatedDate"] != nil) {
				users.UpdatedDate = mUser["UpdatedDate"].(string)
			}
			if(mUser["CreatedDate"] != nil) {
				users.CreatedDate = mUser["CreatedDate"].(string)
			}
			users.Id= int(mUser["Id"].(int64))
			dataUser = append(dataUser,users)
		}
	}
	resUser.Data = dataUser
	return resUser
}

func UpdateUser(uid string, uu *Users) (a *Users, err error) {
	return
}

func Login(username, password string) bool {

	hasher := md5.New()
	hasher.Write([]byte(password))
	password_md5 := hex.EncodeToString(hasher.Sum(nil))

	oRM := orm.NewOrm()
	queryString := oRM.QueryTable("users")
	cond := orm.NewCondition()

	num,err := queryString.SetCond(cond.And("username",username).And("password",password_md5)).Count()

	if(num > 0 && err == nil){
		return true
	}else{
		return false
	}

}

func DeleteUser(uid string) {
	return
}
