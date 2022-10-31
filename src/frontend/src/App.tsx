import { RouterProvider } from 'react-router-dom'
import { positions, Provider } from 'react-alert'
import AlertMUITemplate from 'react-alert-template-mui'
import { router } from '@/routes'
import './App.scss'

const options = {
  timeout: 3000,
  position: positions.BOTTOM_CENTER,
  transition: 'fade',
}

function App() {
  return (
    <Provider template={AlertMUITemplate} {...options}>
      <RouterProvider router={router} />
    </Provider>
  )
}

export default App
