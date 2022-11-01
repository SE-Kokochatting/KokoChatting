import { useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import { useAlert } from 'react-alert'
import { ChatType } from '@/enums'
import { IGroup } from '@/types'
import { DefaultGroupAvatar } from '@/consts'
import { getGroupList } from '@/network/group/getGroupList'
import { getToken } from '@/utils/token'
import ListItem from './components/ListItem'
import Loading from '@/components/Loading'
import './index.scss'

interface ChatListProps {
  chatType: ChatType
}

function ChatList({ chatType }: ChatListProps) {
  const navigate = useNavigate()
  const alert = useAlert()
  const [chatListData, setChatListData] = useState<IGroup[]>([])
  const [isLoading, setIsLoading] = useState(true)

  useEffect(() => {
    setIsLoading(true)
    const token = getToken()
    if (!token) {
      alert.show('请先登录', {
        title: '创建群失败',
        onClose: () => {
          navigate('/login')
        },
      })
      return
    } else {
      switch (chatType) {
        case ChatType.Mixed:
          setIsLoading(false)
          setChatListData([])
          break
        case ChatType.Private:
          setIsLoading(false)
          setChatListData([])
          break
        case ChatType.Group:
          getGroupList(token).then((res) => {
            const resData = res.data
            const { data } = resData
            const { group } = data
            setChatListData(group)
            setIsLoading(false)
          })
      }
    }
  }, [chatType])

  return (
    <div className='c-chat_list'>
      {chatListData.map(({ gid, avatarUrl, name, extract, lastTime }) => (
        <ListItem
          key={gid}
          avatarUrl={avatarUrl ? avatarUrl : DefaultGroupAvatar}
          name={name}
          extract={extract}
          lastTime={lastTime}
        />
      ))}
      {isLoading && <Loading />}
    </div>
  )
}
export default ChatList
