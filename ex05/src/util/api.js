import axios from "axios"

const API = axios.create({
    baseURL: process.env.API_BASE_URL || "localhost:9000"
})

export default API