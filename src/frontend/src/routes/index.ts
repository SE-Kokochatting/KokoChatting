import loadable from '@loadable/component'
import { RouteConfig } from 'react-router-config'

const routesConfig: RouteConfig[] = [
  {
    path: '/home',
    exact: true,
    element: loadable(() => import('@/pages/Home')),
  },
  {
    path: '/private',
    exact: true,
    element: loadable(() => import('@/pages/Private')),
  },
  {
    path: '/group',
    exact: true,
    element: loadable(() => import('@/pages/Group')),
  },
]

export default routesConfig
