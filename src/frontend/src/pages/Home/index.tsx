import Header from '@/components/Header'
import ChatList from '@/components/ChatList'
import ChatWindow from '@/components/ChatWindow'
import { ChatType } from '@/enums'
import './index.scss'

function Private() {
  return (
    <div className='home'>
      <Header name='华小科' online={true} />
      <div className='home-main'>
        <ChatList />
        <ChatWindow chatType={ChatType.Private} />
      </div>
    </div>
  )
}

export default Private
