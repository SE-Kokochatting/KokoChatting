/**
 * description: 请求添加好友
 * author: Yuming Cui
 * date: 2022-11-01 19:51:42 +0800
 */

import request from '../request'

export interface IBlockFriend {
  fid: number
}

export async function blockFriend(reqData: IBlockFriend): Promise<any> {
  const url = '/api/user/list_block'
  const { code, data } = await request(
    url,
    { method: 'POST', useToken: true },
    reqData,
  )
  return { code, data }
}
