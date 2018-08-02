package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"
)

var (
	UserList map[string]*User
)

func init() {
	UserList = make(map[string]*User)
	u := User{"user_11111", "astaxie", "11111"}
	UserList["user_11111"] = &u
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
	Status int `sql:"status"`
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

func GetAllUsers(offset int,limit int, filter *UserFilter) []orm.Params {
	oRM := orm.NewOrm()
	var mapsUser []orm.Params
	whereArr := make(map[int]string)
	whereCondition := ""
	limitOffset := ""

	i:= 0
	if(filter.Username != ""){
		whereArr[i] = whereCondition+" username LIKE '%"+filter.Username+"%'"
		i++
	}

	if(filter.Status != "" || filter.Status != 0){
		whereArr[i] = whereCondition+" status ="+filter.Status
		i++
	}

	if(len(whereArr)>0){
		whereCondition = "WHERE "+strings.Join(whereArr," AND ")
	}

	if(limit !="" && offset !="") {
		limitOffset = "LIMIT " + strconv.Itoa(offset) + "," + strconv.Itoa(limit)
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
