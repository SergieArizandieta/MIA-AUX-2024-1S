import { Routes, Route, HashRouter } from 'react-router-dom'

import Home from '../pages/Home/Home'
import Commands from '../pages/Commands/Commands'

export default function AppNavigator() {
  return (
    <HashRouter>
      <Routes>
 
          <Route path="/" element={<Home/>} />
          <Route path="/Commands" element={<Commands/>} />

      </Routes>
    </HashRouter>
  )
}