import { useState } from "react"

const Cart = () => {
    var [cart, setCart] = useState(sessionStorage.getItem("cart") == null ? [] : JSON.parse(sessionStorage.getItem("cart")))

    const buyForProducts = (event) => {

    }

    return (
        <div className="App">
            <table>
                {
                    cart.map((element, i) => {
                        return <tr><td>element.product_name</td><td>element.quantity</td></tr>
                    })
                }
            </table>
            <button onClick={(e) => buyForProducts(e)}>Pay</button>
        </div>
    )
}

export default Cart