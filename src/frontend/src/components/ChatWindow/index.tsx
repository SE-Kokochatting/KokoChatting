import { observer } from 'mobx-react-lite'
import { Direction, ChatType, Theme } from '@/enums'
import ThemeStore from '@/mobx/theme'
import Bubble from './components/Bubble'
import Sender from './components/Sender'
import './index.scss'

interface ChatWindowProps {
  chatType: ChatType
}

function _ChatWindow(props: ChatWindowProps) {
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
      <div
        className='c-chat_window-chat_area'
        style={{
          backgroundImage:
            ThemeStore.theme === Theme.Light
              ? 'url(https://linhong.me/2019/11/02/telegram-background/bg.jpg)'
              : 'none',
          backgroundPosition: 'center',
        }}
      >
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

const ChatWindow = observer(_ChatWindow)

export default ChatWindow
