/**
 * description: 当前聊天
 * author: Yuming Cui
 * date: 2022-11-01 14:33:52 +0800
 */

import { makeAutoObservable } from 'mobx'
import { quitGroup } from '@/network/group/quitGroup'

class ChatState {
  public currentChat: any = null

  public constructor() {
    makeAutoObservable(this)
  }

  /**
   * 设置当前显示的聊天
   * @param val 要设置的值
   * @returns void
   */
  public setCurrentChat(val: any) {
    this.currentChat = val
  }

  /**
   * 退出群
   * @param val 要设置的值
   * @returns void
   */
  public async quitGroup() {
    const { code, data } = await quitGroup({
      gid: this.currentChat?.gid as number,
    })
    return { code, data }
  }
}

const ChatStore = new ChatState()

export default ChatStore
