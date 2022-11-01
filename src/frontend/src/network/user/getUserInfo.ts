/**
 * description: 获取用户信息
 * author: Yuming Cui
 * date: 2022-10-28 23:35:52 +0800
 */

import request from '../request'

export interface IGetUserInfo {
  uid: number
}

export async function getUserInfo(reqData: IGetUserInfo): Promise<any> {
  const url = '/api/user'
  const { code, data } = await request(
    url,
    { method: 'POST', useToken: true },
    reqData,
  )
  return { code, data }
}
