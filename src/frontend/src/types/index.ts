/**
 * description: 类型
 * author: Yuming Cui
 * date: 2022-10-31 20:39:10 +0800
 */

import { MessageType } from '@/enums'

export interface IUser {
  uid: number
  name: string
  avatarUrl: string
}

export interface IGroup {
  gid: number
  avatarUrl: string
  name: string
  count?: number
}

export interface IMessage {
  senderId: number
  groupId: number
  messageType: MessageType
  messageNum: number
  lastMessageTime: string
}

export type IChat = IUser & IGroup & IMessage
