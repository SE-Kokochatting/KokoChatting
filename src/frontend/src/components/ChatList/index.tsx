import { observer } from 'mobx-react-lite'
import { useState, useEffect } from 'react'
import { ChatType, MessageType } from '@/enums'
import { DefaultGroupAvatarUrl, DefaultAvatarUrl } from '@/consts'
import { IGroup, IMessageOutline, IUser } from '@/types'
import ChatListStore from '@/mobx/chatlist'
import MsgStore from '@/mobx/msg'
import ListItem from './components/ListItem'
import Loading from '@/components/Loading'
import './index.scss'

function _ChatList() {
  const [isLoading, setIsLoading] = useState(true)

  function handleFetchData() {
    setIsLoading(true)

    if (ChatListStore.chatType === ChatType.Message) {
      ChatListStore.updateMsgOutline()
      MsgStore.pullMsgContent(MessageType.SingleMessage)
      // Todo: 也需要拉取群相关的消息
    } else if (ChatListStore.chatType === ChatType.Private) {
      ChatListStore.updateFriend()
    } else {
      ChatListStore.updateGroup()
    }

    setIsLoading(false)
  }

  useEffect(() => {
    handleFetchData()
  }, [ChatListStore.chatType])

  return (
    <div className='c-chat_list'>
      {ChatListStore.chatType === ChatType.Message &&
        ChatListStore.msgData !== null &&
        ChatListStore.msgData.map(
          ({
            senderId,
            groupId,
            messageType,
            messageNum,
            lastMessageTime,
            name,
            avatarUrl,
          }: Partial<IMessageOutline>) => (
            <ListItem
              key={senderId ? `u${senderId}` : `g${groupId}`}
              uid={senderId}
              gid={groupId}
              messageType={messageType}
              messageNum={messageNum}
              lastMessageTime={lastMessageTime}
              chatType={ChatType.Message}
              name={name}
              avatarUrl={
                avatarUrl
                  ? avatarUrl
                  : senderId
                  ? DefaultAvatarUrl
                  : DefaultGroupAvatarUrl
              }
            />
          ),
        )}
      {ChatListStore.chatType === ChatType.Private &&
        ChatListStore.friendData !== null &&
        ChatListStore.friendData.map(({ uid, avatarUrl, name }: IUser) => (
          <ListItem
            key={`u${uid}`}
            uid={uid}
            avatarUrl={avatarUrl ? avatarUrl : DefaultAvatarUrl}
            name={name}
            chatType={ChatType.Private}
          />
        ))}
      {ChatListStore.chatType === ChatType.Group &&
        ChatListStore.groupData !== null &&
        ChatListStore.groupData.map(({ gid, avatarUrl, name }: IGroup) => (
          <ListItem
            key={`g${gid}`}
            gid={gid}
            avatarUrl={avatarUrl ? avatarUrl : DefaultGroupAvatarUrl}
            name={name}
            chatType={ChatType.Group}
          />
        ))}
      {isLoading && <Loading />}
    </div>
  )
}

const ChatList = observer(_ChatList)

export default ChatList
