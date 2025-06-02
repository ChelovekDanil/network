import React, { useState } from 'react';
import { useLocation } from 'react-router-dom';
import './contact.css'; // Предполагается, что этот файл существует для стилей

const ContactPage = () => {
    const location = useLocation();
    const { login } = location.state || {};

    const [message, setMessage] = useState('');
    const [messages, setMessages] = useState([]); // Состояние для массива сообщений
    const [sender, setSender] = useState("вы"); // Имя отправителя
    const [receiver, setReceiver] = useState('Пользователь'); // Имя получателя

    const handleSubmit = (e) => {
        e.preventDefault();

        // Добавление сообщения в массив
        if (message.trim()) { // Проверяем, что сообщение не пустое
            const msgObject = {
                content: message.trim(),
                sender,
                receiver,
                timestamp: new Date().toLocaleTimeString(), // Добавьте временную метку, если нужно
            };
            setMessages([...messages, msgObject]);
            setMessage(''); // Очистить поле ввода после отправки
        }
    };

    return (
        <div className="contact-page">
            {login ? <p className="message">Сообщение: {login}</p> : <p className="message">Нет выбранного пользователя.</p>}
            <div className="messages-list">
                <h3>Список сообщений:</h3>
                {messages.length === 0 ? (
                    <p>Нет сообщений.</p>
                ) : (
                    <ul>
                        {messages.map((msg, index) => (
                            <li key={index}>
                                <strong>{msg.sender}</strong> (выслано в {msg.timestamp}):
                                <p>{msg.content}</p>
                            </li>
                        ))}
                    </ul>
                )}
            </div>
            <form onSubmit={handleSubmit} className="contact-form">
                <textarea
                    className="input-message"
                    value={message}
                    onChange={(e) => setMessage(e.target.value)}
                    placeholder="Введите ваше сообщение"
                    required
                />
                <button type="submit" className="send-button">Отправить</button>
            </form>
        </div>
    );
};

export default ContactPage;
