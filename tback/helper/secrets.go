package helper

var JWTSECRET = "mowili mi ze nikim nie bede udowodnie im ze byli w bledzie "
var SHOPAPISECRET = struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Key      string `json:"key"`
}{
	Username: "e",
	Email:    "e@e.com",
	Key:      "1pG3MWsEflP7HG8hpClS62Zlmgp",
}
var SHOPAPIWEBTOKEN = ""

var SHOPAPIKEY = struct {
	Token string `json:"key"`
}{
	Token: SHOPAPIWEBTOKEN,
}
