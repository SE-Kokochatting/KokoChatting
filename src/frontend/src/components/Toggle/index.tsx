import { observer } from 'mobx-react-lite'
import { useForm } from 'react-hook-form'
import { useAlert } from 'react-alert'
import { useState } from 'react'
import { ToggleType, MessageType } from '@/enums'
import { ICreateGroup, createGroup } from '@/network/group/createGroup'
import { IAddFriend, addFriend } from '@/network/friend/addFriend'
import ToggleStore from '@/mobx/toggle'
import ChatListStore from '@/mobx/chatlist'
import SvgIcon from '../SvgIcon'
import './index.scss'

function _Toggle() {
  const {
    register,
    reset,
    handleSubmit,
    formState: { errors },
  } = useForm<Omit<IAddFriend, 'messageType'> | ICreateGroup>()

  enum AddType {
    Friend,
    Group,
  }

  const alert = useAlert()
  // 此时是添加好友还是群
  const [addType, setAddType] = useState<AddType>(AddType.Friend)

  async function onAddContactSubmit(reqData: any) {
    reqData.messageType = MessageType.FriendRequestNotify
    reqData.receiver = parseInt(reqData.receiver, 10)
    const { data } = await addFriend(reqData)
    if (!data) {
      alert.show('发送添加好友请求失败', {
        onClose: () => {
          reset()
        },
      })
      return
    }
    alert.show('已发送添加好友请求', {
      onClose: async () => {
        ToggleStore.setShowToggle(false)
        reset()
      },
    })
  }

  async function onCreateGroupSubmit(reqData: any) {
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
          <div className='c-toggle-tab'>
            <div
              className={
                addType === AddType.Friend
                  ? 'c-toggle-tab-item selected'
                  : 'c-toggle-tab-item'
              }
              onClick={() => {
                setAddType(AddType.Friend)
                reset()
              }}
            >
              添加好友
            </div>
            <div
              className={
                addType === AddType.Group
                  ? 'c-toggle-tab-item selected'
                  : 'c-toggle-tab-item'
              }
              onClick={() => {
                setAddType(AddType.Group)
                reset()
              }}
            >
              添加群
            </div>
          </div>
          <form
            className='c-toggle-form'
            onSubmit={handleSubmit(onAddContactSubmit)}
          >
            <input
              type='text'
              placeholder={addType === AddType.Friend ? 'uid' : 'gid'}
              className='c-toggle-form-input'
              {...register('receiver', { required: true, pattern: /^[0-9]+$/ })}
            />
            {errors.receiver?.type === 'required' && (
              <span className='entrance-window-form-hint'>uid 不能为空</span>
            )}
            <input
              type='text'
              placeholder='备注信息'
              className='c-toggle-form-input'
              {...register('messageContent', {
                required: true,
                pattern: /^[0-9]+$/,
              })}
            />
            {errors.messageContent?.type === 'required' && (
              <span className='entrance-window-form-hint'>
                备注信息不能为空
              </span>
            )}
            <button className='c-toggle-form-btn'>发送请求</button>
          </form>
        </>
      )}
      {ToggleStore.toggleType === ToggleType.CreateGroup && (
        <>
          <h1 className='c-toggle-title'>创建群聊</h1>
          <form
            className='c-toggle-form'
            onSubmit={handleSubmit(onCreateGroupSubmit)}
          >
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
