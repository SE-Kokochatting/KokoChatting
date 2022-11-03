import { observer } from 'mobx-react-lite'
import { ChatType, MessageType } from '@/enums'
import ChatStore from '@/mobx/chat'
import Bubble from './components/Bubble'
import Sender from './components/Sender'
import './index.scss'

function _ChatWindow() {
  return (
    <div className='c-chat_window'>
      <div className='c-chat_window-chat_area'>
        {ChatStore.bubblesData.map(
          ({
            lastMessageTime,
            readUids,
            messageContent,
            messageId,
            senderId,
            messageType,
          }) => (
            <Bubble
              key={messageId}
              lastMessageTime={lastMessageTime}
              readUids={readUids}
              messageContent={messageContent}
              senderId={senderId}
              chatType={
                messageType === MessageType.SingleMessage
                  ? ChatType.Private
                  : ChatType.Group
              }
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
