/**
 * description: 封装 fetch
 * 1. 自动带上 token
 * 2. token 过期，实现重定向
 * author: Yuming Cui
 * date: 2022-10-31 22:27:09 +0800
 */

import { getToken } from '@/utils/token'

interface IOptions {
  method: 'GET' | 'POST'
  useToken: boolean
}

export default async function request(
  url: string,
  options: IOptions,
  reqData?: any,
): Promise<any> {
  const { method, useToken } = options
  const headers: any = {
    'Content-Type': 'application/json',
  }
  if (useToken) {
    headers.Authorization = getToken()
  }
  try {
    const res = await fetch(url, {
      method,
      headers,
      mode: 'cors',
      body: JSON.stringify(reqData),
    })
    const resData = await res.json()
    // console.log(resData)

    // 后端发送的数据中 data 套 data
    const { code, data } = resData
    const realData = data ? data.data : null
    // token 错误或已过期
    if (code === 1007) {
      window.location.href = '/login'
    }
    return { code, data: realData }
  } catch (err) {
    console.error(err)
  }
}
