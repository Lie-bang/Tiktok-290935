namespace go douyinmessage

struct BaseResp {
    1: required i32 status_code
    2: optional string status_message
    3: required i64 service_time
}

struct Message {
    1: required i64 id
    2: required i64 to_user_id
    3: required i64 from_user_id
    4: required string content
    5: optional i64 create_time
}

struct ChatRecordRequest{
    1: required i64 to_user_id(vt.gt="0")
    2: required i64 user_id(vt.gt="0")
}

struct ChatRecordResponse{
    1: required BaseResp base_resp
    2: required list<Message> msg_list
}

struct SendMessageRequest{
    1: required i64 to_user_id(vt.gt="0")
    2: required i64 user_id(vt.gt="0")
    3: required i64 action_type(vt.in="1")
    4: required string content
}

struct SendMessageResponse{
    1: required BaseResp base_resp
}

struct GetFirstMessagesRequest{
    1: required list<i64> to_user_ids
    2: required i64 user_id(vt.gt="0")
}

struct GetFirstMessagesResponse{
    1: required BaseResp base_resp
    2: required list<Message> messages
}

service MessageService{
    ChatRecordResponse ChatRecord(1:ChatRecordRequest req)
    SendMessageResponse SendMessage (1:SendMessageRequest req)
    GetFirstMessagesResponse GetFirstMessages(1:GetFirstMessagesRequest req)
}