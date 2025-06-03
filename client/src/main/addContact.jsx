import { useNavigate } from 'react-router-dom'
import './addContact.css'

function AddContact() {
    const navigate = useNavigate();

    const handlerGoBack = () => {
        navigate(-1)
    }

    const addContact = () => {

    }

    return (
        <>
            <div className='add-contact-page'>
                <button onClick={handlerGoBack}>Назад</button>
                <div className="add-contact-page-block">    
                    <p>Добавить контакт</p>
                    <p id='login-contact'>Логин контакта</p>
                    <input type="text" />
                    <button onClick={addContact}>Добавить</button>
                </div>
            </div>
        </>
    )
}

export default AddContact