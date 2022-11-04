/**
 * description: 当前聊天
 * author: Yuming Cui
 * date: 2022-11-01 14:33:52 +0800
 */

import { makeAutoObservable } from 'mobx'
import { quitGroup } from '@/network/group/quitGroup'
import { IChat, IMessage } from '@/types'

class ChatState {
  public currentChat: Partial<IChat> | null = null
  public bubblesData: IMessage[] = []

  public constructor() {
    makeAutoObservable(this)
  }

  /**
   * 设置当前显示的聊天
   * @param val 要设置的值
   * @returns void
   */
  public setCurrentChat(val: Partial<IChat> | null) {
    this.currentChat = val
  }

  /**
   * 设置气泡
   * @param val 要设置的值
   * @returns void
   */
  public setBubblesData(val: IMessage[]) {
    this.bubblesData = val
  }

  /**
   * 退出群
   * @param
   * @returns void
   */
  public async quitGroup() {
    const { code, data } = await quitGroup({
      gid: this.currentChat?.gid as number,
    })
    return { code, data }
  }

  /**
   * 删除好友
   * @param
   * @returns void
   */
  public async deleteFriend() {}

  /**
   * 拉黑好友
   * @param
   * @returns void
   */
  public async blockFriend() {}
}

const ChatStore = new ChatState()

export default ChatStore
