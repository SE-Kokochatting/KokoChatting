import { makeAutoObservable } from 'mobx'
import { pullMsg } from '@/network/message/pullMsg'
import { pullHistoryMsg } from '@/network/message/pullHistoryMsg'
import { MessageType } from '@/enums'
import { IMessageContent, IMessage } from '@/types'
import { getMsgId, setMsgId } from '@/utils/message'
import { pullMsgOutline } from '@/network/message/pullMsgOutline'

class MsgState {
  // 好友请求
  public friendRequest: IMessageContent[] = []
  public friendRequestIsPull = false
  // 群通知
  public groupNotify: IMessageContent[] = []
  public groupNotifyIsPull = false
  // 好友消息
  public friendMsg: Partial<IMessage>[] = []
  public friendMsgIsPull = false
  // 群消息
  public groupMsg: Partial<IMessage>[] = []
  public groupMsgIsPull = false

  public constructor() {
    makeAutoObservable(this)
  }

  /**
   * 移除请求好友请求
   * @param mid 消息id
   * @returns void
   */
  public removeFriendRequest(mid: number) {
    this.friendRequest = this.friendRequest.filter(
      (item) => item.messageId !== mid,
    )
  }

  /**
   * 移除群通知
   * @param mid 消息id
   * @returns void
   */
  public removeGroupNotify(mid: number) {
    this.groupNotify = this.groupNotify.filter((item) => item.messageId !== mid)
  }

  /**
   * 拉取具体消息
   * @param msgType 消息类型
   * @returns void
   */
  public async pullMsgContent(msgType: MessageType) {
    if (msgType === MessageType.FriendRequestNotify) {
      if (this.friendRequestIsPull) return
      this.friendRequestIsPull = true
    } else if (msgType === MessageType.JoinGroupRequestNotify) {
      if (this.groupNotifyIsPull) return
      this.groupNotifyIsPull = true
    } else if (msgType === MessageType.SingleMessage) {
      if (this.friendMsgIsPull) return
      this.friendMsgIsPull = true
    } else if (msgType === MessageType.GroupMessage) {
      if (this.groupMsgIsPull) return
      this.groupMsgIsPull = true
    }

    // 获取消息概要数组
    const { data } = await pullMsgOutline({ lastMessageId: 0 })
    const { message } = data

    this.init()
    if (!message) return
    let maxMsgId = 0

    const reqArr = []

    // 拉取消息
    for (const outlineMsg of message) {
      if (outlineMsg.messageType === msgType) {
        reqArr.push(
          pullHistoryMsg({
            firstMessageId: 100000,
            id:
              outlineMsg.groupId === 0
                ? outlineMsg.senderId
                : outlineMsg.groupId,
            msgType: msgType,
            pageNum: 1,
            pageSize: 100000,
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
      } else if (
        msgType === MessageType.JoinGroupRequestNotify ||
        msgType === MessageType.QuitGroupNotify ||
        msgType === MessageType.JoinGroupNotify
      ) {
        this.groupNotify.push(message)
      } else if (msgType === MessageType.SingleMessage) {
        this.friendMsg.push(message)
      } else if (msgType === MessageType.GroupMessage) {
        this.groupMsg.push(message)
      }
    }
  }

  /**
   * 发送消息
   * @param message 发送的消息
   * @param msgType 消息类型
   * @returns void
   */
  public sendMsg(message: Partial<IMessage>, msgType: MessageType) {
    if (msgType === MessageType.SingleMessage) {
      this.friendMsg.push(message)
    } else if (msgType === MessageType.GroupMessage) {
      this.groupMsg.push(message)
    }
  }

  private init() {
    this.friendMsg = []
    this.groupMsg = []
    this.friendRequest = []
    this.groupNotify = []
  }
}

const MsgStore = new MsgState()

export default MsgStore
