namespace go douyinfavorite

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

struct Douyin_favorite_judge_request {
    1: required i64 user_id
    2: required i64 video_id
}

struct Douyin_favorite_judge_response {
    1: required i32 status_code // 状态码，0-成功，其他值-失败
    2: optional string status_msg  // 返回状态描述
    3: required i32 if_fav // 表明是否点赞
}

struct Douyin_favorite_count_user_request {
    1: required i64 user_id
}

struct Douyin_favorite_count_user_response {
    1: required i32 status_code
    2: required i64 favorite_count
    3: optional string status_msg
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

struct User {
  1: required i64 id  // 用户id
  2: required string name  // 用户名称
  3: optional i64 follow_count // 关注总数
  4: optional i64 follower_count  // 粉丝总数
  5: required bool is_follow  // true-已关注，false-未关注
  6: optional string avatar //用户头像
  7: optional string background_image  //用户个人页顶部大图
  8: optional string signature  //个人简介
  9: optional i64 total_favorited  //获赞数量
  10: optional i64 work_count  //作品数量
  11: optional i64 favorite_count  //点赞数量
}


service FavoriteService {
    Douyin_favorite_action_response FavoriteAction (1: Douyin_favorite_action_request request)
    Douyin_favorite_list_response FavoriteList (1: Douyin_favorite_list_request request)
    Douyin_favorite_judge_response FavoriteJudge (1: Douyin_favorite_judge_request request)
    Douyin_favorite_count_user_response FavoriteCountUser (1: Douyin_favorite_count_user_request request)

}