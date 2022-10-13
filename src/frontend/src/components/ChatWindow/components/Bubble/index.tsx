import { Direction } from '@/enums'
import SvgIcon from '@/components/SvgIcon'
import './index.scss'

interface BubbleProps {
  content: string
  direction: Direction
  read?: boolean
  time: string
}

function Bubble(props: BubbleProps) {
  const { content, direction, read, time } = props
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
            fill: 'var(--light)',
            alignSelf: 'end',
            marginRight: '10px',
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
