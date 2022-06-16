import { useEffect } from "react"
import { useNavigate } from "react-router-dom"
import { API_PROTECTED } from "../../util/api"

const Logout = () => {
    const navigate = useNavigate()

    useEffect(() => {
        setTimeout(() => {
            API_PROTECTED.get("/auths/logout").then(() => {
                localStorage.removeItem("userinfo")
                navigate("/", { replace: true })
            }).catch(() => {
                localStorage.removeItem("userinfo")
                navigate("/", { replace: true })
            })
        }, 1000)
    })

    return (
        <div className="modal" tabIndex="-1">
            <div className="modal-dialog">
                <div className="modal-content">
                <div className="modal-header">
                    <h5 className="modal-title">Logout</h5>
                </div>
                <div className="modal-body">
                    <p>You have been logged out.</p>
                </div>
                </div>
            </div>
        </div>
    )
}

export default Logout