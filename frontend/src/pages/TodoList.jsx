import { useEffect, useState } from 'react'
import api from '../api.js'
import fileDownload from 'js-file-download'
import TodoItem from '../components/TodoItem'

export default function TodoList() {
    const [todos, setTodos] = useState([])
    const [title, setTitle] = useState('')

    const load = () => api.get('/todos').then(r => setTodos(r.data))
    useEffect(load, [])

    const add = async e => {
        e.preventDefault()
        if (!title.trim()) return
        await api.post('/todos', { title })
        setTitle('')
        load()
    }

    const download = async fmt => {
        const res = await api.get('/todos/export', {
            params: { fmt },
            responseType: 'blob',
        })
        fileDownload(res.data, `todos.${fmt}`)
    }

    return (
        <div className="max-w-lg mx-auto">
            <form onSubmit={add} className="flex space-x-2 mb-4">
                <input
                    value={title}
                    onChange={e => setTitle(e.target.value)}
                    placeholder="Новая задача"
                    className="flex-1 p-2 border rounded"
                />
                <button className="p-2 bg-blue-600 text-white rounded">+</button>
            </form>

            <div className="flex space-x-2 mb-4">
                <button
                    type="button"
                    onClick={() => download('json')}
                    className="px-3 py-1 rounded bg-blue-600 text-white hover:bg-blue-700"
                >
                    Экспорт JSON
                </button>
                <button
                    type="button"
                    onClick={() => download('csv')}
                    className="px-3 py-1 rounded bg-blue-600 text-white hover:bg-blue-700"
                >
                    Экспорт CSV
                </button>
            </div>

            <ul className="space-y-2">
                {todos.map(t => (
                    <TodoItem key={t.id} todo={t} reload={load}/>
                ))}
            </ul>
        </div>
    )
}
