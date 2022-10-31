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
  Mixed,
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
}
