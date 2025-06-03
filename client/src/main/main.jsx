import { useNavigate } from 'react-router-dom'
import './main.css'
import { useState } from 'react'

function Main() {
    const navigate = useNavigate()    

    const unRegister = () => {
        localStorage.setItem("isAuth", "false")
        navigate("/")
    }

    const goProfile = () => {
        navigate("/profile")
    }

    const clickContact = (user) => {
        navigate("/contact", { state: { login: user.login } });
    };

    const addContact = () => {
        navigate("/addContact")
    }

    const users = [
        { login: 'danil' },
        { login: 'ivan' },
        { login: 'anna' },
        { login: 'maria' }
    ];

    return (
        <> 
            <div className="main-page">
                <button onClick={unRegister}>Выйти</button>
                <button onClick={goProfile}>Профиль</button>
                <button onClick={addContact}>Добавить</button>
                <h1>Список контактов</h1>
                <ul>
                    {users.map((user, index) => (
                        <li 
                            className="order-items-main-page" 
                            onClick={() => clickContact(user)} 
                            key={index}
                        >
                            {user.login}
                        </li>
                    ))}
                </ul>
            </div> 
        </>
    );
}

export default Main;