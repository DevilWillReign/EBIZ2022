import { useEffect, useState } from "react"
import API from "../util/api"
import { myRange } from "../util/util"

const Products = () => {
    var [products, setProducts] = useState([])

    useEffect(() => {
        console.log(API.defaults)
        API.get("/products").then((response) => {
            if (response.status === 200) {
                const products = response.data
                setProducts(products)
            }
        })
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
                <tbody>
                {
                    products.map((element, i) => {
                        var quantity = []
                        for (const k of myRange(10, 1)) {
                            quantity.push(<option key={k} value={k}>{k}</option>)
                        }
                        return (
                            <tr key={element.id}>
                                <td>{element.id}</td>
                                <td>{element.name}</td>
                                <td>{element.code}</td>
                                <td>{element.price}</td>
                                <td>
                                    <select id={"quantity" + element.id}>
                                        {quantity}
                                    </select>
                                </td>
                                <td><button onClick={(event) => addProductToBasket(event, element)}>Add product to basket</button></td>
                            </tr>
                        )
                    })
                }
                </tbody>
            </table>
        </div>
    )
}

export default Products