import { useAlert } from 'react-alert'
import ChatListStore from '@/mobx/chatlist'
import ChatStore from '@/mobx/chat'
import './index.scss'

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
      {ChatStore.currentChat?.gid && (
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
