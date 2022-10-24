import { Direction, ChatType } from '@/enums'
import SvgIcon from '@/components/SvgIcon'
import './index.scss'

interface BubbleProps {
  content: string
  chatType: ChatType
  direction: Direction
  read?: boolean
  time: string
}

function Bubble(props: BubbleProps) {
  const { content, chatType, direction, read, time } = props
  return (
    <div
      style={{
        alignSelf: direction === Direction.Left ? 'start' : 'end',
        margin:
          direction === Direction.Left
            ? '15px 0 15px 30px'
            : '15px 30px 15px 0',
        maxWidth: '50%',
        display: 'flex',
      }}
    >
      {read && (
        <SvgIcon
          name='done'
          style={{
            width: '25px',
            height: '25px',
            color: 'var(--light)',
            alignSelf: 'end',
            marginRight: '10px',
          }}
        />
      )}
      {chatType === ChatType.Group && direction === Direction.Left && (
        <img
          src='https://p.qqan.com/up/2021-2/16137992359659254.jpg'
          style={{
            width: '50px',
            height: '50px',
            borderRadius: '50%',
            alignSelf: 'center',
            marginRight: '20px',
            cursor: 'pointer',
          }}
        />
      )}
      <div
        className={
          direction === Direction.Left
            ? 'c-chat_window-chat_area-bubble'
            : 'c-chat_window-chat_area-bubble right'
        }
      >
        <div className='c-chat_window-chat_area-bubble-content'>{content}</div>
        <span className='c-chat_window-chat_area-bubble-time'>{time}</span>
      </div>
    </div>
  )
}
export default Bubble
