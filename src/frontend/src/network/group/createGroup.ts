/**
 * description: 创建群聊
 * author: Yuming Cui
 * date: 2022-10-31 17:03:51 +0800
 */

import request from '../request'

export interface ICreateGroup {
  // 群名称
  name: string
  // 用户 id 列表
  member?: {
    mid: number
  }[]
}

export async function createGroup(reqData: ICreateGroup): Promise<any> {
  const url = '/api/create_group'
  const { code, data } = await request(
    url,
    { method: 'POST', useToken: true },
    reqData,
  )
  return { code, data }
}
