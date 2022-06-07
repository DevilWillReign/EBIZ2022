import { useEffect, useState } from "react"
import { NavLink, useParams } from "react-router-dom"
import myRange from "../../util/myRange"
import { API } from "../../util/api"

const Product = () => {
    const [ product, setProduct ] = useState({})
    const [ quantity, setQuantity ] = useState(1)
    const { productId } = useParams()

    useEffect(() => {
        API().get("/products/" + productId).then((response) => {
            if (response.status === 200) {
                setProduct(response.data)
            }
        }).catch((reason) => {
            console.log(reason)
        })
    }, [productId])
    

    const findProductInArray = (array, productToFind) => {
        for (var i = 0; i < array.length; i++) {
            if (array[i].code === productToFind.code) {
                return i
            }
        }
        return -1
    }

    const addProductToCart = () => {
        let productWithQuantity = product
        productWithQuantity.quantity = Number(quantity)
        let cart = localStorage.getItem("cart") ? JSON.parse(localStorage.getItem("cart")) : []
        let index = findProductInArray(cart, productWithQuantity)
        if (index > -1) {
            cart[index] = productWithQuantity
        } else {
            cart.push(productWithQuantity)
        }
        localStorage.setItem("cart", JSON.stringify(cart))
    }

    return (
        <>
            <ul id="product" className="list-group list-group-flush">
                <li id="product-name" className="list-group-item">Product name: {product.name ? product.name : "no data"}</li>
                <li id="product-code" className="list-group-item">Product code: {product.code ? product.code : "no data"}</li>
                <li id="product-availability" className="list-group-item">Product availability: {product.availability ? product.availability : "no data"}</li>
                <li id="product-price" className="list-group-item">Product price: {product.price ? product.price : "no data"}</li>
                <li id="product-description" className="list-group-item">Product description: {product.description ? product.description : "no data"}</li>
                <li id="product-quantity" className="list-group-item">
                    <select onChange={(e) => setQuantity(e.target.value)} value={quantity} id="product-quantity-selector">
                        {myRange(product.availability ? product.availability : 0, 1).map(value => {
                            return (
                                <option key={value} value={value}>{value}</option>
                            )
                        })}
                    </select>
                </li>
            </ul>
            <button id="add-product" className="btn btn-primary" onClick={() => addProductToCart()}>Add to cart</button>
            <NavLink to="/products" className="btn btn-primary">Back to product list</NavLink>
        </>
    )
}

export default Product