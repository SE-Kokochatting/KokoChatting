/**
 * description: 拉取历史消息
 * author: Yuming Cui
 * date: 2022-11-04 20:48:27 +0800
 */

import { MessageType } from '@/enums'
import request from '../request'

export interface IPullHistoryMsg {
  firstMessageId: number
  id: number
  msgType: MessageType
  pageNum: number
  pageSize: number
}

export async function pullHistoryMsg(reqData: IPullHistoryMsg): Promise<any> {
  const url = '/api/user/friend/history'
  const { code, data } = await request(
    url,
    { method: 'POST', useToken: true },
    reqData,
  )
  return { code, data }
}
