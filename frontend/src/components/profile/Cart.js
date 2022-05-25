import { useState } from "react"
import { NavLink } from "react-router-dom"

const Cart = () => {
    const [cart, setCart] = useState(localStorage.getItem("cart") ? JSON.parse(localStorage.getItem("cart")) : [])

    const buyForProducts = () => {
        console.log("Paying")
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
                        <button id="cart-pay" className="btn btn-primary" onClick={(e) => buyForProducts(e)}>Pay</button>
                    </>
                ) 
            }
        </>
    )
}

export default Cart