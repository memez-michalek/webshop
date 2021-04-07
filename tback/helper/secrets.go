package helper

var JWTSECRET = "mowili mi ze nikim nie bede udowodnie im ze byli w bledzie "
var SHOPAPISECRET = struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Key      string `json:"key"`
}{
	Username: "taco",
	Email:    "taco@taco.com",
	Key:      "1qUl8ylhDoUH5vmwkwEWMO2pwKX",
}
var SHOPAPIWEBTOKEN = ""

var SHOPAPIKEY = struct {
	Token string `json:"key"`
}{
	Token: SHOPAPIWEBTOKEN,
}
