import Bubble from './components/Bubble'
import Sender from './components/Sender'
import { Direction, ChatType } from '@/enums'
import './index.scss'

interface ChatWindowProps {
  chatType: ChatType
}

function ChatWindow(props: ChatWindowProps) {
  const { chatType } = props
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
            chatType={chatType}
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
export default ChatWindow
