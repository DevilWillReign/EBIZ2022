import { useEffect, useState } from "react"
import { API } from "../../util/api"
import Dropdown from "../layout/Dropdown"

const CategoryDropdown = (props) => {
    const [categories, setCategories] = useState([])

    useEffect(() => {
        API.get("/categories").then(response => {
            setCategories([...response.data.elements])
        }).catch(() => {})
    }, [])

    return (
        <Dropdown path={props.path} name={props.name} elements={categories} />
    )
}

export default CategoryDropdown