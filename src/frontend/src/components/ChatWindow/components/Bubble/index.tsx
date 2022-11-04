import { ChatType } from '@/enums'
import { IMessage } from '@/types'
import { getUid } from '@/utils/uid'
// import SvgIcon from '@/components/SvgIcon'
import './index.scss'

function Bubble({
  lastMessageTime,
  readUids,
  messageContent,
  senderId,
  chatType,
}: Partial<IMessage> & { chatType: ChatType }) {
  const uid = getUid()
  return (
    <div
      style={{
        alignSelf: uid === senderId ? 'start' : 'end',
        margin: uid === senderId ? '15px 0 15px 30px' : '15px 30px 15px 0',
        maxWidth: '50%',
        display: 'flex',
      }}
    >
      {/* {read && (
        <SvgIcon
          name='done'
          style={{
            width: '25px',
            height: '25px',
            color: 'var(--global-font-primary_lighter)',
            alignSelf: 'end',
            marginRight: '10px',
          }}
        />
      )} */}
      {chatType === ChatType.Group && uid !== senderId && (
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
          uid === senderId
            ? 'c-chat_window-chat_area-bubble'
            : 'c-chat_window-chat_area-bubble right'
        }
      >
        <div className='c-chat_window-chat_area-bubble-content'>
          {messageContent}
        </div>
        <span className='c-chat_window-chat_area-bubble-time'>
          {lastMessageTime}
        </span>
      </div>
    </div>
  )
}
export default Bubble
