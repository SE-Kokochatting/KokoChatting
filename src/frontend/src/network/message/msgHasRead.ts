/**
 * description: 消息已读未读
 * author: Yuming Cui
 * date: 2022-11-05 16:18:30 +0800
 */

import request from '../request'

export interface IMsgHasRead {
  msgids: number[]
}

export async function msgHasRead(reqData: IMsgHasRead): Promise<any> {
  const url = '/api/user/friend/read_message'
  const { code, data } = await request(
    url,
    { method: 'POST', useToken: true },
    reqData,
  )
  return { code, data }
}
