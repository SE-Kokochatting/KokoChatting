import { observer } from 'mobx-react-lite'
import { DefaultAvatarUrl } from '@/consts'
import UserStore from '@/mobx/user'
import SvgIcon from '../SvgIcon'
import './index.scss'

function _UserInfo() {
  return (
    <div
      className='user_info'
      style={{ maxWidth: UserStore.showUserInfo ? '400px' : 0 }}
    >
      <h1 className='user_info-title'>
        <SvgIcon
          name='cross'
          style={{
            color: 'var(--global-font-primary)',
            width: '20px',
            height: '20px',
            marginRight: '10px',
            cursor: 'pointer',

            '&:hover': {
              color: 'var(--global-font-primary_lighter)',
            },
          }}
          onClick={() => {
            UserStore.setShowUserInfo(false)
          }}
        />
        用户信息
      </h1>
      <img
        className='user_info-avatar'
        src={UserStore.avatarUrl ? UserStore.avatarUrl : DefaultAvatarUrl}
      />
      <p className='user_info-item'>uid: {UserStore.uid}</p>
      <p className='user_info-item'>用户名: {UserStore.name}</p>
    </div>
  )
}

const UserInfo = observer(_UserInfo)

export default UserInfo
