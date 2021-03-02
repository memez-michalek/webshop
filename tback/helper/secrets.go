package helper

var JWTSECRET = "mowili mi ze nikim nie bede udowodnie im ze byli w bledzie "
var SHOPAPISECRET = struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Key      string `json:"key"`
}{
	Username: "cypis",
	Email:    "cypis@cypis.com",
	Key:      "1ntoZdCfBUb6u1RTUfkXKioqvxc",
}
var SHOPAPIWEBTOKEN = ""

var SHOPAPIKEY = struct {
	Token string `json:"key"`
}{
	Token: SHOPAPIWEBTOKEN,
}
