import { useEffect, useState } from "react"
import { NavLink, useNavigate } from "react-router-dom"

const Profile = () => {
    const [profile, _setProfile] = useState(localStorage.getItem("userinfo") ? JSON.parse(localStorage.getItem("userinfo")) : {})
    const navigate = useNavigate()

    useEffect(() => {
        if (profile || Object.keys(profile).length === 0 || Object.getPrototypeOf(profile) === Object.prototype) {
            navigate("/auth", { replace: true })
        }
    }, [])

    return (
        <>
            <ul id="user-info" className="list-group">
                <li id="user-name" className="list-group-item">Name: {profile.name}</li>
                <li id="user-email" className="list-group-item">Email: {profile.email}</li>
            </ul>
            <NavLink className="btn btn-primary" to="/profile/cart">Cart</NavLink>
            <NavLink className="btn btn-primary" to="/profile/transactions">Transactions</NavLink>
            <NavLink className="btn btn-primary" to="/profile/cart">Cart</NavLink>
        </>
    )
}

export default Profile