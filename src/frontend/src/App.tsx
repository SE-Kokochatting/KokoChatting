import { RouterProvider } from 'react-router-dom'
import { useEffect } from 'react'
import { configure } from 'mobx'
import { positions, Provider } from 'react-alert'
import { router } from '@/routes'
import { WsHost } from '@/consts'
import { getToken } from '@/utils/token'
import Emitter from '@/utils/eventEmitter'
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

let socket: WS | null = null

function connectWebSocket() {
  const token = getToken()
  if (token) {
    socket = new WS(`${WsHost}/upgrade_protocol`, ['chat', token])
    socket.init(
      {
        time: 20 * 1000,
        timeout: 10 * 1000,
        reconnect: 5 * 1000,
      },
      true,
    )
  }
}

function reconnectWebSocket() {
  // 入口函数
  if (socket !== null) {
    socket.clear()
    socket = null
  }
  connectWebSocket()
}

function scrollToBottom() {
  const ChatWindowDom = document.querySelector(
    '.c-chat_window-chat_area',
  ) as Element
  ChatWindowDom.scrollTo(0, ChatWindowDom.scrollHeight)
}

function App() {
  useEffect(() => {
    if (socket !== null) {
      return
    }
    connectWebSocket()
  }, [])

  useEffect(() => {
    Emitter.on('reconnect', reconnectWebSocket)
    Emitter.on('scrollToBottom', scrollToBottom)
    return () => {
      Emitter.removeListener('reconnect')
      Emitter.removeListener('scrollToBottom')
    }
  }, [])

  return (
    <Provider template={AlertMUITemplate} {...options}>
      <RouterProvider router={router} />
    </Provider>
  )
}

export default App
