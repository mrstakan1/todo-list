import axios from 'axios'

export default axios.create({
    // локально переопределяем через VITE_API_URL,
    // в контейнере запрос идёт на /api/*
    baseURL: import.meta.env.VITE_API_URL || '/api'
})