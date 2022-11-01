/**
 * description: 当前聊天
 * author: Yuming Cui
 * date: 2022-11-01 14:33:52 +0800
 */

import { makeAutoObservable } from 'mobx'
import { IChat } from '@/types'

class CurrentChatState {
  public currentChat: IChat | null = null

  public constructor() {
    makeAutoObservable(this)
  }

  /**
   * 设置当前显示的聊天
   * @param val 要设置的值
   * @returns void
   */
  public setCurrentChat(val: IChat) {
    this.currentChat = val
  }

  /**
   * 退出群
   * @param val 要设置的值
   * @returns void
   */
  public quitGroup() {}
}

const CurrentChatStore = new CurrentChatState()

export default CurrentChatStore
