import SvgIcon from '@/components/SvgIcon/index'
import './index.scss'

function Home() {
  return (
    <div className='home'>
      Hello world
      <SvgIcon name='vite' style={{ width: '20px', height: '20px', marginLeft: '10px' }} />
    </div>
  )
}

export default Home
