import Header from '@/components/Header'
import ChatList from '@/components/ChatList'
import './index.scss'
import ChatWindow from '@/components/ChatWindow'

function Home() {
  return (
    <div className='home'>
      <Header />
      <div className='home-main'>
        <ChatList />
        <ChatWindow />
      </div>
    </div>
  )
}

export default Home
