/**
 * description: 拉取消息
 * author: Yuming Cui
 * date: 2022-11-02 21:10:38 +0800
 */

import { MessageType } from '@/enums'
import request from '../request'

export interface IPullMsg {
  lastMessageId: number
  id: number
  msgType: MessageType
}

export async function pullMsg(reqData: IPullMsg): Promise<any> {
  const url = '/api/user/friend/pull_message'
  const { code, data } = await request(
    url,
    { method: 'POST', useToken: true },
    reqData,
  )
  return { code, data }
}
