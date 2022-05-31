import { useEffect, useState } from "react"
import { NavLink, useNavigate } from "react-router-dom"
import { API_PROTECTED } from "../../util/api"

const Profile = () => {
    const [loggedIn] = useState(localStorage.getItem("userinfo") !== null)
    const navigate = useNavigate()
    const [profile, setProfile] = useState(null)

    useEffect(() => {
        if (!loggedIn) {
            localStorage.setItem("userinfo", null)
            navigate("/auth/logout", { replace: true })
        }
        API_PROTECTED.get("/user/me").then(response => {
            if (response.status === 200) {
                setProfile(response.data)
            }
        }).catch(() => {
            navigate("/auth/logout", { replace: true })
        })
    }, [navigate, loggedIn])

    if (profile === null) {
        return (
            <div>Loading...</div>
        )
    } else {
        return (
            <>
                <ul id="user-info" className="list-group">
                    <li id="user-name" className="list-group-item">Name: {profile.name}</li>
                    <li id="user-email" className="list-group-item">Email: {profile.email}</li>
                </ul>
                <NavLink className="btn btn-primary" to="/profile/cart">Cart</NavLink>
                <NavLink className="btn btn-primary" to="/profile/payments">Payments</NavLink>
                <NavLink className="btn btn-primary" to="/profile/transactions">Transactions</NavLink>
            </>
        )
    }
}

export default Profile