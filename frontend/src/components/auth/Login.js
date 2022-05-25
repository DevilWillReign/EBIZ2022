import { Formik, Field, Form, ErrorMessage } from "formik"
import * as Yup from "yup"
import { useHref, useLocation, useNavigate } from "react-router-dom"
import API from "../../util/api"
import './Login.css'

const Login = () => {
    const api_url = API.defaults.baseURL
    const location = useLocation()

    return (
        <div className="text-center">
            <div className="form-signin">
                <Formik
                    initialValues={{username: '', password: ''}}
                    validationSchema={Yup.object({
                        username: Yup.string().required("Required"),
                        password: Yup.string().required("Required")
                    })}
                    onSubmit={(values, { setSubmitting }) => {
                        setSubmitting(true);
                        API.post("/auths/login", JSON.stringify(values, null, 2)).then((response) => {
                            if (response.status === 200) {
                                localStorage.setItem("userinfo", JSON.stringify(response.data))
                                setSubmitting(false);
                            }
                        }).catch(reason => {
                            console.log(reason)
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
                        <Form
                            onChange={handleChange}
                            onBlur={handleBlur}
                            onSubmit={handleSubmit}
                        >
                            <div className="mb-1">
                                <label htmlFor="username" className="form-label">Username</label>
                                <Field name="username" type="text" value={values.username} className="form-control" />
                                <div className={touched.username && errors.username ? "alert alert-danger" : null}><ErrorMessage name="username" /></div>
                            </div>
                            <div className="mb-1">
                                <label htmlFor="password" className="form-label">Password</label>
                                <Field name="password" type="password"  value={values.password} className="form-control" />
                                <div className={touched.password && errors.password ? "alert alert-danger" : null}><ErrorMessage name="password" /></div>
                            </div>
                            <button type="submit" className="btn btn-lg btn-primary w-100" disabled={isSubmitting}>Sign in</button>
                        </Form>
                )}
                </Formik>
                <hr />
                <div>
                    <a className="btn btn-lg btn-secondary w-100 mb-1" href={api_url + "/auths/google/login?redirect_url=" + location.pathname} >Google Login</a>
                    <a className="btn btn-lg btn-secondary w-100" href={api_url + "/auths/github/login?redirect_url=" + location.pathname}>Github Login</a>
                </div>
            </div>
        </div>
    )
}

export default Login