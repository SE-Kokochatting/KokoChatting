/**
 * description: 登录
 * author: Yuming Cui
 * date: 2022-10-28 23:35:52 +0800
 */

export interface ILogin {
  uid: string
  password: string
}

export async function login(data: ILogin) {
  const url = '/api/user/login'
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
