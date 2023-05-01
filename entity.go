package main

import "encoding/json"

// registro User con los tags para convertir a json
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
	Phone string `json:"phone"`
}

// interface vacia
type MetadataUser interface{}

func (this *User) toJson() ([]byte, error) {
	return json.Marshal(this)
}
