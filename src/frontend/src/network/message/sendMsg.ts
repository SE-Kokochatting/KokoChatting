/**
 * description: 发送消息
 * author: Yuming Cui
 * date: 2022-11-04 17:03:07 +0800
 */

import { MessageType } from '@/enums'
import request from '../request'

export interface ISendMsg {
  receiver: number
  messageContent: string
  messageType: MessageType
}

export async function sendMsg(reqData: ISendMsg): Promise<any> {
  const url = '/api/user/friend/send_message'
  const { code, data } = await request(
    url,
    { method: 'POST', useToken: true },
    reqData,
  )
  return { code, data }
}
