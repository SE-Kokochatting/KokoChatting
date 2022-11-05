import request from '../request'

export interface IRefuseFriend {
  // 消息的 id
  id: number
}

export async function refuseFriend(reqData: IRefuseFriend): Promise<any> {
  const url = '/api/user/add_friend_refuse'
  const { code, data } = await request(
    url,
    { method: 'POST', useToken: true },
    reqData,
  )
  return { code, data }
}
