import './auth.css'

function authRepasswordAuth() {
    console.log("repassword")
}

function authRegistrationAuth() {
    console.log("registration")
}

function authCheck() {
    console.log("auth")
}

function Login() {
    return(
        <div className="auth-block-fraper">
            <div className="auth-block">
            <p id="auth-block-header">Вход</p>
                <p>Логин</p>
                <input type="text" id="login_auth"/>
                <p>Пароль</p>
                <input type="password" id="password_auth"/>
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