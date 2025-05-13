import axios from 'axios'

const api = axios.create({ baseURL: '77.221.135.36:8080' })

api.interceptors.request.use(cfg => {
    const tk = localStorage.getItem('tk')
    if (tk) cfg.headers.Authorization = `Bearer ${tk}`
    return cfg
})

export default api
