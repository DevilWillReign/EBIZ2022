import { NavLink } from "react-router-dom"

const Dropdown = (props) => {
    return (
        <li key={props.path} className="nav-item">
            <NavLink type="button" className="btn" to={props.path}>{props.name}</NavLink>
            <button type="button" className="btn dropdown-toggle dropdown-toggle-split" data-bs-toggle="dropdown" aria-expanded="false">
                <span className="visually-hidden">Toggle Dropdown</span>
            </button>
            <ul className="dropdown-menu">
                {
                    props.elements.map(element => {
                        return (
                            <li key={props.path + "/" + element.id}>
                                <NavLink className="dropdown-item" to={props.path + "/" + element.id}>{element.name}</NavLink>
                            </li>
                        )
                    })
                }
            </ul>
        </li>
    )
}

export default Dropdown