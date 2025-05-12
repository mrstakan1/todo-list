import React from 'react'
import ReactDOM from 'react-dom/client'
import { HashRouter, Routes, Route, Navigate } from 'react-router-dom'
import App from './App'
import Login from './pages/Login'
import Register from './pages/Register'
import TodoList from './pages/TodoList'
import './index.css'

ReactDOM.createRoot(document.getElementById('root')).render(
    <HashRouter>
        <Routes>
            <Route path="/" element={<App />}>
                <Route index element={<Navigate to="/todos" replace />} />
                <Route path="login" element={<Login />} />
                <Route path="register" element={<Register />} />
                <Route path="todos" element={<TodoList />} />
            </Route>
        </Routes>
    </HashRouter>
)
