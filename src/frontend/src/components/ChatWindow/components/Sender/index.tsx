import { useForm } from 'react-hook-form'
import { useAlert } from 'react-alert'
import { sendMsg } from '@/network/message/sendMsg'
import { MessageType } from '@/enums'
import ChatStore from '@/mobx/chat'
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
      const { code } = await sendMsg({
        receiver: ChatStore.currentChat.uid as number,
        messageContent: content,
        messageType: ChatStore.currentChat.uid
          ? MessageType.SingleMessage
          : MessageType.GroupMessage,
      })
      if (code === 3000) {
        alert.show('发生了未知的错误', {
          title: '消息发送失败',
        })
      }
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
