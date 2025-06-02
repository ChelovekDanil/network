import './profile.css'

function Profile() {
    const login = localStorage.getItem("login")
    
    return (
        <>
            <div className="profile-page">
                <img src="./icon.png" alt="icon" />
                <p>Логин</p>
                <input type="text" value={login} />    
                <p>Пароль</p>
                <input type="text" />
                <button>Изменить</button>
            </div>
        </>
    )
}

export default Profile
