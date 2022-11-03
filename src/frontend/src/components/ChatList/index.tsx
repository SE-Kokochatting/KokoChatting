import { observer } from 'mobx-react-lite'
import { useState, useEffect } from 'react'
import { ChatType } from '@/enums'
import { DefaultGroupAvatar } from '@/consts'
import { IMessage, IGroup } from '@/types'
import ChatListStore from '@/mobx/chatlist'
import ListItem from './components/ListItem'
import Loading from '@/components/Loading'
import './index.scss'

function _ChatList() {
  const [isLoading, setIsLoading] = useState(true)

  async function handleFetchData() {
    setIsLoading(true)

    if (ChatListStore.chatType === ChatType.Message) {
      await ChatListStore.updateMsgOutline()
    } else if (ChatListStore.chatType === ChatType.Private) {
      await ChatListStore.updateFriend()
    } else {
      await ChatListStore.updateGroup()
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
          }: IMessage) => (
            <ListItem
              key={senderId ? `u${senderId}` : `g${groupId}`}
              uid={senderId}
              gid={groupId}
              messageType={messageType}
              messageNum={messageNum}
              lastMessageTime={lastMessageTime}
            />
          ),
        )}
      {ChatListStore.chatType === ChatType.Group &&
        ChatListStore.groupData !== null &&
        ChatListStore.groupData.map(({ gid, avatarUrl, name }: IGroup) => (
          <ListItem
            key={`g${gid}`}
            gid={gid}
            avatarUrl={avatarUrl ? avatarUrl : DefaultGroupAvatar}
            name={name}
          />
        ))}
      {isLoading && <Loading />}
    </div>
  )
}

const ChatList = observer(_ChatList)

export default ChatList
