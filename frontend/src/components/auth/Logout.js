import { useEffect } from "react"
import { useNavigate } from "react-router-dom"

const Logout = () => {
    const navigate = useNavigate()

    useEffect(() => {
        setTimeout(() => {
            localStorage.removeItem("userinfo")
            navigate("/", { replace: true })
        }, 1000)
    })

    return (
        <div className="modal" tabindex="-1">
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