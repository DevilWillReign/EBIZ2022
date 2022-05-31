import { useEffect, useState } from "react"
import { NavLink } from "react-router-dom"
import { API } from "../../util/api"

const Categories = () => {
    const [categories, setCategories] = useState([])

    useEffect(() => {
        API.get("/categories").then(response => {
            setCategories([...response.data])
        }).catch(() => {})
    }, [])

    return (
        <div className="card-group">
            {
                categories.map(category => {
                    return (
                        <div key={category.name} className="card">
                            <img className="card-img-top" src="..." alt="Card image cap" />
                            <div className="card-body">
                                <h5 className="card-title">{category.name}</h5>
                                <p className="card-text">{category.description}</p>
                                <NavLink className="btn btn-primary" to={"" + category.id}>Go category</NavLink>
                            </div>
                        </div>
                    )
                })
            }
        </div>
    )
}

export default Categories