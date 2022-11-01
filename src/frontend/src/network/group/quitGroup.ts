/**
 * description: 退群
 * author: Yuming Cui
 * date: 2022-11-01 15:23:05 +0800
 */

import request from '../request'

export interface IQuitGroup {
  gid: number
}

export async function quitGroup(reqData: IQuitGroup): Promise<any> {
  const url = '/api/group/quit'
  const { code, data } = await request(
    url,
    { method: 'POST', useToken: true },
    reqData,
  )
  return { code, data }
}
