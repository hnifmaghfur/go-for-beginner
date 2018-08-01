package models

import (
	"errors"
	"strconv"
	"time"
	"github.com/astaxie/beego/orm"
	"crypto/md5"
	"encoding/hex"
)

var (
	UserList map[string]*User
)

func init() {
	UserList = make(map[string]*User)
	u := User{"user_11111", "astaxie", "11111", Profile{"male", 20, "Singapore", "astaxie@gmail.com"}}
	UserList["user_11111"] = &u
}

type User struct {
	Id       string
	Username string
	Password string
	Profile  Profile
}

type Profile struct {
	Gender  string
	Age     int
	Address string
	Email   string
}

//struct for Response Login
type ResponseLogin struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Tokensting string `json:"tokensting"`
}

func AddUser(u User) string {
	u.Id = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	UserList[u.Id] = &u
	return u.Id
}

func GetUser(uid string) (u *User, err error) {
	if u, ok := UserList[uid]; ok {
		return u, nil
	}
	return nil, errors.New("User not exists")
}

func GetAllUsers() map[string]*User {
	return UserList
}

func UpdateUser(uid string, uu *User) (a *User, err error) {
	if u, ok := UserList[uid]; ok {
		if uu.Username != "" {
			u.Username = uu.Username
		}
		if uu.Password != "" {
			u.Password = uu.Password
		}
		if uu.Profile.Age != 0 {
			u.Profile.Age = uu.Profile.Age
		}
		if uu.Profile.Address != "" {
			u.Profile.Address = uu.Profile.Address
		}
		if uu.Profile.Gender != "" {
			u.Profile.Gender = uu.Profile.Gender
		}
		if uu.Profile.Email != "" {
			u.Profile.Email = uu.Profile.Email
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
