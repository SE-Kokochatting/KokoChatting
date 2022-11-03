import { observer } from 'mobx-react-lite'
import { IChat } from '@/types'
import { MessageType } from '@/enums'
import ChatStore from '@/mobx/chat'
import './index.scss'

function handleClick({ uid, gid, avatarUrl, name }: any) {
  ChatStore.setCurrentChat({
    uid,
    gid,
    avatarUrl,
    name,
  })
}

function _ListItem({
  uid,
  gid,
  avatarUrl,
  name,
  messageType,
  messageNum,
  lastMessageTime,
}: IChat) {
  return (
    <div
      className='c-chat_list-item'
      style={{
        display:
          messageType === MessageType.SingleMessage ||
          messageType === MessageType.GroupMessage
            ? 'flex'
            : 'none',
      }}
      onClick={() => handleClick({ uid, gid, avatarUrl, name })}
    >
      <div className='c-chat_list-item-avatar'>
        <img className='c-chat_list-item-avatar-img' src={avatarUrl} />
      </div>
      <div className='c-chat_list-item-main'>
        <span className='c-chat_list-item-main-name'>{name}</span>
        {/* <span className='c-chat_list-item-main-content'>{extract}</span> */}
      </div>
      <div className='c-chat_list-item-time'>{lastMessageTime}</div>
    </div>
  )
}

const ListItem = observer(_ListItem)

export default ListItem
