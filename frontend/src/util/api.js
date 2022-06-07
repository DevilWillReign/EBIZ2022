import axios from "axios"

const API = () => axios.create({
    baseURL: process.env.REACT_APP_API_BASE_URL || "http://localhost:9000/api/v1"
})

const API_PROTECTED = () => axios.create({
    baseURL: process.env.REACT_APP_API_BASE_URL || "http://localhost:9000/api/v1",
    withCredentials: true
})

export {
    API,
    API_PROTECTED
}