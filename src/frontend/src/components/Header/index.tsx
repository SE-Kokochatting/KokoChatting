import { observer } from 'mobx-react-lite'
import { useState } from 'react'
import { useShowDropDown } from './hooks/useShowDropdown'
import CurrentChatStore from '@/mobx/currentChat'
import SvgIcon from '@/components/SvgIcon'
import Search from './components/Search'
import LeftDropdown from './components/LeftDropdown'
import RightDropdown from './components/RightDropdown'
import './index.scss'

function _Header() {
  const [showLeftDropdown, setShowLeftDropdown] = useState(false)
  // right dropdown
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
            {CurrentChatStore.currentChat?.name}
          </span>
          {CurrentChatStore.currentChat?.count &&
            CurrentChatStore.currentChat?.gid && (
              <span className='c-header-right-info-num'>
                {CurrentChatStore.currentChat?.count} members
              </span>
            )}
          {CurrentChatStore.currentChat?.uid && (
            <span className='c-header-right-info-state'>online</span>
          )}
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
            cursor: 'pointer',

            '&:hover': {
              color: 'var(--global-font-primary)',
            },
          }}
          onClick={(e) => {
            e.stopPropagation()
            setShowDropDown(true)
          }}
        />
        <RightDropdown showDropdown={showDropDown} />
      </div>
    </div>
  )
}

const Header = observer(_Header)

export default Header
