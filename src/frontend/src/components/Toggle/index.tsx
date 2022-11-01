import { observer } from 'mobx-react-lite'
import { useForm } from 'react-hook-form'
import { useAlert } from 'react-alert'
import { ToggleType } from '@/enums'
import { ICreateGroup, createGroup } from '@/network/group/createGroup'
import ToggleStore from '@/mobx/toggle'
import ChatListStore from '@/mobx/chatList'
import SvgIcon from '../SvgIcon'
import './index.scss'

function _Toggle() {
  const {
    register,
    reset,
    handleSubmit,
    formState: { errors },
  } = useForm<ICreateGroup>()

  const alert = useAlert()
  const onSubmit = async (reqData: any) => {
    const { code, data } = await createGroup(reqData)
    if (!data) {
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
    }
    alert.show('创建群聊成功', {
      onClose: async () => {
        // 更新群列表
        await ChatListStore.updateGroup()
        ToggleStore.setShowToggle(false)
        reset()
      },
    })
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
