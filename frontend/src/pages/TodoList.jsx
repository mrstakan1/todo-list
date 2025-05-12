import { useEffect, useState } from 'react'
import api from '../api'
import TodoItem from '../components/TodoItem'

export default function TodoList() {
    const [todos, set] = useState([])
    const [title, setTitle] = useState('')
    const load = () => api.get('/todos').then(r => set(r.data))
    useEffect(load, [])
    const add = async e => { e.preventDefault(); await api.post('/todos', { title }); setTitle(''); load() }
    return (
        <div className="max-w-lg mx-auto">
            <form onSubmit={add} className="flex space-x-2 mb-4">
                <input value={title} onChange={e => setTitle(e.target.value)}
                       placeholder="Новая задача" className="flex-1 p-2 border"/>
                <button className="p-2 bg-blue-600 text-white">+</button>
            </form>
            <ul className="space-y-2">
                {todos.map(t => <TodoItem key={t.id} todo={t} reload={load}/>)}
            </ul>
        </div>
    )
}
