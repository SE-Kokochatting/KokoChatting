import { observer } from 'mobx-react-lite'
import { ChatType } from '@/enums'
import { getMsgId } from '@/utils/message'
import { getUid } from '@/utils/uid'
import { useEffect } from 'react'
import { msgHasRead } from '@/network/message/msgHasRead'
import MsgStore from '@/mobx/msg'
import ChatStore from '@/mobx/chat'
import Emitter from '@/utils/eventEmitter'
import Bubble from './components/Bubble'
import Sender from './components/Sender'
import './index.scss'

function _ChatWindow() {
  const mid = getMsgId()
  const uid = getUid()

  function handleMsgRead() {
    // 在页面上渲染的消息
    const renderedMsg =
      ChatStore.currentChat?.uid > 0
        ? MsgStore.friendMsg.filter(
            ({ senderId, receiverId }) =>
              senderId === ChatStore.currentChat?.uid ||
              (senderId === uid && ChatStore.currentChat?.uid === receiverId),
          )
        : MsgStore.groupMsg.filter(
            ({ groupId }) => groupId === ChatStore.currentChat?.gid,
          )
    // 气泡 Dom 数组
    const bubbleDomArr = [
      ...document.querySelectorAll('.c-chat_window-chat_area-bubble_wrapper'),
    ]
    // key: DOM 值: message
    const bubbleMap = {}

    bubbleDomArr.forEach((item, i) => {
      bubbleMap[item.id] = renderedMsg[i]
    })

    const io = new IntersectionObserver((entries) => {
      entries.forEach((item) => {
        if (item.isIntersecting) {
          const msgObj = bubbleMap[item.target.id]
          const { senderId, messageId, readUids } = msgObj
          if (senderId !== uid && !readUids.includes(uid)) {
            msgHasRead({ msgids: [messageId] }).then(() => {
              io.unobserve(item.target)
            })
          }
        }
      })
    })

    bubbleDomArr.forEach((item) => io.observe(item))
  }

  useEffect(() => {
    handleMsgRead()
    Emitter.emit('scrollToBottom')
  }, [ChatStore.currentChat])

  useEffect(() => {
    Emitter.on('updateIntersect', handleMsgRead)
    return () => {
      Emitter.removeListener('updateIntersect')
    }
  }, [])

  return (
    <div className='c-chat_window'>
      <div className='c-chat_window-chat_area'>
        {ChatStore.currentChat?.uid > 0 &&
          MsgStore.friendMsg.map(
            ({
              sendTime,
              readUids,
              messageContent,
              messageId,
              senderId,
              receiverId,
            }) =>
              (senderId === ChatStore.currentChat?.uid ||
                (senderId === uid &&
                  ChatStore.currentChat?.uid === receiverId)) && (
                // (messageId as number) > mid &&
                <Bubble
                  key={messageId}
                  sendTime={sendTime}
                  readUids={readUids}
                  messageContent={messageContent}
                  senderId={senderId}
                  chatType={ChatType.Private}
                />
              ),
          )}
        {ChatStore.currentChat?.gid > 0 &&
          MsgStore.groupMsg.map(
            ({
              sendTime,
              readUids,
              messageContent,
              messageId,
              groupId,
              name,
              avatarUrl,
            }) =>
              groupId === ChatStore.currentChat?.gid &&
              (messageId as number) > mid && (
                <Bubble
                  key={messageId}
                  sendTime={sendTime}
                  readUids={readUids}
                  messageContent={messageContent}
                  groupId={groupId}
                  chatType={ChatType.Group}
                  name={name}
                  avatarUrl={avatarUrl}
                />
              ),
          )}
      </div>
      <Sender />
    </div>
  )
}

const ChatWindow = observer(_ChatWindow)

export default ChatWindow
