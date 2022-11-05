import { RouterProvider } from 'react-router-dom'
import { useEffect } from 'react'
import { configure } from 'mobx'
import { positions, Provider } from 'react-alert'
import { router } from '@/routes'
import { WsHost } from './consts'
import { getToken } from './utils/token'
import AlertMUITemplate from 'react-alert-template-mui'
import WS from '@/ws'

import './App.scss'

configure({
  enforceActions: 'never',
})

const options = {
  timeout: 3000,
  position: positions.BOTTOM_CENTER,
  transition: 'fade',
}

function App() {
  useEffect(() => {
    const token = getToken()
    if (token) {
      const socket = new WS(`${WsHost}/upgrade_protocol`, ['chat', token])
      socket.init(
        {
          time: 60 * 1000,
          timeout: 60 * 1000,
          reconnect: 60 * 1000,
        },
        true,
      )
    }
  }, [])

  return (
    <Provider template={AlertMUITemplate} {...options}>
      <RouterProvider router={router} />
    </Provider>
  )
}

export default App
