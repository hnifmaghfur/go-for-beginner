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


//struct for Response Login
type ResponseLogin struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Tokensting string `json:"tokensting"`
}

type ResponseUser struct {
	Count int
	Page int
	Perpage int
	Data []orm.Params
}

//struct for Response Get User
type ResponseGetUser struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Page int `json:"page"`
	Perpage int `json:"perpage"`
	Count int `json:"count"`
	Data  []orm.Params `json:"data"`
}

//struct for Response Get All User
type ResponseGetAllUser struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Page int `json:"page"`
	Perpage int `json:"perpage"`
	Count int `json:"count"`
	Data  []orm.Params `json:"data"`
	Pages PagesList `json:"pages"`
}

type PagesList struct{
	First string `json:"first"`
	Prev []string `json:"prev"`
	Next []string `json:"next"`
	Last string `json:"last"`
}

type ResponseInsertUser struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Data  struct{Id_user int64 `json:"id_user"`} `json:"data"`
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

	whereCondition := ""
	limitOffset := " LIMIT 0,1 "
	var resUser ResponseUser

	if(username != ""){
		fUsername := " username = '"+username+"'"
		whereCondition = " WHERE "+fUsername
	}

	num,_ :=oRM.Raw("Select username, status, created_date, updated_date FROM users "+whereCondition +limitOffset).Values(&mapsUser)

	resUser.Count = int(num)
	resUser.Data = mapsUser
	return resUser

}

func GetAllUsers(page interface{},perpage interface{}, filter UserFilter) ResponseUser {
	oRM := orm.NewOrm()
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
	resUser.Data = mapsUser

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
	var maps []orm.Params
	num, err := oRM.Raw("SELECT username, password FROM users WHERE username = ? AND password = ? LIMIT 1", username,password_md5).Values(&maps)

	if(num > 0 && err == nil){
		return true
	}else{
		return false
	}

}

func DeleteUser(uid string) {
	return
}
