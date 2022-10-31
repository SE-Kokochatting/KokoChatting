/**
 * description: 保存和获取用户token
 * author: Yuming Cui
 * date: 2022-10-28 20:52:13 +0800
 */

const TOKEN_KEY = 'token'

export function setToken(token: string) {
  window.localStorage.setItem(TOKEN_KEY, token)
}

export function clearToken() {
  window.localStorage.removeItem(TOKEN_KEY)
}

export function getToken(): string | null {
  const tokenStr = window.localStorage.getItem(TOKEN_KEY)
  if (!tokenStr) return null

  return tokenStr
}
