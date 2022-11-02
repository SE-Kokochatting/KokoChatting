/**
 * description: 消息状态管理
 * author: Yuming Cui
 * date: 2022-11-01 12:54:59 +0800
 */

import { makeAutoObservable } from 'mobx'
import { ChatType } from '@/enums'
import { IGroup, IMessage } from '@/types'
import { getMsgId } from '@/utils/message'
import { getGroupList } from '@/network/group/getGroupList'
import { pullMsgOutline } from '@/network/message/pullMsgOutline'

class ChatListState {
  public chatType: ChatType = ChatType.Message
  public groupData: IGroup[] = []
  public msgData: IMessage[] = []

  public constructor() {
    makeAutoObservable(this)
  }

  /**
   * 设置类型
   * @param val 要设置的值
   * @returns void
   */
  public setChatType(val: ChatType) {
    this.chatType = val
  }

  /**
   * 请求群列表
   * @param val 要设置的值
   * @returns void
   */
  public updateGroup() {
    getGroupList().then(({ data }) => {
      const { group } = data
      this.groupData = group
    })
  }

  /**
   * 请求好友列表
   * @param val 要设置的值
   * @returns void
   */
  public updateFriend() {}

  /**
   * 请求消息纲要
   * @param val 要设置的值
   * @returns void
   */
  public updateMsgOutline() {
    const mid = getMsgId()
    pullMsgOutline({ lastMessageId: mid }).then(({ data }) => {
      const { message } = data
      this.msgData = message
    })
  }
}

const ChatListStore = new ChatListState()

export default ChatListStore
