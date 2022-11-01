/**
 * description: chatList 状态管理
 * author: Yuming Cui
 * date: 2022-11-01 12:54:59 +0800
 */

import { makeAutoObservable } from 'mobx'
import { IGroup } from '@/types'
import { getGroupList } from '@/network/group/getGroupList'

class ChatListState {
  public data: IGroup[] = []

  public constructor() {
    makeAutoObservable(this)
  }

  /**
   * 请求群列表
   * @param val 要设置的值
   * @returns void
   */
  public updateGroup() {
    getGroupList().then(({ data }) => {
      const { group } = data
      this.data = group
    })
  }

  /**
   * 请求好友列表
   * @param val 要设置的值
   * @returns void
   */
  public updateFriend() {}
}

const ChatListStore = new ChatListState()

export default ChatListStore
