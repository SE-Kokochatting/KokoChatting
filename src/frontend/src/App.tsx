import { Route, Routes, Navigate } from 'react-router-dom'
import routes from '@/routes'
import './App.scss'

function App() {
  return (
    <Routes>
      <Route path='/' element={<Navigate to='/home' replace />} />
      {routes.map((item, i) => {
        return (
          <Route
            key={i}
            path={item.path as string}
            element={<item.element />}
          />
        )
      })}
    </Routes>
  )
}

export default App
