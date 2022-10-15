import { useLocation } from 'react-router-dom'
import { ChatType } from '@/enums'
import Header from '@/components/Header'
import ChatList from '@/components/ChatList'
import ChatWindow from '@/components/ChatWindow'

import './index.scss'

function Chat() {
  const location = useLocation()
  const { pathname } = location

  return (
    <div className='chat'>
      {pathname !== '/group' ? (
        <Header name='华小科' online={true} />
      ) : (
        <Header name='芝士软工' peopleNum={5} />
      )}
      <div className='chat-main'>
        <ChatList />
        {pathname !== '/group' ? (
          <ChatWindow chatType={ChatType.Private} />
        ) : (
          <ChatWindow chatType={ChatType.Group} />
        )}
      </div>
    </div>
  )
}

export default Chat
