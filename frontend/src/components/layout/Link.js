import { NavLink } from "react-router-dom"

const Link = (props) => {
    return (
        <li key={props.path} className="nav-item">
            <NavLink className="nav-link active" aria-current="page" to={props.path}>{props.name}</NavLink>
        </li>
    )
}

export default Link