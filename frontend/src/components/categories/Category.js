import { useEffect, useState } from "react"
import { NavLink, useParams } from "react-router-dom"
import { API } from "../../util/api"

const Category = () => {
    const [ category, setCategory ] = useState({})
    const { categoryId } = useParams()

    useEffect(() => {
        API.get("/categories/" + categoryId + "/extended").then((response) => {
            if (response.status === 200) {
                setCategory(response.data)
            }
        }).catch((reason) => {
            console.log(reason)
        })
    }, [categoryId])

    return (
        <>
            <ul id="category" className="list-group list-group-flush">
                <li id="category-name" className="list-group-item">Category name: {category.name ? category.name : "no data"}</li>
                <li id="category-description" className="list-group-item">Category description: {category.description ? category.description : "no data"}</li>
                <li id="category-products" className="list-group-item">
                    <ul id="products" className="list-group">
                        {
                            category.products ? category.products.map(product => {
                                return (
                                    <li key={"product" + product.id} className="list-group-item">
                                        <span>
                                            <NavLink to={"/products/" + product.id}>
                                                {"name: "  +product.name + ", price: " + product.price + ", avaliability: " + product.availability}
                                            </NavLink>
                                        </span>
                                    </li>
                                )
                            }) : ""
                        }
                    </ul>
                </li>
            </ul>
            <NavLink id="category-back" to="/categories" className="btn btn-primary">Back to categories list</NavLink>
        </>
    )
}

export default Category