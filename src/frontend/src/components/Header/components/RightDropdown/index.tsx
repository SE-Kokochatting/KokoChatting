import { useAlert } from 'react-alert'
import ChatListStore from '@/mobx/chatList'
import CurrentChatStore from '@/mobx/currentChat'
import './index.scss'

interface RightDropdownProps {
  showDropdown: boolean
}

function RightDropdown({ showDropdown }: RightDropdownProps) {
  const alert = useAlert()

  async function handleQuit() {
    const { code } = await CurrentChatStore.quitGroup()
    if (code === 200) {
      alert.show('已退群')
      ChatListStore.updateGroup()
      CurrentChatStore.setCurrentChat(null)
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
      {CurrentChatStore.currentChat?.gid && (
        <div
          className='c-header-right-dropdown-item'
          onClick={() => handleQuit()}
        >
          退出群聊
        </div>
      )}
    </div>
  )
}
export default RightDropdown
