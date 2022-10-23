import { RouteConfig } from 'react-router-config'
import Chat from '@/pages/Chat'
import Login from '@/pages/Login'

const routesConfig: RouteConfig[] = [
  {
    path: '/home',
    exact: true,
    element: Chat,
  },
  {
    path: '/private',
    exact: true,
    element: Chat,
  },
  {
    path: '/group',
    exact: true,
    element: Chat,
  },
  {
    path: '/login',
    exact: true,
    element: Login,
  },
  {
    path: '/register',
    exact: true,
    element: Login,
  },
]

export default routesConfig
