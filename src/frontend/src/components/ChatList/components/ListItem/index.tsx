import { observer } from 'mobx-react-lite'
import { IChat } from '@/types'
import CurrentChatStore from '@/mobx/currentChat'
import './index.scss'

function handleClick({ uid, gid, avatarUrl, name, extract, lastTime }: IChat) {
  CurrentChatStore.setCurrentChat({
    uid,
    gid,
    avatarUrl,
    name,
    extract,
    lastTime,
  })
}

function _ListItem({ uid, gid, avatarUrl, name, extract, lastTime }: IChat) {
  return (
    <div
      className='c-chat_list-item'
      onClick={() =>
        handleClick({ uid, gid, avatarUrl, name, extract, lastTime })
      }
    >
      <div className='c-chat_list-item-avatar'>
        <img className='c-chat_list-item-avatar-img' src={avatarUrl} />
      </div>
      <div className='c-chat_list-item-main'>
        <span className='c-chat_list-item-main-name'>{name}</span>
        <span className='c-chat_list-item-main-content'>{extract}</span>
      </div>
      <div className='c-chat_list-item-time'>{lastTime}</div>
    </div>
  )
}

const ListItem = observer(_ListItem)

export default ListItem
