package types

import "gorm.io/gorm"

type Task map[interface{}]interface{}

type Credentials struct {
	gorm.Model
	Host     string `json:"host"`
	From     string `json:"from"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Agent struct {
	gorm.Model
	Pid      int           `json:"pid"`
	Arch     string        `json:"arch"`
	Tasks    []Task        `json:"tasks"`
	Creds    []Credentials `json:"creds"`
	Process  string        `json:"process"`
	Username string        `json:"username"`
	Computer string        `json:"computer"`
}
