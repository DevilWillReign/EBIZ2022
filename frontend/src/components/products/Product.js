import { useEffect, useState } from "react"
import { useParams } from "react-router-dom"
import myRange from "../../util"

const Product = () => {
    const { product, setProduct } = useState({})
    const { productId } = useParams()

    useEffect(() => {
        console.log(API.defaults)
        API.get("/products/" + productId).then((response) => {
            if (response.status === 200) {
                const product = response.data
                setProduct(product)
            }
        }).catch((reason) => {
            console.log(reason)
        })
    }, [])

    const addProductToCart = (event) => {
        let productWithQuantity = product
        productWithQuantity.quantity = document.getElementById("product-quantity").value
        let cart = sessionStorage.getItem("cart") == null ? JSON.parse(sessionStorage.getItem("cart")) : []
        let index = cart.indexOf(productWithQuantity)
        if (index > -1) {
            cart[index] = productWithQuantity
        } else {
            cart.push(productWithQuantity)
        }
        sessionStorage.setItem("cart", JSON.stringify(cart))
    }

    return (
        <>
            <ul id="product" className="list-group list-group-flush">
                <li id="product-name" className="list-group-item">Product name: {product.name ? product.name : "no data"}</li>
                <li id="product-code" className="list-group-item">Product code: {product.code ? product.code : "no data"}</li>
                <li id="product-availability" className="list-group-item">Product availability: {product.availability ? product.availability : "no data"}</li>
                <li id="product-price" className="list-group-item">Product price: {product.price ? product.price : "no data"}</li>
                <li id="product-description" className="list-group-item">Product description: {product.description ? product.description : "no data"}</li>
            </ul>
            <select id="product-quantity">
            {myRange(product.availability ? product.availability : 0, 1).map(value => {
                return (
                    <option value={value}>{value}</option>
                )
            })}
            </select>
            <button id="add-product" className="btn btn-lg btn-primary" onClick={(e) => addProductToCart(e)}>Add to cart</button>
        </>
    )
}

export default Product