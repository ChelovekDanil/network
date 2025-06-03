import { useNavigate } from 'react-router-dom'
import './profile.css'

function Profile() {
    const login = localStorage.getItem("login")

    const navigate = useNavigate()

    const handlerGoBack = () => {
        navigate(-1)
    }
    
    return (
        <>
            <button onClick={handlerGoBack}>Назад</button>
            <div className="profile-page">
                <img src="./icon.png" alt="icon" />
                <p>Логин</p>
                <input type="text" value={login} />    
                <p>Пароль</p>
                <input type="text" />
                <button>Изменить</button>
                <button>Удалить пользователя</button>
            </div>
        </>
    )
}

export default Profile
