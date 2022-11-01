import { observer } from 'mobx-react-lite'
import { CSSObject } from '@emotion/react'
import { useNavigate } from 'react-router-dom'
import { useAlert } from 'react-alert'
import { getUserInfo } from '@/network/user/getUserInfo'
import { getUid } from '@/utils/uid'
import { ToggleType } from '@/enums'
import UserStore from '@/mobx/user'
import ToggleStore from '@/mobx/toggle'
import SvgIcon from '@/components/SvgIcon'
import Switch from './components/Switch'
import './index.scss'

interface LeftDropdownProps {
  showLeftDropdown: boolean
  setShowLeftDropdown: (val: boolean) => void
}

const iconStyle: CSSObject = {
  width: '25px',
  height: '25px',
  color: 'var(--global-font-primary_lighter)',
  verticalAlign: 'middle',
  marginRight: '20px',
}

async function handleUserInfo(uid: number) {
  const { data } = await getUserInfo({ uid })
  const { name, avatarUrl } = data
  UserStore.setUserInfo({ uid, name, avatarUrl })
  UserStore.setShowUserInfo(!UserStore.showUserInfo)
}

function handleToggle(type: ToggleType) {
  ToggleStore.setShowToggle(true)
  ToggleStore.setToggleType(type)
}

function _LeftDropdown({
  showLeftDropdown,
  setShowLeftDropdown,
}: LeftDropdownProps) {
  const navigate = useNavigate()
  const alert = useAlert()
  return (
    <ul
      className='c-header-left-dropdown'
      style={{ display: showLeftDropdown ? 'block' : 'none' }}
    >
      <li
        className='c-header-left-dropdown-item'
        onClick={() => {
          setShowLeftDropdown(false)
          navigate('/home')
        }}
      >
        <SvgIcon name='message' style={iconStyle} />
        最近消息
      </li>
      <li
        className='c-header-left-dropdown-item'
        onClick={() => {
          setShowLeftDropdown(false)
          navigate('/private')
        }}
      >
        <SvgIcon name='contact' style={iconStyle} />
        联系人
      </li>
      <li
        className='c-header-left-dropdown-item'
        onClick={() => {
          setShowLeftDropdown(false)
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
      <li
        className='c-header-left-dropdown-item'
        onClick={() => {
          const uid = getUid()
          if (!uid) {
            alert.show('请尝试重新登录！', {
              title: '操作失败',
              onClose: () => {
                navigate('/login')
              },
            })
          } else {
            handleUserInfo(uid)
          }
        }}
      >
        <SvgIcon name='myself' style={iconStyle} />
        个人信息
      </li>
      <li
        className='c-header-left-dropdown-item'
        onClick={() => {
          handleToggle(ToggleType.AddContact)
        }}
      >
        <SvgIcon name='addContact' style={iconStyle} />
        添加联系人
      </li>
      <li
        className='c-header-left-dropdown-item'
        onClick={() => {
          handleToggle(ToggleType.CreateGroup)
        }}
      >
        <SvgIcon name='createGroup' style={iconStyle} />
        创建群聊
      </li>
      <li className='c-header-left-dropdown-item' id='dark_mode'>
        <SvgIcon name='moon' style={iconStyle} />
        暗黑模式
        <Switch />
      </li>
    </ul>
  )
}

const LeftDropdown = observer(_LeftDropdown)

export default LeftDropdown
