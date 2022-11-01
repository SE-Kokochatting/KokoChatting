import { useNavigate, useLocation } from 'react-router-dom'
import { useForm } from 'react-hook-form'
import { useAlert } from 'react-alert'
import { sha256 } from 'js-sha256'
import { Theme } from '@/enums'
import { IRegister, register as postRegisterReq } from '@/network/user/register'
import { ILogin, login as postLoginReq } from '@/network/user/login'
import ThemeStore from '@/mobx/theme'
import { setToken } from '@/utils/token'
import { setUid } from '@/utils/uid'
import './index.scss'

// Todo: 如果 token 有效，重定向到 /home
function Login() {
  const {
    register,
    reset,
    handleSubmit,
    formState: { errors },
  } = useForm<ILogin & IRegister>()
  const navigate = useNavigate()
  const location = useLocation()
  const { pathname } = location
  const alert = useAlert()

  const onSubmit = async (reqData: any) => {
    if (pathname === '/login') {
      // sha256 加密
      reqData.password = sha256(reqData.password)
      reqData.uid = parseInt(reqData.uid, 10)
      const { code, data } = await postLoginReq(reqData)
      // 登录失败
      if (!data) {
        switch (code) {
          case 1002:
            alert.show('密码错误！', {
              title: '登录失败',
            })
            break
        }
        return
      }
      const { token } = data
      // 设置 token
      setToken(token)
      // 设置 uid
      setUid(reqData.uid)
      alert.show('登录成功', {
        onClose: () => {
          navigate('/home')
        },
      })
      navigate('/home')
    } else {
      const { code, data } = await postRegisterReq(reqData)
      // 注册失败
      if (!data) {
        switch (code) {
          case 1001:
            alert.show('该用户已注册！', {
              title: '注册失败',
            })
            break
        }
        return
      }
      const { uid } = data
      alert.show(`您的 uid 为 ${uid}，请及时保存！`, {
        timeout: 60000,
        title: '注册成功',
        onClose: () => {
          reset()
          navigate('/login')
        },
      })
    }
  }

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
              reset()
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
              reset()
            }}
          >
            注册
          </div>
        </div>
        <form
          className='entrance-window-form'
          onSubmit={handleSubmit(onSubmit)}
        >
          {pathname === '/login' ? (
            <>
              <input
                type='text'
                placeholder='uid'
                className='entrance-window-form-input'
                {...register('uid', { required: true, pattern: /^[0-9]+$/ })}
              />
              {errors.uid?.type === 'required' && (
                <span className='entrance-window-form-hint'>uid 不能为空</span>
              )}
              {errors.uid?.type === 'pattern' && (
                <span className='entrance-window-form-hint'>
                  uid 必须为数字
                </span>
              )}
              <input
                type='password'
                placeholder='密码'
                className='entrance-window-form-input'
                {...register('password', { required: true })}
              />
              {errors.password?.type === 'required' && (
                <span className='entrance-window-form-hint'>密码不能为空</span>
              )}
              <button className='entrance-window-form-btn'>登录</button>
            </>
          ) : (
            <>
              <input
                type='text'
                placeholder='用户名'
                className='entrance-window-form-input'
                {...register('name', { required: true })}
              />
              {errors.name?.type === 'required' && (
                <span className='entrance-window-form-hint'>
                  用户名不能为空
                </span>
              )}
              <input
                type='password'
                placeholder='密码'
                className='entrance-window-form-input'
                {...register('password', { required: true })}
              />
              {errors.password?.type === 'required' && (
                <span className='entrance-window-form-hint'>密码不能为空</span>
              )}
              <button className='entrance-window-form-btn'>注册</button>
            </>
          )}
        </form>
      </div>
    </div>
  )
}
export default Login
