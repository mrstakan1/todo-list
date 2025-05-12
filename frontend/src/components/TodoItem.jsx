import api from '../api'

export default function TodoItem({ todo, reload }) {
    const toggle = () => api.put(`/todos/${todo.id}`, { ...todo, completed: !todo.completed }).then(reload)
    const del = () => api.delete(`/todos/${todo.id}`).then(reload)
    return (
        <li className="p-2 bg-white flex justify-between items-center">
            <label className="flex items-center space-x-2">
                <input type="checkbox" checked={todo.completed} onChange={toggle}/>
                <span className={todo.completed ? 'line-through' : ''}>{todo.title}</span>
            </label>
            <button
                onClick={del}
                className="text-red-600 hover:text-red-800 text-xl leading-none"
                title="Удалить"
            >
                &times;
            </button>
        </li>
    )
}
