import { observer } from 'mobx-react-lite'
import { useForm } from 'react-hook-form'
import { useAlert } from 'react-alert'
import { useState, useEffect } from 'react'
import { ToggleType, MessageType, AddType, NotifyType } from '@/enums'
import { IMessageContent } from '@/types'
import { ICreateGroup, createGroup } from '@/network/group/createGroup'
import { IAddFriend, addFriend } from '@/network/friend/addFriend'
import ToggleStore from '@/mobx/toggle'
import ChatListStore from '@/mobx/chatlist'
import MsgStore from '@/mobx/msg'
import NotifyItem from './components/NotifyItem'
import SvgIcon from '../SvgIcon'
import './index.scss'

function _Toggle() {
  const {
    register,
    reset,
    handleSubmit,
    formState: { errors },
  } = useForm<Partial<IMessageContent & ICreateGroup & IAddFriend>>()

  const alert = useAlert()
  // 添加好友/群
  const [addType, setAddType] = useState<AddType>(AddType.Friend)
  // 好友请求/群通知
  const [notifyType, setNotifyType] = useState<NotifyType>(
    NotifyType.friendRequest,
  )

  useEffect(() => {
    MsgStore.pullMsgContent(
      notifyType === NotifyType.friendRequest
        ? MessageType.FriendRequestNotify
        : MessageType.JoinGroupRequestNotify,
    )
  }, [ToggleStore.showToggle, ToggleStore.toggleType])

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
      {ToggleStore.toggleType === ToggleType.Notify && (
        <>
          <div className='c-toggle-tab'>
            <div
              className={
                notifyType === NotifyType.friendRequest
                  ? 'c-toggle-tab-item selected'
                  : 'c-toggle-tab-item'
              }
              onClick={() => {
                setNotifyType(NotifyType.friendRequest)
              }}
            >
              好友请求
            </div>
            <div
              className={
                notifyType === NotifyType.groupManageNotify
                  ? 'c-toggle-tab-item selected'
                  : 'c-toggle-tab-item'
              }
              onClick={() => {
                setNotifyType(NotifyType.groupManageNotify)
              }}
            >
              群通知
            </div>
          </div>
          {notifyType === NotifyType.friendRequest &&
            MsgStore.friendRequest.map((msg: IMessageContent) => (
              <NotifyItem
                publisherName={`${msg.senderId}`}
                info={msg.messageContent}
                mid={msg.messageId}
                type={MessageType.FriendRequestNotify}
                key={msg.messageId}
              />
            ))}
          {notifyType === NotifyType.groupManageNotify &&
            MsgStore.groupNotify.map((msg: IMessageContent) => (
              <NotifyItem
                publisherName={`${msg.groupId}`}
                info={msg.messageContent}
                mid={msg.messageId}
                type={MessageType.JoinGroupRequestNotify}
                key={msg.messageId}
              />
            ))}
        </>
      )}
    </div>
  )
}

const Toggle = observer(_Toggle)

export default Toggle
