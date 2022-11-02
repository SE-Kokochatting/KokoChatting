import { observer } from 'mobx-react-lite'
import { Direction } from '@/enums'
import ChatListStore from '@/mobx/chatlist'
import Bubble from './components/Bubble'
import Sender from './components/Sender'
import './index.scss'

function _ChatWindow() {
  // 之后定义其类型
  const chatInfo = [
    {
      id: 1,
      content: '70周年校庆快乐！',
      direction: Direction.Left,
      time: '8:01',
    },
    {
      id: 2,
      content: '同乐！',
      direction: Direction.Right,
      time: '8:02',
      read: true,
    },
    {
      id: 3,
      content: '70周年校庆快乐！',
      direction: Direction.Left,
      time: '8:03',
    },
  ]

  return (
    <div className='c-chat_window'>
      <div className='c-chat_window-chat_area'>
        {chatInfo.map(({ id, content, direction, time, read }) => (
          <Bubble
            key={id}
            chatType={ChatListStore.chatType}
            content={content}
            direction={direction}
            read={read}
            time={time}
          />
        ))}
      </div>
      <Sender />
    </div>
  )
}

const ChatWindow = observer(_ChatWindow)

export default ChatWindow
