import { useAlert } from 'react-alert'
import ChatListStore from '@/mobx/chatlist'
import ChatStore from '@/mobx/chat'
import './index.scss'
import { delFriend } from '@/network/friend/delFriend'
import { blockFriend } from '@/network/friend/blockFriend'

interface RightDropdownProps {
  showDropdown: boolean
}

function RightDropdown({ showDropdown }: RightDropdownProps) {
  const alert = useAlert()

  async function handleQuit() {
    const { code } = await ChatStore.quitGroup()
    if (code === 200) {
      alert.show('已退群')
      ChatListStore.updateGroup()
      ChatStore.setCurrentChat(null)
    }else{
      alert.show('操作异常')
    }
  }

  async function deleteFriend() {
    if (ChatStore.currentChat === null){
      return
    }
    if (ChatStore.currentChat.uid === undefined){
      return
    }
    const { code } = await delFriend({fid:ChatStore.currentChat.uid})
    if(code === 200){
      // 刷新好友列表
      alert.show("删除成功")
      ChatListStore.updateFriend()
    }else{
      alert.show("删除好友失败")
    }
  }

  async function _blockFriend() {
    if (ChatStore.currentChat === null){
      return
    }
    if (ChatStore.currentChat.uid === undefined){
      return
    }
    const {code} = await blockFriend({fid:ChatStore.currentChat.uid})
    if(code === 200){
      // 刷新好友列表
      alert.show("拉黑成功")
      ChatListStore.updateFriend()
    }else{
      alert.show("拉黑好友失败")
    }
  }

  return (
    <div
      className='c-header-right-dropdown'
      style={{ display: showDropdown ? 'flex' : 'none' }}
      onClick={(e) => {
        e.stopPropagation()
      }}
    >
      {ChatStore.currentChat?.gid > 0 && (
        <div
          className='c-header-right-dropdown-item'
          onClick={() => handleQuit()}
        >
          退出群聊
        </div>
      )}

      {ChatStore.currentChat?.uid > 0 && (
        <>
          <div
            className='c-header-right-dropdown-item'
            onClick={() => _blockFriend()}
          >
            拉黑好友
          </div>
          <div
          className='c-header-right-dropdown-item'
          onClick={() => deleteFriend()}
          >
            删除好友
          </div>
        </>
      )}
    </div>
  )
}
export default RightDropdown
