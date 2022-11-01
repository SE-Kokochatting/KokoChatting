import { useState, useEffect } from 'react'
// import { useAlert } from 'react-alert'
import { ChatType } from '@/enums'
import { DefaultGroupAvatar } from '@/consts'
import ChatListStore from '@/mobx/chatList'
import ListItem from './components/ListItem'
import Loading from '@/components/Loading'
import './index.scss'

interface ChatListProps {
  chatType: ChatType
}

function ChatList({ chatType }: ChatListProps) {
  // const alert = useAlert()
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
      {ChatListStore.data.map(({ gid, avatarUrl, name, extract, lastTime }) => (
        <ListItem
          key={gid}
          avatarUrl={avatarUrl ? avatarUrl : DefaultGroupAvatar}
          name={name}
          extract={extract}
          lastTime={lastTime}
        />
      ))}
      {isLoading && <Loading />}
    </div>
  )
}
export default ChatList
