import { makeAutoObservable } from 'mobx'
import { pullMsg } from '@/network/message/pullMsg'
import { MessageType } from '@/enums'
import { IMessageContent, IMessage } from '@/types'
import { setMsgId } from '@/utils/message'
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

    // const mid = getMsgId()
    const mid = 0
    // 获取消息概要数组
    const { data } = await pullMsgOutline({ lastMessageId: mid })
    const { message } = data

    this.init()
    if (!message) return

    let maxMsgId = 0

    const reqArr = []

    // 拉取消息
    for (const outlineMsg of message) {
      if (outlineMsg.messageType === msgType) {
        reqArr.push(
          pullMsg({
            // 由于在消息页面，拿好友申请举例，需要渲染出以前没有同意或拒绝的请求
            // 因此，需要从0开始请求，对于已处理的请求，后端不会返回给前端
            // 好友聊天界面中，拿到maxMsgId，只渲染大于此id的消息
            lastMessageId: 0,
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
    console.log(msgArr)

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
    setMsgId(maxMsgId)
  }

  // /**
  //  * 发送消息
  //  * @param val 要设置的值
  //  * @returns void
  //  */
  // public sendMsg() {

  // }

  private init() {
    this.friendMsg = []
    this.groupMsg = []
    this.friendRequest = []
    this.groupNotify = []
  }
}

const MsgStore = new MsgState()

export default MsgStore
