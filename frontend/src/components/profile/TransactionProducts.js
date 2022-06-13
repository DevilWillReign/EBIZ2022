import { NavLink } from "react-router-dom"

const TransactionProducts = (props) => {
    return (
        <ul id="transaction-products" className="list-group">
            {
                props.products.map(product => {
                    return (
                        <li key={product.id} id={"transaction-product-" + product.id} className="list-group-item">
                            <NavLink to={"/products/" + product.id}>Name: {product.name}, Quantity: {product.quantity}, Price: {product.price}</NavLink>
                        </li>
                    )
                })
            }
        </ul>
    )
}

export default TransactionProducts