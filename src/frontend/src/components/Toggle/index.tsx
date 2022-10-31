import { observer } from 'mobx-react-lite'
import { useNavigate } from 'react-router-dom'
import { useForm } from 'react-hook-form'
import { useAlert } from 'react-alert'
import { ToggleType } from '@/enums'
import { ICreateGroup, createGroup } from '@/network/group/createGroup'
import { getToken } from '@/utils/token'
import ToggleStore from '@/mobx/toggle'
import SvgIcon from '../SvgIcon'
import './index.scss'

function _Toggle() {
  const {
    register,
    reset,
    handleSubmit,
    formState: { errors },
  } = useForm<ICreateGroup>()

  const navigate = useNavigate()
  const alert = useAlert()
  const onSubmit = async (data: any) => {
    const token = getToken()
    if (!token) {
      alert.show('请先登录', {
        title: '创建群失败',
        onClose: () => {
          reset()
          navigate('/login')
        },
      })
      return
    } else {
      const res = await createGroup(data, token)
      const resData = res.data
      const _data = resData.data
      if (!_data) {
        const { code } = resData
        switch (code) {
          case 2003:
            alert.show('该群名称已存在', {
              title: '创建群失败',
              onClose: () => {
                reset()
              },
            })
            return
        }
      } else {
        alert.show('创建群聊成功', {
          onClose: () => {
            reset()
          },
        })
      }
    }
  }

  return (
    <div
      className='c-toggle'
      style={{ display: ToggleStore.showToggle ? 'block' : 'none' }}
    >
      <SvgIcon
        name='cross'
        style={{
          width: '30px',
          height: '30px',
          float: 'right',
          cursor: 'pointer',

          '&:hover': {
            color: 'var(--global-font-primary)',
          },
        }}
        onClick={() => {
          ToggleStore.setShowToggle(false)
        }}
      />
      {ToggleStore.toggleType === ToggleType.AddContact && (
        <>
          <h1 className='c-toggle-title'>添加</h1>
          <form className='c-toggle-form' onSubmit={handleSubmit(onSubmit)}>
            <button className='c-toggle-form-btn'>添加</button>
          </form>
        </>
      )}
      {ToggleStore.toggleType === ToggleType.CreateGroup && (
        <>
          <h1 className='c-toggle-title'>创建群聊</h1>
          <form className='c-toggle-form' onSubmit={handleSubmit(onSubmit)}>
            <input
              type='text'
              placeholder='群名称'
              className='c-toggle-form-input'
              {...register('name', { required: true })}
            />
            {errors.name?.type === 'required' && (
              <span className='c-toggle-form-hint'>群名称不能为空</span>
            )}
            <button className='c-toggle-form-btn'>创建</button>
          </form>
        </>
      )}
    </div>
  )
}

const Toggle = observer(_Toggle)

export default Toggle
