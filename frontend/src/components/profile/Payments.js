import { useEffect, useState } from "react"
import { NavLink, useNavigate } from "react-router-dom"
import { API_PROTECTED } from "../../util/api"

const Payments = () => {
    const [loggedIn, ] = useState(localStorage.getItem("userinfo") !== null)
    const [payments, setPayments] = useState([])
    const navigate = useNavigate()

    useEffect(() => {
        if (!loggedIn) {
            localStorage.setItem("userinfo", null)
            navigate("/auth/logout", { replace: true })
        }
        API_PROTECTED.get("/user/payments").then(response => {
            if (response.status === 200) {
                setPayments([...response.data.elements])
            }
        }).catch(() => {
            navigate("/auth/logout", { replace: true })
        })
    }, [loggedIn, navigate])

    return (
        <>
            <ul id="payments-list" className="list-group">
                {
                    payments.map(payment => {
                        return (
                            <li id={"payments-item-" + payment.id} key={payment.id} className="list-group-item">
                                <span id={"payments-item-" + payment.id + "-text"}>{payment.id + " " + payment.paymenttype}</span>
                                <NavLink id={"payments-item-" + payment.id + "-transaction"} to={"/profile/transactions/" + payment.transactionid}>Go to transaction</NavLink>
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