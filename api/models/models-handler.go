package models


type Product struct {
	name     string
	quantity int
	price    float64
	category string
	id       string
}

type Key struct {
	id       string
	username string
	value    string
}
