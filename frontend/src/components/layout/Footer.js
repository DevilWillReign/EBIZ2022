import { NavLink } from "react-router-dom"

const Footer = () => {
    return (
        <footer id="footer" className="bg-light text-center text-lg-start">
            <div id="footer-links" className="container p-4">
                <div className="row">
                    <div className="col-lg-3 col-md-6 mb-4 mb-md-0">
                        <h5 className="text-uppercase">Links</h5>

                        <ul className="list-unstyled mb-0">
                            <li>
                                <NavLink to="/about">About</NavLink>
                            </li>
                        </ul>
                    </div>
                </div>
            </div>
            <div id="footer-copyright" className="text-center p-3" style={{backgroundColor: 'rgba(0, 0, 0, 0.2)'}}>
                © 2022 Copyright
            </div>
        </footer>
    )
}

export default Footer