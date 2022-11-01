/**
 * description: 注册
 * author: Yuming Cui
 * date: 2022-10-28 23:35:52 +0800
 */

import request from '../request'

export interface IRegister {
  name: string
  password: string
}

export async function register(reqData: IRegister): Promise<any> {
  const url = '/api/user/register'
  const { code, data } = await request(
    url,
    { method: 'POST', useToken: false },
    reqData,
  )
  return { code, data }
}
