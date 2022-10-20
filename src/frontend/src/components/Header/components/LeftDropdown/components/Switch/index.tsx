import { observer } from 'mobx-react-lite'
import { Theme } from '@/enums'
import ThemeStore from '@/mobx/theme'
import './index.scss'

// type SwitchProps = {
// };
function _Switch(/* props: SwitchProps */) {
  // const {} = props;
  return (
    <div className='switch_wrapper'>
      <label className='switch_wrapper-switch'>
        <input
          className='switch_wrapper-switch-input'
          checked={ThemeStore.theme === Theme.Dark ? true : false}
          type='checkbox'
          onChange={(e) => {
            const value = e.target.checked
            ThemeStore.setTheme(value ? Theme.Dark : Theme.Light)
          }}
        />
        <span className='switch_wrapper-switch-slider' />
      </label>
    </div>
  )
}

const Switch = observer(_Switch)

export default Switch
