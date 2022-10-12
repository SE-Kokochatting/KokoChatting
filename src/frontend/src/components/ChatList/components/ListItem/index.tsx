import './index.scss'
// type ListItemProps = {
// };
function ListItem(/* props: ListItemProps */) {
  // const {} = props;
  return (
    <div className='c-chat_list-item'>
      <div className='c-chat_list-item-avatar'>
        <img
          className='c-chat_list-item-avatar-img'
          src='https://p.qqan.com/up/2021-2/16137992359659254.jpg'
        />
      </div>
      <div className='c-chat_list-item-main'>
        <span className='c-chat_list-item-main-name'>华小科</span>
        <span className='c-chat_list-item-main-content'>70周年校庆快乐！</span>
      </div>
      <div className='c-chat_list-item-time'>12:21</div>
    </div>
  )
}
export default ListItem
