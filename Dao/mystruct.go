package Dao

// User 用户结构体（可根据需要扩展）
type User struct {
	/*
		构建用户对象存放在JWT中
		UserId ： 用户ID
		Username ： 用户名
		Role ： 角色ID
		PublishKey: 用户公钥
	*/
	UserId    int64  `json:"userId"`
	Username  string `json:"userName"`
	Role      int    `json:"role"`
	PublicKey string `json:"publicKey"`
}
type FileListElement struct {
	Expire   int64  `json:"expire"`
	FileName string `json:"fileName"`
	FileSize int64  `json:"fileSize"`
	IsAllow  int64  `json:"isAllow"`
	Owner    string `json:"owner"`
	UseCount int64  `json:"use"`
	UseLimit int64  `json:"useLimit"`
}
type ShareFile struct {
	Id       int64  `json:"id"`
	Expire   int64  `json:"expire"`
	FileName string `json:"fileName"`
	IsAllow  int64  `json:"isAllow"`
	Target   string `json:"target"`
	UseCount int64  `json:"use"`
	UseLimit int64  `json:"useLimit"`
	FileSize int64  `json:"fileSize"`
}
type BeShareFile struct {
	Id       int64  `json:"id"`
	Expire   int64  `json:"expire"`
	FileName string `json:"fileName"`
	IsAllow  int64  `json:"isAllow"`
	From     string `json:"from"`
	UseCount int64  `json:"use"`
	UseLimit int64  `json:"useLimit"`
	FileSize int64  `json:"fileSize"`
}
