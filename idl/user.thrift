namespace go douyinuser

enum ErrCode {
    SuccessCode                = 0
    ServiceErrCode             = 10001
    ParamErrCode               = 10002
    UserAlreadyExistErrCode    = 10003
    AuthorizationFailedErrCode = 10004
}

struct BaseResp {
    1: required i32 status_code
    2: optional string status_message
    3: required i64 service_time
}

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
struct CreateUserRequest {
    1: required string username (vt.min_size = "1",vt.max_size = "32")
    2: required string password (vt.min_size = "1",vt.max_size = "32")
}

struct CreateUserResponse {
    1: required i64 user_id
    2: required BaseResp base_resp
}

struct CheckUserRequest {
    1: required string username (vt.min_size = "1",vt.max_size = "32")
    2: required string password (vt.min_size = "1",vt.max_size = "32")
}

struct CheckUserResponse {
    1: required i64 user_id
    2: required BaseResp base_resp
}

struct GetUserRequest {
    1: required i64 user_id(vt.gt="0")
    2: required i64 to_user_id(vt.gt="0")
}

struct GetUserResponse {
    1: required User user
    2: required BaseResp base_resp
}

struct MGetUserNameRequest {
    1: required list<i64> user_ids (vt.min_size = "1")
    2: required i64 user_id(vt.gt="0")
}

struct MGetUserNameResponse {
    1: required map<i64,string> usernames
    2: required BaseResp base_resp
}

service UserService {
    CreateUserResponse CreateUser(1: CreateUserRequest req)
    CheckUserResponse CheckUser(1: CheckUserRequest req)
    GetUserResponse GetUser(1: GetUserRequest req)
    MGetUserNameResponse MGetUserName(1: MGetUserNameRequest req)
}