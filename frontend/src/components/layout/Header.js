import { useState } from "react"
import { NavLink } from "react-router-dom"

const routes = (loggedIn) => [
    {
        path: "/",
        name: "Home"
    },
    {
        path: "/profile/cart",
        name: "Cart"
    },
    {
        path: "/products",
        name: "Products"
    },
    {
        path: loggedIn ? "/profile" : "/auth",
        name: loggedIn ? "Profile" : "Login"
    },
    {
        path: loggedIn ? "/auth/logout" : "/auth/register",
        name: loggedIn ? "Logout" : "Register"
    },
]

const Header = () => {
    var [loggedIn, setLoggedIn] = useState(sessionStorage.getItem("user") !== null)

    return (
        <nav className="navbar navbar-expand-lg navbar-light bg-light">
            <div className="container-fluid">
                <NavLink className="navbar-brand" to="/">Apprit Store</NavLink>
                <button className="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent"
                        aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
                    <span className="navbar-toggler-icon"></span>
                </button>
                <div className="collapse navbar-collapse" id="navbarSupportedContent">
                    <ul className="navbar-nav me-auto mb-2 mb-lg-0">
                        {
                            routes(loggedIn).map(route => {
                                return (
                                    <li key={route.path} className="nav-item">
                                        <NavLink className="nav-link active" aria-current="page" to={route.path}>{route.name}</NavLink>
                                    </li>
                                )
                            })
                        }
                    </ul>
                </div>
            </div>
        </nav>
    )
}

export default Header