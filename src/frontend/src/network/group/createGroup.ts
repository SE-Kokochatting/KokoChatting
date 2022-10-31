/**
 * description: 创建群聊
 * author: Yuming Cui
 * date: 2022-10-31 17:03:51 +0800
 */

export interface ICreateGroup {
  // 群名称
  name: string
  // 用户 id 列表
  member?: {
    mid: number
  }[]
}

export async function createGroup(
  data: ICreateGroup,
  token: string,
): Promise<any> {
  const url = '/api/create_group'
  try {
    const res = await fetch(url, {
      method: 'POST',
      mode: 'cors',
      headers: {
        Authorization: token,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    })
    return res.json()
  } catch (err) {
    console.error(err)
  }
}
