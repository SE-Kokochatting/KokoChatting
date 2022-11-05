/**
 * description: 枚举
 * author: Yuming Cui
 * date: 2022-10-13 12:46:06 +0800
 */

// 气泡方向
export enum Direction {
  Left,
  Right,
}

// 聊天类型
export enum ChatType {
  Message,
  Private,
  Group,
}

// 主题
export enum Theme {
  Light = 'light',
  Dark = 'dark',
}

// Toggle 类型
export enum ToggleType {
  AddContact,
  CreateGroup,
  Notify,
}

// 消息类型
export enum MessageType {
  // 0 单聊消息
  SingleMessage,
  // 1 群聊消息
  GroupMessage,
  // 2 好友请求
  FriendRequestNotify,
  // 3 撤回单聊消息通知
  RevertSingleMessageNotify,
  // 4 撤回群聊消息通知
  RevertGroupMessageNotify,
  // 5 单聊消息消息已读通知
  HasReadSingleNotify,
  // 6 群聊消息已读通知
  HasReadGroupNotify,
  // 7 入群申请通知
  JoinGroupRequestNotify,
  // 8 退群通知
  QuitGroupNotify,
  // 9 进群通知
  JoinGroupNotify,
  // 10 同意好友请求通知
  AddFriendResponseNotify,
  // 11 删除好友通知
  DeleteFriendNotify,
  // 12 心跳包
  PongMessage,
}

export enum AddType {
  Friend,
  Group,
}

export enum NotifyType {
  friendRequest,
  groupManageNotify,
}
