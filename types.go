package main

import (
	"encoding/json"
)

type Data struct{
	Id int `json:"id"`
	User string `json:"user"`
	Status string `json:"status"`
	Createdate string `json:"createData"`
	Updatedate string `json:"updateDate"`
}

func (d *Data) ToJson() ([]byte, error){
	return json.Marshal(d)
}