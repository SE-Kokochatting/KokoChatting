/**
 * description: 获取用户信息
 * author: Yuming Cui
 * date: 2022-10-28 23:35:52 +0800
 */

export interface IGetUserInfo {
  uid: number
}

export async function getUserInfo(data: IGetUserInfo): Promise<any> {
  const url = '/api/user'
  try {
    const res = await fetch(url, {
      method: 'POST',
      mode: 'cors',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    })
    return res.json()
  } catch (err) {
    console.error(err)
  }
}
