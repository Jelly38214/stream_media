package defs

// UserCredential struct
// 打tag的方式， 序列化和反序列化时使用
// {user_name: xxx, pwd: xxx}
type UserCredential struct {
	Username string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

// VideoInfo struct
type VideoInfo struct {
	ID           string
	AuthodID     int
	Name         string
	DisplayCtime string
}
