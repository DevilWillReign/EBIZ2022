import { useEffect, useState } from "react"
import API from "../util/api"

const Products = () => {
    var [products, setProducts] = useState([])

    useEffect(() => {
        if (!products) {
            API.get("/products").then((response) => {
                const products = JSON.parse(response.data)
                setProducts(products)
            })
        }
    }, [])

    const addProductToBasket = (event, product) => {
        let productWithQuantity = product
        productWithQuantity.quantity = Math.random() * 100 + 1
        let basket = sessionStorage.getItem("basket") == null ? JSON.parse(sessionStorage.getItem("basket")) : []
        basket.push(productWithQuantity)
        sessionStorage.setItem("basket", JSON.stringify(basket))
    }

    return (
        <div className="App">
            <table>
                {
                    products.map((element, i) => {
                        return <tr>
                            <td>element.id</td>
                            <td>element.name</td>
                            <td>element.code</td>
                            <td>element.price</td>
                            <td><button onClick={(event) => addProductToBasket(event, element)}></button></td>
                        </tr>
                    })
                }
            </table>
        </div>
    )
}

export default Products