import {useEffect, useState} from "react";
import axios from "axios";

const login = () => (
    window.location.href = "http://localhost:8080/login"
)

const logout = () => (
    window.location.href = "http://localhost:8080/logout"
)

const User = () => {
  const [isLogin, setIsLogin] = useState(false)
  const [user, setUser] = useState("")

  useEffect(() => {
    // ログイン状況を取得
    axios.get("http://localhost:8080/me",
        {
          withCredentials: true
        })
    .then((res) => {
      setUser(res.data.name)
      setIsLogin(user !== "")
      console.log(res.data)
    })
    .catch((error) => {
      setUser("")
      setIsLogin(false)
    })
  })

  let status = "Not logged in"
  let button = <button onClick={login}>Login</button>
  if (isLogin) {
    status = "Logged in"
    button = <button onClick={logout}>Logout</button>
  }

  return (
      <div className="User">
        Welcome to auth0 sample page!<br/>
        Your status: {status}<br/>
        {(isLogin) && "Your name: " + user}<br/>
        {button}
      </div>
  )
}

export default User;
