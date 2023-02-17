namespace go douyinrelation

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

struct FriendUser {
    1: required User user
    2: optional string message
    3: required i64 msgType
}

struct ActionRequest{
    1: required i64 user_id(vt.gt="0")
    2: required i64 to_user_id(vt.gt="0")
    3: required i32 action_type(vt.in="1",vt.in="2")
}

struct ActionResponse{
    1: required BaseResp base_resp
}

struct FollowListRequest{
    1: required i64 user_id(vt.gt="0")
    2: required i64 to_user_id(vt.gt="0")
}

struct FollowListResponse{
    1: required BaseResp base_resp
    2: required list<User> user_list
}

struct FollowerListRequest{
    1: required i64 user_id(vt.gt="0")
    2: required i64 to_user_id(vt.gt="0")
 }

struct FollowerListResponse{
    1: required BaseResp base_resp
    2: required list<User> user_list
}

struct FriendListRequest{
    1: required i64 user_id(vt.gt="0")
    2: required i64 to_user_id(vt.gt="0")
}

struct FriendListResponse{
    1: required BaseResp base_resp
    2: required list<FriendUser> user_list
}

struct GetRelationInfoRequest{
    1: required i64 user_id(vt.gt="0")
    2: required list<i64> to_user_ids
}

struct GetRelationInfoResponse{
    1: required BaseResp base_resp
    2: required list<User> user_list
}

service RelationService{
    ActionResponse Action (1:ActionRequest req)
    FollowListResponse FollowList (1:FollowListRequest req)
    FollowerListResponse FollowerList (1:FollowerListRequest req)
    FriendListResponse FriendList (1:FriendListRequest req)
    GetRelationInfoResponse GetRelationInfo(1:GetRelationInfoRequest req)
}