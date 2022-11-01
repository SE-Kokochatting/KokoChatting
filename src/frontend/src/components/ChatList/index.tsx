import { observer } from 'mobx-react-lite'
import { useState, useEffect } from 'react'
import { ChatType } from '@/enums'
import { DefaultGroupAvatar } from '@/consts'
import { IChat } from '@/types'
import ChatListStore from '@/mobx/chatList'
import ListItem from './components/ListItem'
import Loading from '@/components/Loading'
import './index.scss'

interface ChatListProps {
  chatType: ChatType
}

function _ChatList({ chatType }: ChatListProps) {
  const [isLoading, setIsLoading] = useState(true)

  async function handleFetchData() {
    setIsLoading(true)

    if (chatType === ChatType.Mixed) {
      //
    } else if (chatType === ChatType.Private) {
      //
    } else {
      await ChatListStore.updateGroup()
    }
    setIsLoading(false)
  }

  useEffect(() => {
    handleFetchData()
  }, [chatType])

  return (
    <div className='c-chat_list'>
      {ChatListStore.data.map(
        ({ uid, gid, avatarUrl, name, extract, lastTime }: IChat) => (
          <ListItem
            key={uid ? `u${uid}` : `g${gid}`}
            uid={uid}
            gid={gid}
            avatarUrl={avatarUrl ? avatarUrl : DefaultGroupAvatar}
            name={name}
            extract={extract}
            lastTime={lastTime}
          />
        ),
      )}
      {isLoading && <Loading />}
    </div>
  )
}

const ChatList = observer(_ChatList)

export default ChatList
