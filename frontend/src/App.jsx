import { Outlet, Link, useNavigate, useLocation, Navigate } from 'react-router-dom'

export default function App() {
    const nav = useNavigate()
    const loc = useLocation()
    const tk = localStorage.getItem('tk')
    if (!tk && !['/login', '/register'].includes(loc.pathname))
        return <Navigate to="/login" replace />

    const logout = () => { localStorage.removeItem('tk'); nav('/login') }

    return (
        <div className="min-h-screen bg-gray-100 p-4">
            <header className="flex justify-between mb-4">
                <h1 className="text-2xl font-bold"><Link to="/todos">Todo-App</Link></h1>
                {tk && (
                    <nav className="space-x-4">
                        <Link to="/todos">Список</Link>
                        <button
                            onClick={logout}
                            className="text-red-600 hover:text-red-700 transition-colors"
                        >
                            Выйти
                        </button>
                    </nav>
                )}
            </header>
            <Outlet/>
        </div>
    )
}
