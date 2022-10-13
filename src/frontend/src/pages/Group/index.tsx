import Header from '@/components/Header'
import ChatList from '@/components/ChatList'
import ChatWindow from '@/components/ChatWindow'
import { ChatType } from '@/enums'
import './index.scss'

function Group() {
  return (
    <div className='group'>
      <Header name='芝士软工' online={true} peopleNum={5} />
      <div className='group-main'>
        <ChatList />
        <ChatWindow chatType={ChatType.Group} />
      </div>
    </div>
  )
}

export default Group
