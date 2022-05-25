import { Formik, Field, Form, ErrorMessage } from "formik"
import { useNavigate } from "react-router-dom"
import * as Yup from "yup"
import API from "../../util/api"
import './Login.css'

const Register = () => {
    const navigate = useNavigate()

    return (
        <div className="text-center">
            <div className="form-signin">
                <Formik
                    initialValues={{username: '', email: '', password: '', repeat_password: ''}}
                    validationSchema={Yup.object({
                        username: Yup.string().required("Required"),
                        email: Yup.string().email().required("Required"),
                        password: Yup.string().required("Required"),
                        repeat_password: Yup.string().oneOf([Yup.ref("password"), null], "Password must match.")
                    })}
                    onSubmit={(values, { setSubmitting, resetForm }) => {
                        setSubmitting(true);
                        API.post("/users", values).then((response) => {
                            if (response.status === 201) {
                                localStorage.setItem("userinfo", JSON.stringify(response.data))
                                setSubmitting(false);
                                navigate("/", { replace: true });
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
                                <Field name="username" value={values.username} type="text" className="form-control" />
                                <div className={touched.username && errors.username ? "alert alert-danger" : null}><ErrorMessage name="username" className="valid-tooltip" /></div>
                            </div>
                            <div className="mb-1">
                                <label htmlFor="email" className="form-label">Email</label>
                                <Field name="email" value={values.email} type="email" className="form-control" />
                                <div className={touched.email && errors.email ? "alert alert-danger" : null}><ErrorMessage name="email" className="valid-tooltip" /></div>
                            </div>
                            <div className="mb-1">
                                <label htmlFor="password" className="form-label">Password</label>
                                <Field name="password" value={values.password} type="password" className="form-control" />
                                <div className={touched.password && errors.password ? "alert alert-danger" : null}><ErrorMessage name="password" className="valid-tooltip" /></div>
                            </div>
                            <div className="mb-1">
                                <label htmlFor="repeat_password" className="form-label">Repeat Password</label>
                                <Field name="repeat_password" value={values.repeat_password} type="password" className="form-control" />
                                <div className={touched.repeat_password && errors.repeat_password ? "alert alert-danger" : null}><ErrorMessage name="repeat_password" className="valid-tooltip" /></div>
                            </div>
                            <button type="submit" className="btn btn-lg btn-primary w-100" disabled={isSubmitting}>Sign up</button>
                        </Form>
                )}
                </Formik>
            </div>
        </div>
    )
}

export default Register