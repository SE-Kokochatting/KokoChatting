import { observer } from 'mobx-react-lite'
import { DefaultAvatarUrl } from '@/consts'
import UserStore from '@/mobx/user'
import SvgIcon from '../SvgIcon'
import { useAlert } from 'react-alert'
import  { ModifyUserAvatar }  from "@/network/user/modifyUserAvatar"
import './index.scss'

function _UserInfo() {
  const alert = useAlert()
  function upload(e: any){
    const formData: FormData = new FormData()
    console.log("upload ",e.target.files[0])
    formData.append("file",e.target.files[0])
    fetch("/api/upload",{
      method:"POST",
      body: formData
    }).then(async (res) => {
      const resData = await res.json()
      console.log(resData)
      const { code, data } = resData
      if(code !== 200){
        console.log(code,data)
        alert.show("上传失败")
        return
      }
      const url = data.data.url
      ModifyUserAvatar({avatarUrl:url}).then(({ code,data }) => {
        if(code === 200){
          UserStore.setUserInfo({ uid:UserStore.uid,name:UserStore.name,avatarUrl:url })
        }else{
          alert.show("修改头像失败")
        }
      })
    })

  }
  return (
    <div
      className='c-user_info'
      style={{ maxWidth: UserStore.showUserInfo ? '400px' : 0 }}
    >
      <h1 className='c-user_info-title'>
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
        className='c-user_info-avatar'
        src={UserStore.avatarUrl ? UserStore.avatarUrl : DefaultAvatarUrl}
      />

      <p className='c-user_info-item'>账号: {UserStore.uid}</p>
      <p className='c-user_info-item'>用户名: {UserStore.name}</p>
      <div className="c-user_info-btn">
        修改头像
        <input className="c-user_info-btn-uploader" type="file" accept="image/gif, image/jpeg, image/png, image/jpg" onChange={upload}/>
      </div>
    </div>
  )
}

const UserInfo = observer(_UserInfo)

export default UserInfo
