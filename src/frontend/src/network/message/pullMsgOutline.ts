/**
 * description: 拉取消息纲要
 * author: Yuming Cui
 * date: 2022-11-02 21:39:50 +0800
 */

import request from '../request'

export interface IPullMsgOutline {
  lastMessageId: number
}

export async function pullMsgOutline(reqData: IPullMsgOutline): Promise<any> {
  const url = '/api/user/friend/pull_msg_outline'
  const { code, data } = await request(
    url,
    { method: 'POST', useToken: true },
    reqData,
  )
  return { code, data }
}
