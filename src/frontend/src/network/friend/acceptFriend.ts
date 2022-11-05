/**
 * description: 同意添加好友
 * author: Yuming Cui
 * date: 2022-11-01 19:56:41 +0800
 */

import request from '../request'

export interface IAcceptFriend {
  // 消息的 id
  id: number
}

export async function acceptFriend(reqData: IAcceptFriend): Promise<any> {
  const url = '/api/user/add_friend'
  const { code, data } = await request(
    url,
    { method: 'POST', useToken: true },
    reqData,
  )
  return { code, data }
}
