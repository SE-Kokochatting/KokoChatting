import { RouteConfig } from 'react-router-config'
import Chat from '@/pages/Chat'

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
]

export default routesConfig
