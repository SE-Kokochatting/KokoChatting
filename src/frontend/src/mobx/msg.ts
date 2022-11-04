import { makeAutoObservable } from 'mobx'
import { pullMsg } from '@/network/message/pullMsg'
import { MessageType } from '@/enums'
import { IMessageContent } from '@/types'
import { getMsgId, setMsgId } from '@/utils/message'
import { pullMsgOutline } from '@/network/message/pullMsgOutline'

class MsgState {
  // 好友请求
  public friendRequest: IMessageContent[] = []
  public friendIsPull = false
  // 群通知
  public groupNotify: IMessageContent[] = []
  public groupIsPull = false

  public constructor() {
    makeAutoObservable(this)
  }

  public removeFriendRequest(mid: number) {
    this.friendRequest = this.friendRequest.filter(
      (item) => item.messageId !== mid,
    )
  }

  public removeGroupRequest(mid: number) {
    this.groupNotify = this.groupNotify.filter((item) => item.messageId !== mid)
  }

  /**
   * 拉取具体消息
   * @param val 要设置的值
   * @returns void
   */
  public async pullMsgContent(msgType: MessageType) {
    if (msgType === MessageType.FriendRequestNotify) {
      if (this.friendIsPull) return
      this.friendIsPull = true
    } else if (msgType === MessageType.JoinGroupRequestNotify) {
      if (this.groupIsPull) return
      this.groupIsPull = true
    }

    const mid = getMsgId()
    // 获取消息概要数组
    const { data } = await pullMsgOutline({ lastMessageId: mid })
    const { message } = data

    this.init()
    if (!message) return

    let maxMsgId = 0

    const reqArr = []

    for (const outlineMsg of message) {
      if (outlineMsg.messageType === msgType) {
        reqArr.push(
          pullMsg({
            lastMessageId: mid,
            id:
              outlineMsg.groupId === 0
                ? outlineMsg.senderId
                : outlineMsg.groupId,
            msgType: msgType,
          }),
        )
      }
    }

    const resData = await Promise.all(reqArr)
    const msgArr = resData.map((item: any) => item.data.message).flat()

    for (const message of msgArr) {
      maxMsgId = Math.max(maxMsgId, message.messageId)
      if (msgType === MessageType.FriendRequestNotify) {
        this.friendRequest.push(message)
      } else {
        this.groupNotify.push(message)
      }
    }
    setMsgId(maxMsgId)
  }

  private init() {
    this.friendRequest = []
    this.groupNotify = []
  }
}

const MsgStore = new MsgState()

export default MsgStore
