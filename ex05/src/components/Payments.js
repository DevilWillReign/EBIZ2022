import { useState } from "react"
import API from "../util/api"

const Payments = () => {
    var [payments, setPayments] = useState([])

    var addPayment = (event) => {
        var form = document.forms.form1
        event.preventDefault()
    }

    return (
        <div>
        <form name="form1" onSubmit={addPayment}>
            <input name="userId" hidden={true} value={1} />
            <input name="" />
            <input />
        </form>
        <table>
            {

            }
        </table>
        </div>
    )
}

export default Payments