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
import About from "../components/common/About"
import Logout from "../components/auth/Logout"
import Categories from "../components/categories/Categories"
import Category from "../components/categories/Category"
import Transaction from "../components/profile/Transaction"
import PayForm from "../components/profile/PayForm"

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
                path: ":productId",
                element: <Product />
            }
        ]
    },
    {
        path: "/categories",
        element: <MainLayout />,
        children: [
            {
                index: true,
                element: <Categories />
            },
            {
                path: ":categoryId",
                element: <Category />
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
            },
            {
                path: "logout",
                element: <Logout />
            }
        ]
    },
    {
        path: "/profile",
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
                path: "transactions/:transactionId",
                element: <Transaction />
            },
            {
                path: "cart",
                element: <Cart />
            },
            {
                path: "payments",
                element: <Payments />
            },
            {
                path: "payments/form/:transactionId",
                element: <PayForm />
            }
        ]
    },
    {
        path: "/about",
        element: <MainLayout />,
        children: [
            {
                index: true,
                element: <About />
            }
        ]
    }
]

const Router = () => {
    return useRoutes(routes)
}

export default Router