import SvgIcon from '@/components/SvgIcon'
import './index.scss'
// type SenderProps = {
// };
function Sender(/* props: SenderProps */) {
  // const {} = props;
  return (
    <div className='c-chat_window-sender'>
      <SvgIcon
        name='link'
        style={{
          width: '30px',
          height: '30px',
          color: 'var(--light)',
          margin: '0 15px',
          cursor: 'pointer',
        }}
      />
      <input
        className='c-chat_window-sender-input'
        type='text'
        placeholder='Write a message...'
      />
      <div style={{ marginRight: '20px' }}>
        <SvgIcon
          name='expression'
          style={{
            width: '30px',
            height: '30px',
            color: 'var(--light)',
            marginRight: '20px',
            cursor: 'pointer',
          }}
        />
        <SvgIcon
          name='send'
          style={{
            width: '30px',
            height: '30px',
            color: 'var(--light)',
            cursor: 'pointer',
          }}
        />
      </div>
    </div>
  )
}
export default Sender
