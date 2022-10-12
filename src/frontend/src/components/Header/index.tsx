import SvgIcon from '@/components/SvgIcon'
import Search from './components/Search'
import './index.scss'
//  type HeaderProps = {
//  };
function Header(/* props: HeaderProps */) {
  // const {} = props;
  return (
    <div className='c-header'>
      <div className='c-header-left'>
        <SvgIcon
          name='menu'
          style={{
            fill: '#fff',
            width: '35px',
            height: '35px',
            marginLeft: '10px',
            cursor: 'pointer',
          }}
        />
        <Search />
      </div>
      <div className='c-header-right'>
        <div className='c-header-user'>
          <span className='c-header-user-name'>华小科</span>
          <span className='c-header-user-state'>online</span>
        </div>
        <SvgIcon
          name='search'
          style={{
            fill: '#fff',
            width: '35px',
            height: '35px',
            position: 'absolute',
            right: '0.7rem',
            cursor: 'pointer',
          }}
        />
        <SvgIcon
          name='right-menu'
          style={{
            fill: '#fff',
            width: '35px',
            height: '35px',
            position: 'absolute',
            right: '0.2rem',
            cursor: 'pointer',
          }}
        />
      </div>
    </div>
  )
}
export default Header
