import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import api from '../api'

export default function Register() {
    const [d, set] = useState({ username: '', password: '' })
    const nav = useNavigate()
    const submit = async e => {
        e.preventDefault()
        await api.post('/register', d).catch(() => alert('Пользователь существует'))
        nav('/login')
    }
    return (
        <form onSubmit={submit} className="max-w-sm mx-auto space-y-4">
            <input value={d.username} onChange={e => set({ ...d, username: e.target.value })}
                   placeholder="Имя" className="w-full p-2 border"/>
            <input type="password" value={d.password}
                   onChange={e => set({ ...d, password: e.target.value })}
                   placeholder="Пароль" className="w-full p-2 border"/>
            <button className="w-full p-2 bg-green-600 text-white">Создать аккаунт</button>
        </form>
    )
}
