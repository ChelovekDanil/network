import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import './auth.css'

function Registration() {
    const navigate = useNavigate()    

    function authLoginAuth() {
        navigate("/login")
    }
    
    function authCheck() {
        const rlogin = /^[a-zA-Z]{3,}$/
        if (!rlogin.test(login)) {
            alert("Логин должен состоять из латинских символов и длина больше 2 символов")
            return
        }

        if (password !== repassword) {
            alert("Пароли не совподают")
            return
        }

        if (password.length <= 6) {
            alert("Пароль должен быть больше 6 символов")
            return
        }

        localStorage.setItem("isAuth", "true")
        navigate("/")
    }

    const [login, setLogin] = useState("")
    const [password, setPassword] = useState("")
    const [repassword, setRepassword] = useState("")
    
    const ChangeLogin = (e) => {
        setLogin(e.target.value)
    } 
    
    const ChangePassword = (e) => {
        setPassword(e.target.value)
    }
    
    const ChangeRepassword = (e) => {
        setRepassword(e.target.value)
    }

    return(
        <div className="auth-block-fraper">
            <div className="auth-block">
            <p id="auth-block-header">Регистрация</p>
                <p>Логин</p>
                <input type="text" id="login_auth" value={login} onChange={ChangeLogin}/>
                <p>Пароль</p>
                <input type="password" id="password_auth" value={password} onChange={ChangePassword}/>
                <p>Повторите пароль</p>
                <input type="password" id="repassword_auth" value={repassword} onChange={ChangeRepassword}/>
                <div className="auth-block-nav">
                    <p id="auth-block-nav-login" onClick={authLoginAuth}>Войти</p>
                </div>
                <button onClick={authCheck}>войти</button>
            </div>
        </div>
    )
}

export default Registration