/**
 * description: 消息工具函数
 * author: Yuming Cui
 * date: 2022-11-02 21:47:38 +0800
 */

const TOKEN_KEY = 'last_msg_id'

export function setMsgId(mid: number) {
  window.localStorage.setItem(TOKEN_KEY, `${mid}`)
}

export function clearMsgId() {
  window.localStorage.removeItem(TOKEN_KEY)
}

export function getMsgId(): number {
  const mid = window.localStorage.getItem(TOKEN_KEY)
  if (!mid) return 0

  return parseInt(mid, 10)
}
