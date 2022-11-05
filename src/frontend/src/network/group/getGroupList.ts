/**
 * description: 获取群列表
 * author: Yuming Cui
 * date: 2022-10-31 18:37:08 +0800
 */

import request from '../request'

export async function getGroupList(): Promise<any> {
  const url = '/api/group/list'
  const { code, data } = await request(url, { method: 'GET', useToken: true })
  return { code, data }
}
