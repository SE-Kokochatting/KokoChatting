import Header from '@/components/Header'
import ChatList from '@/components/ChatList'
import './index.scss'

function Home() {
  return (
    <div className='home'>
      <Header />
      <ChatList />
    </div>
  )
}

export default Home
