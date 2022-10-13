import Bubble from './components/Bubble'
import Sender from './components/Sender'
import { Direction } from '@/enums'
import './index.scss'
// type ChatWindowProps = {
// };
function ChatWindow(/* props: ChatWindowProps */) {
  // const {} = props;
  return (
    <div className='c-chat_window'>
      <div className='c-chat_window-chat_area'>
        <Bubble
          content='70周年校庆快乐！'
          direction={Direction.Left}
          time='8:01'
        />
        <Bubble
          content='同乐！'
          direction={Direction.Right}
          read={true}
          time='8:02'
        />
        <Bubble
          content='70周年校庆快乐！'
          direction={Direction.Left}
          time='8:03'
        />
      </div>
      <Sender />
    </div>
  )
}
export default ChatWindow
