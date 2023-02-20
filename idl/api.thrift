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