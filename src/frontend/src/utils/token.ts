const TOKEN_KEY = '__koko_chating_token'

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
