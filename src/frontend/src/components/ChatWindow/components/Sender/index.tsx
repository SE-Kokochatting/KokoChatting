import { useForm } from 'react-hook-form'
import { useAlert } from 'react-alert'
import { sendMsg } from '@/network/message/sendMsg'
import { MessageType } from '@/enums'
import Emitter from '@/utils/eventEmitter'
import UserStore from '@/mobx/user'
import ChatStore from '@/mobx/chat'
import MsgStore from '@/mobx/msg'
import SvgIcon from '@/components/SvgIcon'
import './index.scss'

interface ISendMsg {
  content: string
}

function Sender() {
  const alert = useAlert()
  const { register, reset, handleSubmit } = useForm<ISendMsg>()

  async function onSubmit({ content }: ISendMsg) {
    reset()
    if (!ChatStore.currentChat) {
      alert.show('请先选择要发送消息的对象', {
        title: '消息发送失败',
      })
    } else {
      const messageType = ChatStore.currentChat.uid
        ? MessageType.SingleMessage
        : MessageType.GroupMessage
      const { code, data } = await sendMsg({
        receiver: ChatStore.currentChat.uid as number,
        messageContent: content,
        messageType,
      })
      if (code === 3000) {
        alert.show('发生了未知的错误', {
          title: '消息发送失败',
        })
        return
      }
      const { msgid } = data
      const now = new Date()
      MsgStore.sendMsg(
        {
          messageId: msgid,
          messageContent: content,
          messageType: messageType,
          name: UserStore.name,
          senderId: UserStore.uid,
          receiverId: ChatStore.currentChat?.uid,
          groupId: ChatStore.currentChat?.gid,
          readUids: [],
          avatarUrl: UserStore.avatarUrl,
          sendTime: `${now.getHours()}:${now.getMinutes()}`,
        },
        messageType,
      )
    }
  }

  return (
    <div className='c-chat_window-sender'>
      <SvgIcon
        name='link'
        style={{
          width: '30px',
          height: '30px',
          color: 'var(--global-font-primary_lighter)',
          margin: '0 15px',
          cursor: 'pointer',
        }}
      />
      <form
        className='c-chat_window-sender-form'
        onSubmit={handleSubmit(onSubmit)}
      >
        <input
          className='c-chat_window-sender-form-input'
          type='text'
          placeholder='Write a message...'
          autoComplete='off'
          {...register('content', { required: true })}
        />
      </form>

      <div style={{ marginRight: '20px' }}>
        {/* <SvgIcon
          name='expression'
          style={{
            width: '30px',
            height: '30px',
            color: 'var(--global-font-primary_lighter)',
            marginRight: '20px',
            cursor: 'pointer',
          }}
        /> */}
        <SvgIcon
          name='send'
          style={{
            width: '30px',
            height: '30px',
            color: 'var(--global-font-primary_lighter)',
            cursor: 'pointer',
          }}
        />
      </div>
    </div>
  )
}
export default Sender
