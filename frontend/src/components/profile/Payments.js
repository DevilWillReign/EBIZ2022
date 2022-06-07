import { useEffect, useState } from "react"
import { NavLink, useNavigate } from "react-router-dom"
import { API_PROTECTED } from "../../util/api"

const Payments = () => {
    const [loggedIn] = useState(localStorage.getItem("userinfo") !== null)
    const [payments, setPayments] = useState([])
    const navigate = useNavigate()

    useEffect(() => {
        if (!loggedIn) {
            localStorage.setItem("userinfo", null)
            navigate("/auth/logout", { replace: true })
        }
        API_PROTECTED().get("/user/payments").then(response => {
            if (response.status === 200) {
                setPayments(response.data)
            }
        }).catch(() => {
            navigate("/auth/logout", { replace: true })
        })
    }, [loggedIn, navigate])

    return (
        <>
            <ul className="list-group">
                {
                    payments.map(payment => {
                        return (
                            <li className="list-group-item">
                                <NavLink to={"" + payment.id}>{payment.id + " " + payment.paymenttype}</NavLink>
                            </li>
                        )
                    })
                }
            </ul>
            <NavLink className="btn btn-primary" to="/profile">Back to profile</NavLink>
        </>
    )
}

export default Payments