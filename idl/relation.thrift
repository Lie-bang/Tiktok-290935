namespace go douyinrelation

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
}

struct FriendUser {
    1: User user
    2: string message
    3: i64 msgType
}

struct ActionRequest{
    1: i64 user_id(vt.gt="0")
    2: i64 to_user_id(vt.gt="0")
    3: i32 action_type(vt.in="1",vt.in="2")
}

struct ActionResponse{
    1: BaseResp base_resp
}

struct FollowListRequest{
    1: i64 user_id(vt.gt="0")
}

struct FollowListResponse{
    1: BaseResp base_resp
    2: list<User> user_list
}

struct FollowerListRequest{
    1: i64 user_id(vt.gt="0")
 }

struct FollowerListResponse{
    1: BaseResp base_resp
    2: list<User> user_list
}

struct FriendListRequest{
    1: i64 user_id(vt.gt="0")
}

struct FriendListResponse{
    1: BaseResp base_resp
    2: list<FriendUser> user_list
}

struct CountFollowRequest{
    1: i64 user_id(vt.gt="0")
}

struct CountFollowResponse{
    1: BaseResp base_resp
    2: i64 follow_count
}

struct CountFollowerRequest{
    1: i64 user_id(vt.gt="0")
}

struct CountFollowerResponse{
    1: BaseResp base_resp
    2: i64 follower_count
}

struct IsFollowRequest{
    1: i64 user_id(vt.gt="0")
    2: i64 to_user_id(vt.gt="0")
}

struct IsFollowResponse{
    1: BaseResp base_resp
    2: bool is_follow
}

service RelationService{
    ActionResponse Action (1:ActionRequest req)
    FollowListResponse FollowList (1:FollowListRequest req)
    FollowerListResponse FollowerList (1:FollowerListRequest req)
    FriendListResponse FriendList (1:FriendListRequest req)
    CountFollowResponse CountFollow(1:CountFollowRequest req)
    CountFollowerResponse CountFollower(1:CountFollowerRequest req)
    IsFollowResponse IsFollow(1:IsFollowRequest req)
}