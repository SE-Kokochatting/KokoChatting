/**
 * description: 保存和获取用户uid
 * author: Yuming Cui
 * date: 2022-10-29 20:52:25 +0800
 */

const UID_KEY = 'uid'

export function setUid(uid: number) {
  window.localStorage.setItem(UID_KEY, `${uid}`)
}

export function clearUid() {
  window.localStorage.removeItem(UID_KEY)
}

export function getUid(): number | null {
  const uidStr = window.localStorage.getItem(UID_KEY)
  if (!uidStr) return null

  return parseInt(uidStr, 10)
}
