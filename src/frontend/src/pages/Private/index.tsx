import Header from '@/components/Header'
import ChatList from '@/components/ChatList'
import ChatWindow from '@/components/ChatWindow'
import { ChatType } from '@/enums'
import './index.scss'

function Private() {
  return (
    <div className='private'>
      <Header name='华小科' online={true} />
      <div className='private-main'>
        <ChatList />
        <ChatWindow chatType={ChatType.Private} />
      </div>
    </div>
  )
}

export default Private
