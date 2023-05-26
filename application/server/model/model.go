package model

// Account记录每个Peer节点的数据
type Account struct {
	AccountID string //账号ID
	UserName  string //节点名字
	Balance   int    //衡量一个节点的交易数量
}

type Picture struct {
	AccountID []byte   //发布者|权限
	Time      []byte   //时间
	Image     []byte   //图像
	Names     [][]byte //拍摄照片内包含的人名列表
}

const (
	AccountKey = "account-key"
	PictureKey = "picture-key"
)
