import { Routes, Route, HashRouter } from 'react-router-dom'

import { useState } from "react";

import Home from '../pages/Home/Home'
import Commands from '../pages/DiskCreen/DiskCreen'
import Partition from '../pages/Partition/Partition'
import SingIn from '../pages/SingIn/SingIn'

export default function AppNavigator() {
  const [ip, setIP] = useState("localhost") 

  const handleChage = (e) => {
    console.log(e.target.value)
    setIP(e.target.value)
  }

  return (
    <HashRouter>
      IP: <input type="text" onChange={handleChage}/> -- {ip}
      <Routes>
         
          <Route path="/" element={<Home/>} />
          <Route path="/DiskCreen" element={<Commands ip={ip}/>} />
          <Route path="/disk/:id/" element={<Partition ip={ip}/>} />
          <Route path="/login/:disk/:part" element={<SingIn ip={ip}/>} />

      </Routes>
    </HashRouter>
  )
}