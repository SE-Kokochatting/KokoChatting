import request from '../request'

export interface IModifyUserAvatar {
  avatarUrl: string
}

export async function ModifyUserAvatar(
  reqData: IModifyUserAvatar,
): Promise<any> {
  const url = '/api/user/avatar'
  const { code, data } = await request(
    url,
    { method: 'POST', useToken: true },
    reqData,
  )
  return { code, data }
}
