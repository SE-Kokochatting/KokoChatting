import { useNavigate, useLocation } from 'react-router-dom'
import { Theme } from '@/enums'
import ThemeStore from '@/mobx/theme'
import './index.scss'
// type LoginProps = {
// };
function Login(/* props: LoginProps */) {
  // const {} = props;
  const navigate = useNavigate()
  const location = useLocation()
  const { pathname } = location
  return (
    <div
      className={ThemeStore.theme === Theme.Dark ? 'entrance dark' : 'entrance'}
    >
      <div className='entrance-window'>
        <div className='entrance-window-tab'>
          <div
            className={
              pathname === '/login'
                ? 'entrance-window-tab-item selected'
                : 'entrance-window-tab-item'
            }
            onClick={() => {
              navigate('/login')
            }}
          >
            登录
          </div>
          <div
            className={
              pathname === '/register'
                ? 'entrance-window-tab-item selected'
                : 'entrance-window-tab-item'
            }
            onClick={() => {
              navigate('/register')
            }}
          >
            注册
          </div>
        </div>
        <div className='entrance-window-form'>
          {pathname === '/login' ? (
            <>
              <input
                type='text'
                placeholder='uid'
                className='entrance-window-form-input'
              />
              <input
                type='password'
                placeholder='密码'
                className='entrance-window-form-input'
              />
              <button className='entrance-window-form-btn'>登录</button>
            </>
          ) : (
            <>
              <input
                type='text'
                placeholder='用户名'
                className='entrance-window-form-input'
              />
              <input
                type='password'
                placeholder='密码'
                className='entrance-window-form-input'
              />
              <button className='entrance-window-form-btn'>注册</button>
            </>
          )}
        </div>
      </div>
    </div>
  )
}
export default Login
