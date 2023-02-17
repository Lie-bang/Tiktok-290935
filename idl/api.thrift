namespace go douyinapi

struct User {
    1: required i64 id
    2: required string name
    3: optional i64 follow_count
    4: optional i64 follower_count
    5: required bool is_follow
    6: optional string avatar
    7: optional string background_image
    8: optional string signature
    9: optional i64 total_favorited
    10: optional i64 work_count
    11: optional i64 favorite_count
}

struct FriendUser {
    1: required i64 id
    2: required string name
    3: optional i64 follow_count
    4: optional i64 follower_count
    5: required bool is_follow
    6: optional string avatar
    7: optional string background_image
    8: optional string signature
    9: optional i64 total_favorited
    10: optional i64 work_count
    11: optional i64 favorite_count
    12: optional string message
    13: required i64 msgType
}

struct Message {
    1: required i64 id
    2: required i64 to_user_id
    3: required i64 from_user_id
    4: required string content
    5: optional i64 create_time
}
struct BaseResponse{
    1: required i32 status_code
    2: optional string status_msg
}

struct ApiUserResponse{
    1: required i32 status_code
    2: optional string status_msg
    3: required User user
}

struct ApiUsersResponse{
    1: required i32 status_code
    2: optional string status_msg
    3: required list<User> user_list
}

struct ApiFriendUsersResponse{
    1: required i32 status_code
    2: optional string status_msg
    3: required list<FriendUser> user_list
}

struct ApiMessageResponse{
    1: required i32 status_code
    2: optional string status_msg
    3: required list<Message> message_list
}


struct CreateUserRequest {
    1: required string username (api.query="username", api.vd="len($) > 0")
    2: required string password (api.query="password", api.vd="len($) > 0")
}

struct CreateUserResponse {
    1: required i32 status_code
    2: optional string status_msg
    3: required i64 user_id
    4: required string token
}

struct CheckUserRequest {
    1: required string username (api.query="username", api.vd="len($) > 0")
    2: required string password (api.query="password", api.vd="len($) > 0")
}

struct CheckUserResponse {
    1: required i32 status_code
    2: optional string status_msg
    3: required i64 user_id
    4: required string token
}

struct GetUserRequest {
    1: required i64 user_id
    2: required string token
}

struct GetUserResponse {
    1: required i32 status_code
    2: optional string status_msg
    3: required User user
}

struct ActionRequest{
    1: required i64 to_user_id
    2: required i32 action_type
    3: required string token
}

struct ActionResponse{
    1: required i32 status_code
    2: optional string status_msg
}

struct FollowListRequest{
    1: required i64 user_id
    2: required string token
}

struct FollowListResponse{
    1: required i32 status_code
    2: optional string status_msg
    3: required list<User> user_list
}

struct FollowerListRequest{
     1: required i64 user_id
     2: required string token
 }

struct FollowerListResponse{
    1: required i32 status_code
    2: optional string status_msg
    3: required list<User> user_list
}

struct FriendListRequest{
    1: required i64 user_id
    2: required string token
}

struct FriendListResponse{
    1: required i32 status_code
    2: optional string status_msg
    3: required list<FriendUser> user_list
}

struct ChatRecordRequest{
    1: i64 to_user_id
    2: required i64 pre_msg_time
    3: required string token
}

struct ChatRecordResponse{
    1: required i32 status_code
    2: optional string status_msg
    3: required list<Message> message_list
}

struct SendMessageRequest{
    1: required i64 to_user_id
    2: required i64 action_type
    3: required string content
    4: required string token
}

struct SendMessageResponse{
    1: required i32 status_code
    2: optional string status_msg
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
}