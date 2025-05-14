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

    function authRepasswordAuth() {
        navigate("/repassword")
    }
    
    function authRegistrationAuth() {
        navigate("/registration")
    }
    
    function authCheck() {
        const rlogin = /^[a-zA-Z]{3,}$/
        if (!rlogin.test(login)) {
            alert("Логин должен состоять из латинских символов и длина больше 2 символов")
            return
        }

        if (password.length <= 6) {
            alert("Пароль должен быть больше 6 символов")
            return
        }

        localStorage.setItem("isAuth", "true")
        navigate("/")
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
                    <p id="auth-block-nav-repassword" onClick={authRepasswordAuth}>Забыл пароль</p>
                    <p id="auth-block-nav-registration" onClick={authRegistrationAuth}>Зарегистрироваться</p>
                </div>
                <button onClick={authCheck}>войти</button>
            </div>
        </div>
    )
}

export default Login