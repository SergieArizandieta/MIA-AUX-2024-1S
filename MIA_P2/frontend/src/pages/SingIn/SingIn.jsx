import { Link, useParams } from "react-router-dom";


export default function SingIn() {
   const { disk, part } = useParams()

   
   const handleSubmit = (e) => {
      e.preventDefault()
      console.log("submit", disk, part)

      const user = e.target.uname.value
      const pass = e.target.psw.value

      console.log("user", user, pass)
   }

  return (
    <>
      <p>Hola mundo Login</p>
      <br/>
      <Link to="/DiskCreen">Commands</Link>
      <br />
      <br />
      <br />
      <br />

      <form onSubmit={handleSubmit}>
         

         <div className="container">
            <label htmlFor="uname"><b>Username</b></label>
            <input type="text" placeholder="Enter Username" name="uname" required/>

            <label htmlFor="psw"><b>Password</b></label>
            <input type="password" placeholder="Enter Password" name="psw" required/>
            
            <button type="submit">Login</button>
           
         </div>
        
      </form>


   </>
  )
}