import { useState } from 'react'
import './auth.css'
import { useNavigate } from 'react-router-dom'

function Login() {
    const navigate = useNavigate()
    const [login, setLogin] = useState("")
    const [password, setPassword] = useState("")

    const ChangeLogin = (e) => {
        setLogin(e.target.value)
    }

    const ChangePassword = (e) => {
        setPassword(e.target.value)
    }
    
    function authRegistrationAuth() {
        navigate("/registration")
    }
    
    const authCheck = async () => {
        const rlogin = /^[a-zA-Z]{3,}$/
        if (!rlogin.test(login)) {
            alert("Логин должен состоять из латинских символов и длина больше 2 символов")
            return
        }

        if (password.length <= 6) {
            alert("Пароль должен быть больше 6 символов")
            return
        }

        try {
            const response = await fetch("http://localhost:8080/auth/login/", {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({login: login, passhash: password}),
            });

            if (!response.ok) {
                throw new Error('Network response was not ok')
            }

            const result = await response.json()

            localStorage.setItem("accessToken", result.accessToken)
            localStorage.setItem("refreshToken", result.refreshToken)
            localStorage.setItem("login", login)
            localStorage.setItem("isAuth", "true")
            navigate("/main")
        } catch (error) {
            console.log('Error: ', error)
        }
    }

    return(
        <div className="auth-block-fraper">
            <div className="auth-block">
            <p id="auth-block-header">Вход</p>
                <p>Логин</p>
                <input type="text" id="login_auth" value={login} onChange={ChangeLogin}/>
                <p>Пароль</p>
                <input type="password" id="password_auth" value={password} onChange={ChangePassword}/>
                <div className="auth-block-nav">
                    <p id="auth-block-nav-registration" onClick={authRegistrationAuth}>Зарегистрироваться</p>
                </div>
                <button onClick={authCheck}>войти</button>
            </div>
        </div>
    )
}

export default Login