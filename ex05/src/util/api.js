import axios from "axios"

const API = axios.create({
    baseURL: process.env.API_BASE_URL || "http://localhost:9000/api/v1"
})

export default API