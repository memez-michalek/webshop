package errorCodes

const (
	USERDOESNOTEXIST     = "User does not exist"
	COULDNOTFINDAPIKEY   = "Could not find api key"
	SHOPDOESNOTEXIST     = "Shop does not exist"
	COULDNOTBIND         = "COULD NOT BIND TO MODEL"
	USERALREADYLOGGEDIN  = "You are already logged in"
	TOKENERROR           = "token related error"
	TOKENEXPIRED         = "token is expired"
	CLAIMERROR           = "Could not read claim values"
	COULDNOTINSERTINTODB = "Could not insert product/ shop into DATABASE"
	SHOPALREADYEXISTS    = "This particular id or name is used in different shop"
	SHOPWASTNINSERTED    = "Shop wasnt added to the database"
	PRODUCTDOESNOTEXIST  = "product was not find in database"
	COULDNOTFINDPRODUCTS = "COULD NOT FIND PRODUCTS BY ID"
	COULDNOTFINDSHOP     = "COULD NOT FIND PARTICULAR SHOP"
)
