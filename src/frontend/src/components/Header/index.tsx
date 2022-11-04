import { observer } from 'mobx-react-lite'
import { useState } from 'react'
import { useShowDropDown } from './hooks/useShowDropdown'
import ChatStore from '@/mobx/chat'
import SvgIcon from '@/components/SvgIcon'
import Search from './components/Search'
import LeftDropdown from './components/LeftDropdown'
import RightDropdown from './components/RightDropdown'
import './index.scss'

function _Header() {
  const [showLeftDropdown, setShowLeftDropdown] = useState(false)
  // Right dropdown
  const { showDropDown, setShowDropDown } = useShowDropDown()

  return (
    <div className='c-header'>
      <div className='c-header-left'>
        <SvgIcon
          name='menu'
          style={{
            color: 'var(--global-font-primary_lighter)',
            width: '35px',
            height: '35px',
            marginLeft: '10px',
            cursor: 'pointer',
          }}
          onClick={() => {
            setShowLeftDropdown(!showLeftDropdown)
          }}
        />
        <LeftDropdown
          showLeftDropdown={showLeftDropdown}
          setShowLeftDropdown={setShowLeftDropdown}
        />
        <Search />
      </div>
      <div className='c-header-right'>
        <div className='c-header-right-info'>
          <span className='c-header-right-info-name'>
            {ChatStore.currentChat?.name}
          </span>
          {ChatStore.currentChat?.count && ChatStore.currentChat?.gid && (
            <span className='c-header-right-info-num'>
              {ChatStore.currentChat?.count} members
            </span>
          )}
          {/* {ChatStore.currentChat?.uid && (
            <span className='c-header-right-info-state'>online</span>
          )} */}
        </div>
        <SvgIcon
          name='search'
          style={{
            color: 'var(--global-font-primary_lighter)',
            width: '35px',
            height: '35px',
            position: 'absolute',
            right: '70px',
            cursor: 'pointer',
          }}
        />
        <SvgIcon
          name='right-menu'
          style={{
            color: 'var(--global-font-primary_lighter)',
            width: '35px',
            height: '35px',
            position: 'absolute',
            right: '20px',
            cursor: ChatStore.currentChat !== null ? 'pointer' : 'default',

            '&:hover': {
              color: 'var(--global-font-primary)',
            },
          }}
          onClick={(e) => {
            if (ChatStore.currentChat !== null) {
              e.stopPropagation()
              setShowDropDown(true)
            }
          }}
        />
        <RightDropdown showDropdown={showDropDown} />
      </div>
    </div>
  )
}

const Header = observer(_Header)

export default Header
