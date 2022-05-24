import Footer from "./layout/Footer"
import Header from "./layout/Header"
import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap/dist/js/bootstrap.min.js'
import { Outlet } from "react-router-dom"

const MainLayout = () => {
    return (
        <>
            <Header />
            <div className="container-fluid">
                <Outlet />
            </div>
            <Footer />
        </>
    )
}

export default MainLayout