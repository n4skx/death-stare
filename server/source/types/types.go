package types

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Type       string `json:"type"`
	Value      string `json:"value"`
	Finished   bool   `json:"finished"`
	AgentRefer uint
}

type Credentials struct {
	gorm.Model
	Username   string `json:"username"`
	Password   string `json:"password"`
	AgentRefer uint
}

type Agent struct {
	gorm.Model
	Pid      int           `json:"pid"`
	Arch     string        `json:"arch"`
	Sleep    int           `json:"sleep"`
	Tasks    []Task        `json:"tasks" gorm:"foreignKey:AgentRefer"`
	Creds    []Credentials `json:"creds" gorm:"foreignKey:AgentRefer;"`
	Process  string        `json:"process"`
	Username string        `json:"username"`
	Computer string        `json:"computer"`
}

type Operator struct {
	gorm.Model
	IsAdmin  bool   `json:"is_admin"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Listener struct {
	gorm.Model
	Name       string `json:"name"`
	Type       string `json:"type"`
	Bind       string `json:"bind"`
	Active     bool   `json:"active"`
	RealIP     string `json:"real_ip"`
	FullChain  string `json:"full_chain"`
	PrivateKey string `json:"private_key"`
}
