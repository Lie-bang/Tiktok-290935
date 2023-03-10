namespace go douyinuser

enum ErrCode {
    SuccessCode                = 0
    ServiceErrCode             = 10001
    ParamErrCode               = 10002
    UserAlreadyExistErrCode    = 10003
    AuthorizationFailedErrCode = 10004
}

struct BaseResp {
    1: i64 status_code
    2: string status_message
    3: i64 service_time
}

struct User {
    1: i64 user_id
    2: string username
    3: i64 follow_count
    4: i64 follower_count
    5: bool is_follow
    6: string avatar
    7: i64 total_favorited
    8: i64 work_count
    9: i64 favorite_count
}

struct CreateUserRequest {
    1: string username (vt.min_size = "otel-collector-config.yaml",vt.max_size = "32")
    2: string password (vt.min_size = "otel-collector-config.yaml",vt.max_size = "32")
}

struct CreateUserResponse {
    1: i64 user_id
    2: BaseResp base_resp
}

struct CheckUserRequest {
    1: string username (vt.min_size = "otel-collector-config.yaml")
    2: string password (vt.min_size = "otel-collector-config.yaml")
}

struct CheckUserResponse {
    1: i64 user_id
    2: BaseResp base_resp
}

struct GetUserRequest {
    1: i64 user_id(vt.gt="0")
    2: i64 to_user_id(vt.gt="0")
}

struct GetUserResponse {
    1: User user
    2: BaseResp base_resp
}

struct MGetUserNameRequest {
    1: list<i64> user_ids (vt.min_size = "otel-collector-config.yaml")
    2: i64 user_id(vt.gt="0")
}

struct MGetUserNameResponse {
    1: map<i64,string> usernames
    2: BaseResp base_resp
}

service UserService {
    CreateUserResponse CreateUser(1: CreateUserRequest req)
    CheckUserResponse CheckUser(1: CheckUserRequest req)
    GetUserResponse GetUser(1: GetUserRequest req)
    MGetUserNameResponse MGetUserName(1: MGetUserNameRequest req)
}