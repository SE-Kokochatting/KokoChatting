import './index.scss'
// type SwitchProps = {
// };
function Switch(/* props: SwitchProps */) {
  // const {} = props;
  return (
    <div className='switch_wrapper'>
      <label className='switch_wrapper-switch'>
        <input className='switch_wrapper-switch-input' type='checkbox' />
        <span className='switch_wrapper-switch-slider' />
      </label>
    </div>
  )
}
export default Switch
