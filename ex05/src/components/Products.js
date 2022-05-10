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

    return (
        <div className="App">
            <table>
                {
                    products.map((element, i) => {
                        return <tr><td>element.id</td><td>element.name</td><td>element.code</td><td>element.price</td></tr>
                    })
                }
            </table>
        </div>
    )
}

export default Products