import { observer } from 'mobx-react-lite'
import { ChatType } from '@/enums'
import { getMsgId } from '@/utils/message'
import { getUid } from '@/utils/uid'
import MsgStore from '@/mobx/msg'
import ChatStore from '@/mobx/chat'
import Bubble from './components/Bubble'
import Sender from './components/Sender'
import './index.scss'

function _ChatWindow() {
  const mid = getMsgId()
  const uid = getUid()

  return (
    <div className='c-chat_window'>
      <div className='c-chat_window-chat_area'>
        {ChatStore.currentChat?.uid > 0 &&
          MsgStore.friendMsg.map(
            ({ sendTime, readUids, messageContent, messageId, senderId }) =>
              (senderId === ChatStore.currentChat?.uid || senderId === uid) && (
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
              lastMessageTime,
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
                  lastMessageTime={lastMessageTime}
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
