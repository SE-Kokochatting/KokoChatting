import { observer } from 'mobx-react-lite'
import { IChat } from '@/types'
import { MessageType, ChatType } from '@/enums'
import ChatStore from '@/mobx/chat'
import { transformTimestamp } from '@/utils/date'
import './index.scss'

function handleClick({ uid, gid, avatarUrl, name }: Partial<IChat>) {
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
  chatType,
}: Partial<IChat> & { chatType: ChatType }) {
  return (
    <div
      className='c-chat_list-item'
      style={{
        display:
          (chatType === ChatType.Message &&
            (messageType === MessageType.SingleMessage ||
              messageType === MessageType.GroupMessage)) ||
          chatType !== ChatType.Message
            ? 'flex'
            : 'none',
      }}
      onClick={() => handleClick({ uid, gid, avatarUrl, name, messageType })}
    >
      <div
        className='c-chat_list-item-avatar'
        style={{
          backgroundImage: `url(${avatarUrl})`,
          backgroundSize: 'cover',
        }}
      />
      <div className='c-chat_list-item-main'>
        <span className='c-chat_list-item-main-name'>{name}</span>
        {/* <span className='c-chat_list-item-main-content'>{extract}</span> */}
      </div>
      <div className='c-chat_list-item-time'>
        {transformTimestamp(lastMessageTime)}
      </div>
    </div>
  )
}

const ListItem = observer(_ListItem)

export default ListItem
