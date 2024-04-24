import diskIMG from "../../assets/disk.png";
import { useState } from "react";
import { Link, useNavigate } from "react-router-dom";

export default function DiskCreen() {
  const [data, setData] = useState([]) 
  const navigate = useNavigate()
  
  // execute the fetch command only once and when the component is loaded
  useState(() => {
 
    // const rawData = {
    //   "rutas":["A.dsk","B.dsk","C.dsk","D.dsk","A.dsk","B.dsk","C.dsk","D.dsk","A.dsk","B.dsk","C.dsk","D.dsk"]
    // }
    // setData(rawData.rutas)

    var dataF = {
      User: 'root',
      Password: 'admin'
    }
    
    fetch('http://localhost:4000/tasks', {
      method: 'POST', 
      headers: {
        'Content-Type': 'application/json' 
      },
      body: JSON.stringify(dataF)
    })
    .then(response => response.json())
    .then(data => {
      console.log(data); // Do something with the response
      setData(data.List)
    })
    .catch(error => {
      console.error('There was an error with the fetch operation:', error);
    });
  }, [])

  const onClick = (objIterable) => {
    //e.preventDefault()
    console.log("click",objIterable)
    navigate(`/disk/${objIterable}`)
  }

  return (
    <>
      <p>Hola Commands</p>
      <br/>
      <Link to="/">Home</Link>

      <br/>
      <br/>
      <br/>
      <br/>

      <div style={{border:"red 1px solid",display: "flex", flexDirection: "row"}}>

        {
          data.map((objIterable, index) => {
            return (
              <div key={index} style={{
                border: "green 1px solid",
                display: "flex",
                flexDirection: "column", // Alinea los elementos en columnas
                alignItems: "center", // Centra verticalmente los elementos
                maxWidth: "100px",
              }}
              onClick={() => onClick(objIterable)}
              >
                <img src={diskIMG} alt="disk" style={{width: "100px"}} />
                <p>{objIterable}</p>
              </div>
            )
          })
        }
      
      </div>
    </>
   )
 }
 