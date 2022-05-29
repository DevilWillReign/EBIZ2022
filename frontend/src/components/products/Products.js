import { useEffect, useState } from "react"
import { NavLink } from 'react-router-dom'
import { API } from "../../util/api"

const Products = () => {
    const [products, setProducts] = useState([])

    useEffect(() => {
        API.get("/products").then((response) => {
            if (response.status === 200) {
                setProducts(response.data)
            }
        }).catch((reason) => {
            console.log(reason)
        })
    }, [])

    return (
        <ul id="product-list" className="list-group">
            {
                products.map(product => {
                    return (
                        <li className="list-group-item" key={product.id}>
                            <NavLink to={"" + product.id}>{product.code} {product.name}</NavLink>
                        </li>
                    )
                })
            }
        </ul>
    )
}

export default Products