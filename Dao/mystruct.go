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
	Expire    int64  `json:"expire"`
	FileName  string `json:"fileName"`
	FileSize  int64  `json:"fileSize"`
	Owner     string `json:"owner"`
	Use_count int64  `json:"use"`
	UseLimit  int64  `json:"useLimit"`
}
type ShareFile struct {
	Expire    int64  `json:"expire"`
	FileName  string `json:"fileName"`
	Target    string `json:"target"`
	Use_count int64  `json:"use"`
	UseLimit  int64  `json:"useLimit"`
	IsGroup   int64  `json:"isGroup"`
	FileSize  int64  `json:"fileSize"`
}
type ReceiveFile struct {
	/*
		文件类型标准
		文件名称、文件ID、文件大小、使用次数、文件拥有者ID、文件被共享者ID、共享策略ID
	*/
	FileName     string `json:"fileName"`
	FileSize     string `json:"fileSize"`
	UseCount     int    `json:"useCount"`
	OwnerId      int    `json:"ownerId"`
	TargetUserId int    `json:"targetUserId"`
	StragetyId   int    `json:"stragetyId"`
}
type Stragety struct {
	Name     string `json:"name"`
	Desc     string `json:"description"`
	Expire   int64  `json:"expire"`
	UseLimit int64  `json:"useLimit"`
}
