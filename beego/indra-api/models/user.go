package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"crypto/md5"
	"encoding/hex"
	"strings"
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

//struct for Response User
type ResponseUser struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Offset int `json:"offset"`
	Limit int `json:"limit"`
	Data  []orm.Params
}

func AddUser(u User) string {
	return "";
}

func GetUser(uid string) (u *User, err error) {
	if u, ok := UserList[uid]; ok {
		return u, nil
	}
	return nil, errors.New("User not exists")
}

func GetAllUsers(offset interface{},limit interface{}, filter UserFilter) []orm.Params {
	oRM := orm.NewOrm()
	var mapsUser []orm.Params
	var whereArr []string
	whereCondition := ""
	limitOffset := " LIMIT 0,25 "

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
	}

	oRM.Raw("Select username, status, created_date, updated_date FROM users "+whereCondition +limitOffset).Values(&mapsUser)

	return mapsUser
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
