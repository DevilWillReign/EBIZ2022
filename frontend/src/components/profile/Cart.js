import { useState } from "react"
import { NavLink, useNavigate } from "react-router-dom"
import { API_PROTECTED } from "../../util/api"

const Cart = () => {
    const [loggedIn, ] = useState(localStorage.getItem("userinfo") !== null)
    const [cart, setCart] = useState(localStorage.getItem("cart") ? JSON.parse(localStorage.getItem("cart")) : [])
    const navigate = useNavigate()

    const buyForProducts = () => {
        API_PROTECTED.post("/user/transactions", {quantifiedproducts: cart})
        .then(response => {
            if (response.status === 201) {
                localStorage.removeItem("cart")
                navigate("/profile/payments/form/" + response.data.transactionid)
            }
        }).catch(() => {console.log("TRANSACTION ERROR")})
    }
    
    const removeProduct = (product) => {
        let currentCart = cart
        let index = currentCart.indexOf(product)
        if (index > -1) {
            currentCart.splice(index, 1)
            localStorage.setItem("cart", JSON.stringify(currentCart))
        }
        setCart([...currentCart])
    }

    const button = (loggedIn) => {
        return loggedIn ? <button id="cart-pay" className="btn btn-primary" onClick={() => buyForProducts()}>Go to payment</button>
        : <span>Login to buy products</span>
    }

    return (
        <>
            { cart.length === 0 ? 
                <h3 id="cart-empty">Cart empty</h3>
                : (
                    <>
                        <ol id="cart-list" className="list-group list-group-numbered">
                            {
                                cart.map((product) => {
                                    return (
                                        <li id={product.id} key={product.id} className="list-group-item d-flex justify-content-between align-items-start">
                                        <div className="ms-2 me-auto">
                                            <div className="fw-bold"><NavLink to={"/products" + product.id}>{product.name}</NavLink></div>
                                            {product.code}
                                        </div>
                                        <span className="badge bg-primary rounded-pill">{product.quantity}</span>
                                        <button className="btn btn-close" onClick={() => removeProduct(product)}></button>
                                        </li>
                                    )
                                })
                            }
                        </ol>
                            { button(loggedIn) }
                    </>
                ) 
            }
        </>
    )
}

export default Cart