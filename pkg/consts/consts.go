package consts

const (
	UserTableName       = "user"
	RelationTableName   = "follows"
	MessageTableName    = "message"
	SecretKey           = "secret key"
	IdentityKey         = "id"
	ApiServiceName      = "douyinapi"
	UserServiceName     = "douyinuser"
	RelationServiceName = "douyinrelation"
	MessageServiceName  = "douyinmessage"
	MySQLDefaultDSN     = "gorm:gorm@tcp(localhost:3306)/gorm?charset=utf8&parseTime=True&loc=Local"
	TCP                 = "tcp"
	UserServiceAddr     = ":9000"
	RelationServiceAddr = ":9001"
	MessageServiceAddr  = ":9002"
	ExportEndpoint      = ":4317"
	ETCDAddress         = "127.0.0.1:2379"
	DefaultLimit        = 10
)
