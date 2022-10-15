import ListItem from './components/ListItem'
import './index.scss'
// type ChatListProps = {
// };

// Todo: 类型定义
const chatListData = [
  {
    id: 1,
    avatarUrl: 'https://p.qqan.com/up/2021-2/16137992359659254.jpg',
    name: '华小科',
    extract: '70周年校庆快乐！',
    lastTime: '8:01',
  },
  {
    id: 2,
    avatarUrl: 'https://p.qqan.com/up/2021-2/16137992359659254.jpg',
    name: '华大科',
    extract: '70周年校庆快乐！',
    lastTime: '8:01',
  },
  {
    id: 3,
    avatarUrl: 'https://p.qqan.com/up/2021-2/16137992359659254.jpg',
    name: '芝士软工',
    extract: '70周年校庆快乐！',
    lastTime: '8:01',
  },
]

function ChatList(/* props: ChatListProps */) {
  // const {} = props;
  return (
    <div className='c-chat_list'>
      {chatListData.map(({ id, avatarUrl, name, extract, lastTime }) => (
        <ListItem
          key={id}
          avatarUrl={avatarUrl}
          name={name}
          extract={extract}
          lastTime={lastTime}
        />
      ))}
    </div>
  )
}
export default ChatList
