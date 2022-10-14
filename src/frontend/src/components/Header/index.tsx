import { useState } from 'react'
import SvgIcon from '@/components/SvgIcon'
import Search from './components/Search'
import LeftDropdown from './components/LeftDropdown'
import './index.scss'

interface HeaderProps {
  name: string
  online: boolean
  // todo
  // interval?: number
  peopleNum?: number
}

function Header(props: HeaderProps) {
  const { name, peopleNum } = props
  const [showLeftDropdown, setShowLeftDropdown] = useState(false)

  return (
    <div className='c-header'>
      <div className='c-header-left'>
        <SvgIcon
          name='menu'
          style={{
            fill: '#fff',
            width: '35px',
            height: '35px',
            marginLeft: '10px',
            cursor: 'pointer',
          }}
          onClick={() => {
            setShowLeftDropdown(!showLeftDropdown)
          }}
        />
        <LeftDropdown shown={showLeftDropdown} />
        <Search />
      </div>
      <div className='c-header-right'>
        <div className='c-header-right-info'>
          <span className='c-header-right-info-name'>{name}</span>
          {peopleNum ? (
            <span className='c-header-right-info-num'>{peopleNum} members</span>
          ) : (
            <span className='c-header-right-info-state'>online</span>
          )}
        </div>
        <SvgIcon
          name='search'
          style={{
            fill: '#fff',
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
            fill: '#fff',
            width: '35px',
            height: '35px',
            position: 'absolute',
            right: '20px',
            cursor: 'pointer',
          }}
        />
      </div>
    </div>
  )
}
export default Header
