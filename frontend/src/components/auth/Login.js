import { Formik, Field, Form, ErrorMessage } from "formik"
import * as Yup from "yup"
import { useNavigate } from "react-router-dom"
import { API, API_PROTECTED } from "../../util/api"
import './Login.css'
import { useCookies } from "react-cookie"
import { useEffect, useState } from "react"

const Login = () => {
    const api_url = API.defaults.baseURL
    const location = "/auth"
    const navigate = useNavigate()
    const [cookies, , removeCookies] = useCookies(["login_state"])
    const [failed, setFailed] = useState(false)

    useEffect(() => {
        if (cookies.login_state === "success") {
            localStorage.setItem("userinfo", true)
            removeCookies("login_state", {path: "/auth"})
            navigate("/", { replace: true })
        } else if (cookies.login_state === "failure") {
            removeCookies("login_state", {path: "/auth"})
            setFailed(true)
        }
    }, [cookies.login_state, navigate, removeCookies])

    return (
        <div className="text-center">
            <div className="form-signin">
                <Formik
                    initialValues={{email: '', password: ''}}
                    validationSchema={Yup.object({
                        email: Yup.string().email().required("Required"),
                        password: Yup.string().required("Required")
                    })}
                    onSubmit={(values, { setSubmitting }) => {
                        setSubmitting(true);
                        API_PROTECTED.post("/auths/login", values).then((response) => {
                            if (response.status === 200) {
                                localStorage.setItem("userinfo", true)
                                removeCookies()
                                setSubmitting(false);
                                navigate("/", { replace: true })
                            }
                        }).catch(reason => {
                            setSubmitting(false);
                            setFailed(true);
                            removeCookies();
                        })
                    }}
                >
                {( {values,
                    errors,
                    touched,
                    handleChange,
                    handleBlur,
                    handleSubmit,
                    isSubmitting }) => (
                        <>
                            <div className={failed ? "alert alert-danger" : "d-none"}>Login failed</div>
                            <Form
                                onChange={handleChange}
                                onBlur={handleBlur}
                                onSubmit={handleSubmit}
                            >
                                <div className="mb-1">
                                    <label htmlFor="email" className="form-label">Email</label>
                                    <Field name="email" type="text" value={values.email} className="form-control" />
                                    <div className={touched.email && errors.email ? "alert alert-danger" : null}><ErrorMessage name="email" /></div>
                                </div>
                                <div className="mb-1">
                                    <label htmlFor="password" className="form-label">Password</label>
                                    <Field name="password" type="password"  value={values.password} className="form-control" />
                                    <div className={touched.password && errors.password ? "alert alert-danger" : null}><ErrorMessage name="password" /></div>
                                </div>
                                <button type="submit" className="btn btn-lg btn-primary w-100" disabled={isSubmitting}>Sign in</button>
                            </Form>
                        </>
                )}
                </Formik>
                <hr />
                <div>
                    <form className="mb-1" action={api_url + "/auths/google/login"}><input type="hidden" value={location} name="redirect_url"/><input type="submit" className="btn btn-lg btn-secondary w-100" value="Google Login" /></form>
                    <form className="mb-1" action={api_url + "/auths/github/login"}><input type="hidden" value={location} name="redirect_url"/><input type="submit" className="btn btn-lg btn-secondary w-100" value="Github Login" /></form>
                    <form action={api_url + "/auths/gitlab/login"}><input type="hidden" value={location} name="redirect_url"/><input type="submit" className="btn btn-lg btn-secondary w-100" value="Gitlab Login" /></form>
                </div>
            </div>
        </div>
    )
}

export default Login