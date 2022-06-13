import { useEffect, useState } from "react"
import { NavLink } from "react-router-dom"
import CategoryDropdown from "../categories/CategoryDropdown"
import Link from "./Link"

const routes = (loggedIn) => {
    return (
        <>
            <Link path={"/"} name={"Home"} />
            <Link path={"/profile/cart"} name={"Cart"} />
            <CategoryDropdown path={"/categories"} name={"Categories"} />
            <Link path={"/products"} name={"Products"} />
            <Link path={loggedIn ? "/profile" : "/auth"} name={loggedIn ? "Profile" : "Login"} />
            <Link path={loggedIn ? "/auth/logout" : "/auth/register"} name={loggedIn ? "Logout" : "Register"} />
        </>
    )
}

const Header = () => {
    const [loggedIn, setLoggedIn] = useState(localStorage.getItem("userinfo") !== null)
    const checkLocalStorage = () => localStorage.getItem("userinfo")
    
    useEffect(() => {
        setLoggedIn(localStorage.getItem("userinfo") !== null)
    }, [checkLocalStorage])

    return (
        <nav className="navbar navbar-expand-lg navbar-light bg-light">
            <div id="header-container" className="container-fluid">
                <NavLink className="navbar-brand" to="/">Apprit Store</NavLink>
                <button className="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent"
                        aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
                    <span className="navbar-toggler-icon"></span>
                </button>
                <div className="collapse navbar-collapse" id="navbarSupportedContent">
                    <ul className="navbar-nav me-auto mb-2 mb-lg-0">
                        {
                            routes(loggedIn)
                        }
                    </ul>
                </div>
            </div>
        </nav>
    )
}

export default Header