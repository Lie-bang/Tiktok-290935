package consts

const (
	CommentServiceName = "douyincomment"
	CommentTableName   = "comment"
	CommentServiceAddr = ":9005"

	FavoriteServiceName = "douyinfavorite"
	FavoriteTableName   = "favorite"
	FavoriteServiceAddr = ":9004"

	VideoServiceName = "douyinvideo"
	VideoTableName   = "Video"
	Endpoint         = "127.0.0.1:9000"
	AccessKeyID      = "minioadmin"
	SecretAccessKey  = "minioadmin"
	UseSSL           = false
	BucketName       = "video"
	VideoServiceAddr = ":9003"

	TempVideoFilePath = "./tempVideoFile/"
	HttpTemplate      = "http://"
	LocalIp           = "192.168.0.109"
	//LocalIp           = "10.11.116.158"
	//LocalIP   = "192.168.30.253"
	MinioPort = ":9000"
	/**/
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
	UserServiceAddr     = ":8999"
	RelationServiceAddr = ":9001"
	MessageServiceAddr  = ":9002"
	ExportEndpoint      = ":4317"
	ETCDAddress         = "127.0.0.1:2379"
	RedisPassword       = "dlb123456"
	RedisAddr           = "localhost:6379"
	RedisChatRecord     = "ChatRecord"
	UserFollowList      = "{userfollowlist}:"
	UserFollowerList    = "{userfollowerlist}:"
)
