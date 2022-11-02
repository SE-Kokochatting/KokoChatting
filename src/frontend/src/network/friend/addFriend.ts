/**
 * description: 请求添加好友
 * author: Yuming Cui
 * date: 2022-11-01 19:51:42 +0800
 */

import { MessageType } from '@/enums'
import request from '../request'

export interface IAddFriend {
  receiver: number
  messageContent: string
  messageType: MessageType
}

export async function addFriend(reqData: IAddFriend): Promise<any> {
  const url = '/api/user/friend/send_message'
  const { code, data } = await request(
    url,
    { method: 'POST', useToken: true },
    reqData,
  )
  return { code, data }
}
