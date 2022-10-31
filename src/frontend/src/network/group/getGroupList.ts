/**
 * description: 获取群列表
 * author: Yuming Cui
 * date: 2022-10-31 18:37:08 +0800
 */

export async function getGroupList(token: string): Promise<any> {
  const url = '/api/group/list'
  try {
    const res = await fetch(url, {
      method: 'GET',
      mode: 'cors',
      headers: {
        Authorization: token,
        'Content-Type': 'application/json',
      },
    })
    return res.json()
  } catch (err) {
    console.error(err)
  }
}
