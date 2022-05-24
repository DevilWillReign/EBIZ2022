import { useEffect, useState } from "react"
import { NavLink } from 'react-router-dom'
import API from "../../util/api"

const Products = () => {
    const [products, setProducts] = useState([])

    useEffect(() => {
        API.get("/products").then((response) => {
            if (response.status === 200) {
                const products = response.data
                setProducts(products)
            }
        }).catch((reason) => {
            console.log(reason)
        })
    }, [])

    return (
        <ul id="product-list" className="list-group">
            {
                products.map(element => {
                    return (
                        <li className="list-group-item" key={element.id}>
                            <NavLink to={element.id}>{element.code} {element.name}</NavLink>
                        </li>
                    )
                })
            }
        </ul>
    )
}

export default Products