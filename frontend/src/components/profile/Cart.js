import { useState } from "react"
import { NavLink } from "react-router-dom"

const Cart = () => {
    const [cart, setCart] = useState(sessionStorage.getItem("cart") == null ? [] : JSON.parse(sessionStorage.getItem("cart")))

    const buyForProducts = (event) => {

    }
    
    const removeProduct = (event, product) => {
        let currentCart = cart
        let index = currentCart.indexOf(product)
        if (index > -1) {
            currentCart.splice(index, 1)
            sessionStorage.setItem("cart", JSON.stringify(currentCart))
            setCart(currentCart)
        }
    }

    return (
        <>
            <ol class="list-group list-group-numbered">
                {
                    cart.map((element) => {
                        return (
                            <li id={element.id} key={element.id} className="list-group-item d-flex justify-content-between align-items-start">
                            <div class="ms-2 me-auto">
                                <div class="fw-bold"><NavLink to={"/products" + element.id}>{element.product_name}</NavLink></div>
                                {element.code}
                            </div>
                            <span class="badge bg-primary rounded-pill">{element.quantity}</span>
                            <button onClick={(e) => removeProduct(e, element)}>Remove</button>
                            </li>
                        )
                    })
                }
            </ol>
            <button onClick={(e) => buyForProducts(e)}>Pay</button>
        </>
    )
}

export default Cart