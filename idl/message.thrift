namespace go douyinmessage

struct BaseResp {
    1: i64 status_code
    2: string status_message
    3: i64 service_time
}

struct Message {
    1: i64 msg_id
    2: i64 to_user_id
    3: i64 from_user_id
    4: string content
    5: i64 create_time
}

struct ChatRecordRequest{
    1: i64 to_user_id(vt.gt="0")
    2: i64 user_id(vt.gt="0")
}

struct ChatRecordResponse{
    1: BaseResp base_resp
    2: list<Message> msg_list
}

struct SendMessageRequest{
    1: i64 to_user_id(vt.gt="0")
    2: i64 user_id(vt.gt="0")
    3: i64 action_type(vt.in="1")
    4: string content
}

struct SendMessageResponse{
    1: BaseResp base_resp
}

struct GetFirstMessagesRequest{
    1: list<i64> to_user_ids
    2: i64 user_id(vt.gt="0")
}

struct GetFirstMessagesResponse{
    1: BaseResp base_resp
    2: list<Message> messages
}

service MessageService{
    ChatRecordResponse ChatRecord(1:ChatRecordRequest req)
    SendMessageResponse SendMessage (1:SendMessageRequest req)
    GetFirstMessagesResponse GetFirstMessages(1:GetFirstMessagesRequest req)
}