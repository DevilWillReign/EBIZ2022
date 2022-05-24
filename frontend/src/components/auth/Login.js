import { Formik, Field, Form, ErrorMessage } from "formik"
import * as Yup from "yup"
import './Login.css'

const Login = () => {
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
                        setTimeout(() => {
                        alert(JSON.stringify(values, null, 2));
                        setSubmitting(false);
                        }, 400);
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
                    <a className="btn btn-lg btn-secondary w-100 mb-1" href="http://localhost:9000/api/v1/auths/google/login?redirect_url=http://localhost:9001/login">Google Login</a>
                    <a className="btn btn-lg btn-secondary w-100" href="http://localhost:9000/api/v1/auths/github/login?redirect_url=http://localhost:9001/login">Github Login</a>
                </div>
            </div>
        </div>
    )
}

export default Login