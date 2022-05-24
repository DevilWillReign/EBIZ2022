import { useRoutes } from "react-router-dom"
import Cart from "../components/profile/Cart"
import Home from "../components/common/Home"
import Login from "../components/auth/Login"
import Payments from "../components/profile/Payments"
import Products from "../components/products/Products"
import Register from "../components/auth/Register"
import Profile from "../components/profile/Profile"
import MainLayout from "../components/MainLayout"
import Product from "../components/products/Product"
import Transactions from "../components/profile/Transactions"

const routes = [
    {
        path: "/",
        element: <MainLayout />,
        children: [
            {
                index: true,
                element: <Home />
            }
        ]
    },
    {
        path: "/products",
        element: <MainLayout />,
        children: [
            {
                index: true,
                element: <Products />
            },
            {
                path: ":id",
                element: <Product />
            }
        ]
    },
    {
        path: "/auth",
        element: <MainLayout />,
        children: [
            {
                index: true,
                element: <Login />
            },
            {
                path: "register",
                element: <Register />
            }
        ]
    },
    {
        path: "/profile/",
        element: <MainLayout />,
        children: [
            {
                index: true,
                element: <Profile />
            },
            {
                path: "transactions",
                element: <Transactions />
            },
            {
                path: "cart",
                element: <Cart />
            },
            {
                path: "payments",
                element: <Payments />
            }
        ]
    }

]

const Router = () => {
    return useRoutes(routes)
}

export default Router