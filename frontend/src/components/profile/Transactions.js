import { useEffect, useState } from "react"
import { NavLink, useNavigate } from "react-router-dom"
import { API_PROTECTED } from "../../util/api"

const Transactions = () => {
    const [loggedIn, ] = useState(localStorage.getItem("userinfo") !== null)
    const [transactions, setTransactions] = useState([])
    const navigate = useNavigate()

    useEffect(() => {
        if (!loggedIn) {
            localStorage.setItem("userinfo", null)
            navigate("/auth/logout", { replace: true })
        }
        API_PROTECTED.get("/user/transactions").then(response => {
            if (response.status === 200) {
                setTransactions([...response.data.elements])
            }
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
                            <li key={transaction.id} className="list-group-item">
                                <NavLink to={"" + transaction.id}>Total: {transaction.total + ", Date: " + transaction.createdat}</NavLink>
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