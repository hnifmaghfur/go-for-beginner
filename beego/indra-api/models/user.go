package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"crypto/md5"
	"encoding/hex"
	"strings"
	"strconv"
)

var (
	UserList map[string]*User
)

func init() {


}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Status int `json:"status"`
	CreatedDate string `json:"created_date"`
	UpdateDate string `json:"update_date"`

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
	Offset int
	Limit int
	Data []orm.Params
}

//struct for Response Get User
type ResponseGetUser struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Offset int `json:"offset"`
	Limit int `json:"limit"`
	Count int `json:"count"`
	Data  []orm.Params `json:"data"`
}

//struct for Response Get All User
type ResponseGetAllUser struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Offset int `json:"offset"`
	Limit int `json:"limit"`
	Count int `json:"count"`
	Data  []orm.Params `json:"data"`
}

func AddUser(u User) string {
	return ""
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

func GetAllUsers(offset interface{},limit interface{}, filter UserFilter) ResponseUser {
	oRM := orm.NewOrm()
	var mapsUser []orm.Params
	var whereArr []string
	var limitOffset string
	var resUser ResponseUser
	whereCondition := ""

	if(filter.Username != ""){
		fUsername := " username LIKE '%"+filter.Username+"%'"
		whereArr = append(whereArr,fUsername)
	}

	if(filter.Status != ""){
		fStatus := " status ="+filter.Status.(string)
		whereArr = append(whereArr,fStatus)
	}

	if(len(whereArr)>0){
		whereCondition = "WHERE "+strings.Join(whereArr," AND ")
	}

	if(offset.(string) !="" && limit.(string) !="") {
		limitOffset = " LIMIT " + offset.(string)+ "," + limit.(string)
		offsetInt,_ := strconv.Atoi(offset.(string))
		limitInt,_ := strconv.Atoi(limit.(string))
		resUser.Offset = offsetInt
		resUser.Limit = limitInt
	}else{
		limitOffset = " LIMIT 0,25 "
		resUser.Offset = 0
		resUser.Limit = 25
	}

	num,_ :=oRM.Raw("Select username, status, created_date, updated_date FROM users "+whereCondition +limitOffset).Values(&mapsUser)
	resUser.Count = int(num)
	resUser.Data = mapsUser
	return resUser
}

func UpdateUser(uid string, uu *User) (a *User, err error) {
	if u, ok := UserList[uid]; ok {
		if uu.Username != "" {
			u.Username = uu.Username
		}
		return u, nil
	}
	return nil, errors.New("User Not Exist")
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
	delete(UserList, uid)
}
