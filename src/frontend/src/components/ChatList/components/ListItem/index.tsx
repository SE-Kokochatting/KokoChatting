import './index.scss'

interface ListItemProps {
  // 必需，如果使用默认头像，后端也应提供默认头像的 url
  avatarUrl: string
  // 用户名 / 群名称
  name: string
  // 摘要，即最新消息；如果是查看好友列表或群列表的场景，则无需此值
  extract?: string
  // 时；如果是查看好友列表或群列表的场景，则无需此值
  lastTime?: string
}
function ListItem(props: ListItemProps) {
  const { avatarUrl, name, extract, lastTime } = props
  return (
    <div className='c-chat_list-item'>
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
export default ListItem
