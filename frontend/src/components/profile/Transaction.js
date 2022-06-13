import { useEffect, useState } from "react"
import { NavLink, useNavigate, useParams } from "react-router-dom"
import { API_PROTECTED } from "../../util/api"
import TransactionProducts from "./TransactionProducts"

const Transaction = () => {
    const [loggedIn, ] = useState(localStorage.getItem("userinfo") !== null)
    const [transaction, setTransaction] = useState(null)
    const navigate = useNavigate()
    const { transactionId } = useParams()

    useEffect(() => {
        if (!loggedIn) {
            localStorage.setItem("userinfo", null)
            navigate("/auth/logout", { replace: true })
        }
        API_PROTECTED.get("/user/transactions/" + transactionId).then(response => {
            if (response.status === 200) {
                console.log(response.data)
                setTransaction(response.data)
            }
        }).catch(() => {
            navigate("/auth/logout", { replace: true })
        })
    }, [loggedIn, navigate, transactionId])

    if (transaction === null) {
        return (
            <div>Loading...</div>
        )
    } else {
        return (
            <>
                <ul id="transaction-info" className="list-group">
                    <li id="transaction-date" className="list-group-item">Date: {transaction.createdat}</li>
                    <li id="transaction-total" className="list-group-item">Total: {transaction.total}</li>
                    <li id="transaction-payment" className="list-group-item">Payment: {transaction.payment || transaction.payment.id !== 0 ? "Paid" : "Not paid"}</li>
                    <li id="transaction-products-label" className="list-group-item">Products:</li>
                    <TransactionProducts products={transaction.quantifiedproducts} />
                </ul>
                <NavLink className="btn btn-primary" to="/profile/transactions">Back to transactions</NavLink>
            </>
        )
    }
}

export default Transaction