import './addContact.css'

function AddContact() {
    return (
        <>
            <div className="add-contact-page">
                <p>Добавить контакт</p>
                <p id='login-contact'>Логин контакта</p>
                <input type="text" />
                <button>Добавить</button>
            </div>
        </>
    )
}

export default AddContact