import { useEffect, useState } from "react"
import { NavLink, useNavigate } from "react-router-dom"
import { API_PROTECTED } from "../../util/api"

const Transactions = () => {
    const [loggedIn] = useState(localStorage.getItem("userinfo") !== null)
    const [transactions, setTransactions] = useState([])
    const navigate = useNavigate()

    useEffect(() => {
        if (!loggedIn) {
            localStorage.setItem("userinfo", null)
            navigate("/auth/logout", { replace: true })
        }
        API_PROTECTED().get("/user/transactions").then(response => {
            setTransactions([...response.data])
        }).catch(() => {
            navigate("/auth/logout", { replace: true })
        })
    }, [loggedIn, navigate])

    return (
        <>
            <ul className="list-group">
                {
                    transactions.map(transaction => {
                        return (
                            <li className="list-group-item">
                                <NavLink to={"" + transaction.id}>{transaction.id + " " + transaction.createdat}</NavLink>
                            </li>
                        )
                    })
                }
            </ul>
            <NavLink className="btn btn-primary" to="/profile">Back to profile</NavLink>
        </>
    )
}

export default Transactions