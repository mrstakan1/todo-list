import { useState } from 'react'
import { useNavigate, Link } from 'react-router-dom'
import api from '../api'

export default function Login() {
    const [d, set] = useState({ username: '', password: '' })
    const nav = useNavigate()
    const submit = async e => {
        e.preventDefault()
        const { data } = await api.post('/login', d).catch(() => alert('Ошибка'))
        if (data?.token) { localStorage.setItem('tk', data.token); nav('/todos') }
    }
    return (
        <form onSubmit={submit} className="max-w-sm mx-auto space-y-4">
            <input value={d.username} onChange={e => set({ ...d, username: e.target.value })}
                   placeholder="Имя" className="w-full p-2 border"/>
            <input type="password" value={d.password}
                   onChange={e => set({ ...d, password: e.target.value })}
                   placeholder="Пароль" className="w-full p-2 border"/>
            <button className="w-full p-2 bg-blue-600 text-white">Войти</button>
            <Link to="/register" className="text-blue-600">Регистрация</Link>
        </form>
    )
}
