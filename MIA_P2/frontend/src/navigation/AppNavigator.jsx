import { Routes, Route, HashRouter } from 'react-router-dom'

import Home from '../pages/Home/Home'
import Commands from '../pages/DiskCreen/DiskCreen'
import Partition from '../pages/Partition/Partition'
import SingIn from '../pages/SingIn/SingIn'

export default function AppNavigator() {
  return (
    <HashRouter>
      <Routes>
 
          <Route path="/" element={<Home/>} />
          <Route path="/DiskCreen" element={<Commands/>} />
          <Route path="/disk/:id/" element={<Partition/>} />
          <Route path="/login/:disk/:part" element={<SingIn/>} />

      </Routes>
    </HashRouter>
  )
}