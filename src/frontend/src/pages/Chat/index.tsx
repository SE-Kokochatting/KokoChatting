import { useLocation } from 'react-router-dom'
import { useEffect } from 'react'
import { observer } from 'mobx-react-lite'
import { ChatType, Theme } from '@/enums'
import Header from '@/components/Header'
import ChatList from '@/components/ChatList'
import ChatWindow from '@/components/ChatWindow'
import UserInfo from '@/components/UserInfo'
import Toggle from '@/components/Toggle'
import ThemeStore from '@/mobx/theme'
import ChatListStore from '@/mobx/chatlist'
import './index.scss'

function _Chat() {
  const location = useLocation()
  const { pathname } = location

  useEffect(() => {
    if (pathname === '/home') {
      ChatListStore.setChatType(ChatType.Message)
    } else if (pathname === '/private') {
      ChatListStore.setChatType(ChatType.Private)
    } else if (pathname === '/group') {
      ChatListStore.setChatType(ChatType.Group)
    }
  }, [pathname])

  // Todo: 窗口类型由此时打开的窗口决定，和 /group 无关
  return (
    <div className={ThemeStore.theme === Theme.Dark ? 'chat dark' : 'chat'}>
      <Header />
      <div className='chat-main'>
        <ChatList />
        <ChatWindow />
        <UserInfo />
      </div>
      <Toggle />
    </div>
  )
}

const Chat = observer(_Chat)

export default Chat
