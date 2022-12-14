import { observer } from 'mobx-react-lite'
import { useAlert } from 'react-alert'
import { acceptFriend } from '@/network/friend/acceptFriend'
import { refuseFriend } from '@/network/friend/refuseFriend'
import { MessageType } from '@/enums'
import MsgStore from '@/mobx/msg'
import './index.scss'

interface NotifyItemProps {
  publisherName: string
  info: string
  mid: number
  type: MessageType
}

function _NotifyItem({ publisherName, info, mid, type }: NotifyItemProps) {
  const alert = useAlert()

  function agree() {
    if (type === MessageType.FriendRequestNotify) {
      acceptFriend({ id: mid }).then((res) => {
        if (res.code === 200) {
          MsgStore.removeFriendRequest(mid)
          alert.show('添加好友成功')
        } else {
          alert.show('添加好友失败')
        }
      })
    }
  }

  function refuse() {
    refuseFriend({ id: mid }).then((res) => {
      if (res.code === 200) {
        MsgStore.removeFriendRequest(mid)
      } else {
        alert.show('操作失败')
      }
    })
  }

  return (
    <div className='notify_item'>
      <div className='notify_item-ctn'>
        <span className='notify_item-ctn-span'>
          {publisherName} 申请添加您为好友
        </span>
        <button className='notify_item-ctn-btn agree' onClick={agree}>
          同意
        </button>
        <button className='notify_item-ctn-btn refuse' onClick={refuse}>
          拒绝
        </button>
      </div>
      <div className='notify_item-info'>{info}</div>
    </div>
  )
}

const NotifyItem = observer(_NotifyItem)
export default NotifyItem
