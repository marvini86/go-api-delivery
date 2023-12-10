package entity

import "time"

type Delivery struct {
	Id          int64     `json:"id"`
	Customer    Customer  `json:"customer"`
	Items       []Item    `json:"items"`
	CreatedDate time.Time `json:"createdDate"`
}

type Item struct {
	Id          int64  `json:"id"`
	Description string `json:"description"`
	Quantity    int64  `json:"quantity"`
}

type Customer struct {
	Id      int64   `json:"id"`
	Name    string  `json:"name"`
	Address Address `json:"address"`
}

type Address struct {
	Id       int64  `json:"id"`
	Address  string `json:"address"`
	District string `json:"district"`
}
