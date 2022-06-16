import { useEffect, useState } from "react"
import { NavLink } from "react-router-dom"
import { API } from "../../util/api"

const Categories = () => {
    const [categories, setCategories] = useState([])

    useEffect(() => {
        API.get("/categories").then(response => {
            setCategories([...response.data.elements])
        }).catch(() => {console.log("CATEGORIES ERROR")})
    }, [])

    return (
        <div id="category-list" className="card-group">
            {
                categories.map(category => {
                    const prefix = "category-" + category.name
                    return (
                        <div id={prefix} key={category.name} className="card">
                            <img id={prefix + "-img"} className="card-img-top" src="..." alt={category.name} />
                            <div id={prefix + "-body"} className="card-body">
                                <h5 id={prefix + "-name"} className="card-title">{category.name}</h5>
                                <p id={prefix + "-description"} className="card-text">{category.description}</p>
                                <NavLink className="btn btn-primary" to={"" + category.id}>Go to category</NavLink>
                            </div>
                        </div>
                    )
                })
            }
        </div>
    )
}

export default Categories