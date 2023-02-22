namespace go douyinapi

struct BaseResp {
    1: i64 status_code
    2: string status_message
    3: i64 service_time
}

struct User {
    1: i64 id
    2: string name
    3: i64 follow_count
    4: i64 follower_count
    5: bool is_follow
    6: string avatar
    7: string background_image
    8: string signature
    9: i64 total_favorited
    10: i64 work_count
    11: i64 favorite_count
}

struct FriendUser {
    1: i64 id
    2: string name
    3: i64 follow_count
    4: i64 follower_count
    5: bool is_follow
    6: string message
    7: i64 msgType
    8: string avatar
}

struct Message {
    1: i64 id
    2: i64 to_user_id
    3: i64 from_user_id
    4: string content
    5: i64 create_time
}


struct CreateUserRequest {
    1: string username (api.query="username", api.vd="len($) > 0")
    2: string password (api.query="password", api.vd="len($) > 0")
}

struct CreateUserResponse {
    1: BaseResp base_resp
}

struct CheckUserRequest {
    1: string username (api.query="username", api.vd="len($) > 0")
    2: string password (api.query="password", api.vd="len($) > 0")
}

struct CheckUserResponse {
    1: BaseResp base_resp
}

struct GetUserRequest {
    1: i64 user_id
    2: string token
}

struct GetUserResponse {
    1: BaseResp base_resp
}

struct ActionRequest{
    1: i64 to_user_id
    2: i32 action_type
}

struct ActionResponse{
    1: BaseResp base_resp
}

struct FollowListRequest{
    1: i64 user_id
}

struct FollowListResponse{
    1: BaseResp base_resp
}

struct FollowerListRequest{
     1: i64 user_id
 }

struct FollowerListResponse{
    1: BaseResp base_resp
}

struct FriendListRequest{
    1: i64 user_id
}

struct FriendListResponse{
    1: BaseResp base_resp
}

struct ChatRecordRequest{
    1: i64 to_user_id
}

struct ChatRecordResponse{
    1: BaseResp base_resp
}

struct SendMessageRequest{
    1: i64 to_user_id
    2: i64 action_type
    3: string content
}

struct SendMessageResponse{
    1: BaseResp base_resp
}


//videoPart:

struct Douyin_feed_request {
    1: optional i64 lastest_time // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
    2: optional string token // 可选参数，登录用户设置
}

struct Douyin_feed_response {
    1: required i32 status_code // 状态码，0-成功，其他值-失败
    2: optional string status_msg // 返回状态描述
    3: list<Video> video_list // 视频列表
    4: optional i64 next_time // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

struct Douyin_publish_action_request {
    1: required string token // 用户鉴权token
    2: required binary data // 视频数据
    3: required string title // 视频标题
//    4: required []byte daasd
}

struct Douyin_publish_action_response {
    1: required i32 status_code // 状态码，0-成功，其他值-失败
    2: optional string status_msg // 返回状态描述
}

struct Douyin_publish_list_request {
  1: required i64 user_id  // 用户id
  2: required string token  // 用户鉴权token
}

struct Douyin_publish_list_response {
  1: required i32 status_code  // 状态码，0-成功，其他值-失败
  2: optional string status_msg // 返回状态描述
  3: list<Video>  video_list  // 用户发布的视频列表
}

struct Douyin_favorite_action_request {
  1: required string token  // 用户鉴权token
  2: required i64 video_id // 视频id
  3: required i32 action_type  // otel-collector-config.yaml-点赞，2-取消点赞
}

struct Douyin_favorite_action_response {
  1: required i32 status_code  // 状态码，0-成功，其他值-失败
  2: optional string status_msg  // 返回状态描述
}

struct Douyin_favorite_list_request {
  1: required i64 user_id  // 用户id
  2: required string token  // 用户鉴权token
}

struct Douyin_favorite_list_response {
  1: required i32 status_code // 状态码，0-成功，其他值-失败
  2: optional string status_msg  // 返回状态描述
  3: list<Video> video_list  // 用户点赞视频列表
}

struct Douyin_comment_action_request {
  1: required string token  // 用户鉴权token
  2: required i64 video_id  // 视频id
  3: required i32 action_type // otel-collector-config.yaml-发布评论，2-删除评论
  4: optional string comment_text  // 用户填写的评论内容，在action_type=1的时候使用
  5: optional i64 comment_id  // 要删除的评论id，在action_type=2的时候使用
}

struct Douyin_comment_action_response {
  1: required i32 status_code  // 状态码，0-成功，其他值-失败
  2: optional string status_msg  // 返回状态描述
  3: optional Comment comment  // 评论成功返回评论内容，不需要重新拉取整个列表
}

struct Douyin_comment_list_request {
  1: required string token // 用户鉴权token
  2: required i64 video_id // 视频id
}

struct Douyin_comment_list_response {
  1: required i32 status_code  // 状态码，0-成功，其他值-失败
  2: optional string status_msg  // 返回状态描述
  3: list <Comment> comment_list  // 评论列表
}

struct Comment {
  1: required i64 id  // 视频评论id
  2: required User user  // 评论用户信息
  3: required string content  // 评论内容
  4: required string create_date  // 评论发布日期，格式 mm-dd
}

struct Video {
    1: required i64 id  // 视频唯一标识
    2: required User author  // 视频作者信息
    3: required string play_url  // 视频播放地址
    4: required string cover_url // 视频封面地址
    5: required i64 favorite_count  // 视频的点赞总数
    6: required i64 comment_count // 视频的评论总数
    7: required bool is_favorite // true-已点赞，false-未点赞
    8: required string title // 视频标题
}

service ApiService {
    CreateUserResponse CreateUser(1: CreateUserRequest req) (api.post="/douyin/user/register/")
    CheckUserResponse CheckUser(1: CheckUserRequest req) (api.post="/douyin/user/login/")
    GetUserResponse GetUser(1: GetUserRequest req) (api.get="/douyin/user/")
    ActionResponse Action (1:ActionRequest req)(api.post="/douyin/relation/action/")
    FollowListResponse FollowList (1:FollowListRequest req)(api.get="/douyin/relation/follow/list/")
    FollowerListResponse FollowerList (1:FollowerListRequest req)(api.get="/douyin/relation/follower/list/")
    FriendListResponse FriendList (1:FriendListRequest req)(api.get="/douyin/relation/friend/list/")
    ChatRecordResponse ChatRecord(1:ChatRecordRequest req)(api.get="/douyin/message/chat/")
    SendMessageResponse SendMessage (1:SendMessageRequest req)(api.post="/douyin/message/action/")

    Douyin_feed_response FeedVideo(1: Douyin_feed_request request) (api.get='/douyin/feed/')
    Douyin_publish_action_response PublishVideo(1: Douyin_publish_action_request request) (api.post='/douyin/publish/action/')
    Douyin_publish_list_response PublishListVideo (1: Douyin_publish_list_request request) (api.get = '/douyin/publish/list/')
    Douyin_favorite_action_response FavoriteAction (1: Douyin_favorite_action_request request) (api.post = '/douyin/favorite/action/')
    Douyin_favorite_list_response FavoriteList (1: Douyin_favorite_list_request request) (api.get = '/douyin/favorite/list/')
    Douyin_comment_action_response CommentAction (1: Douyin_comment_action_request request) (api.post = '/douyin/comment/action/')
    Douyin_comment_list_response CommentList(1: Douyin_comment_list_request request) (api.get = '/douyin/comment/list/')
}