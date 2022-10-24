import SvgIcon from '../SvgIcon'
import './index.scss'
// type UserInfoProps = {
// };
function UserInfo(/* props: UserInfoProps */) {
  // const {} = props;
  return (
    <div className='user_info'>
      <h1 className='user_info-title'>
        <SvgIcon
          name='cross'
          style={{
            fill: '#fff',
            width: '20px',
            height: '20px',
            marginRight: '10px',
            cursor: 'pointer',

            '&:hover': {},
          }}
        />
        用户信息
      </h1>
      <img
        className='user_info-avatar'
        src='https://p.qqan.com/up/2021-2/16137992359659254.jpg'
      />
    </div>
  )
}
export default UserInfo
