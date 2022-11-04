import { makeAutoObservable } from 'mobx'
import { pullMsg } from '@/network/message/pullMsg'
import {  MessageType } from '@/enums'
import { IMessageContent } from '@/types'
import { getMsgId } from '@/utils/message'
// import  ChatListStore  from '@/mobx/chatList'
import { pullMsgOutline } from '@/network/message/pullMsgOutline'

class MsgState{

  public friendRequest: IMessageContent[] = []
  public friendIsPull: boolean = false
  public groupNotify: IMessageContent[] = []
  public groupIsPull: boolean = false

  public constructor(){
    makeAutoObservable(this)
  }


  public removeFriendRequest(mid: number){
    this.friendRequest = this.friendRequest.filter(item => item.messageId != mid)
  }


  public removeGroupRequest(mid: number){
    this.groupNotify = this.groupNotify.filter(item => item.messageId != mid)
  }

  /**
   * 拉取具体消息
   * @param val 要设置的值
   * @returns void
   */
  public async pullMsgContent(msgType: MessageType) {
    if(msgType === MessageType.FriendRequestNotify){
      if(this.friendIsPull){
        return
      }
      this.friendIsPull = true
    }else if(msgType === MessageType.JoinGroupRequestNotify){
      if(this.groupIsPull){
        return
      }
      this.groupIsPull = true
    }
    const mid = getMsgId()
    const {data} = await pullMsgOutline({lastMessageId:mid})
    const {message} = data
    this.friendRequest = []
    this.groupNotify = []
    for(const outlineMsg of message){
      if(outlineMsg.messageType === msgType){
        pullMsg({ lastMessageId: mid, id: outlineMsg.groupId === 0 ? outlineMsg.senderId : outlineMsg.groupId, msgType: msgType }).then(({ data }) => {
          const { message } = data
          if(msgType === MessageType.FriendRequestNotify){
            this.friendRequest.push(...message)
          }else{
            this.groupNotify.push(...message)
          }
        })
      }
    }
  }
}


const MsgStore = new MsgState

export default MsgStore
