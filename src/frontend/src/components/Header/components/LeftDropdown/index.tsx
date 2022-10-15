import { useNavigate } from 'react-router-dom'
import React from 'react'
import SvgIcon from '@/components/SvgIcon'
import Switch from './components/Switch'
import './index.scss'

interface LeftDropdownProps {
  shown: boolean
}

const iconStyle: React.CSSProperties = {
  width: '25px',
  height: '25px',
  fill: 'var(--light)',
  verticalAlign: 'middle',
  marginRight: '20px',
}

function LeftDropdown(props: LeftDropdownProps) {
  const { shown } = props
  const navigate = useNavigate()
  return (
    <ul
      className='c-header-left-dropdown'
      style={{ display: shown ? 'block' : 'none' }}
    >
      <li
        className='c-header-left-dropdown-item'
        onClick={() => {
          navigate('/home')
        }}
      >
        <SvgIcon name='message' style={iconStyle} />
        最近消息
      </li>
      <li
        className='c-header-left-dropdown-item'
        onClick={() => {
          navigate('/private')
        }}
      >
        <SvgIcon name='contact' style={iconStyle} />
        联系人
      </li>
      <li
        className='c-header-left-dropdown-item'
        onClick={() => {
          navigate('/group')
        }}
      >
        <SvgIcon name='group' style={iconStyle} />
        群组
      </li>
      <li className='c-header-left-dropdown-item'>
        <SvgIcon name='notice' style={iconStyle} />
        通知
      </li>
      <li className='c-header-left-dropdown-item'>
        <SvgIcon name='myself' style={iconStyle} />
        个人信息
      </li>
      <li className='c-header-left-dropdown-item' id='dark_mode'>
        <SvgIcon name='moon' style={iconStyle} />
        暗黑模式
        <Switch />
      </li>
    </ul>
  )
}
export default LeftDropdown
