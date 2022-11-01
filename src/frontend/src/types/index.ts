/**
 * description: 类型
 * author: Yuming Cui
 * date: 2022-10-31 20:39:10 +0800
 */

export interface IUser {
  uid: number
  name: string
  avatarUrl: string
}

export interface IGroup {
  gid: number
  avatarUrl: string
  name: string
  extract?: string
  lastTime?: string
  count?: number
}

export type IChat = IUser & IGroup
