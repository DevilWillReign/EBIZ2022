import { Formik, Field, Form, ErrorMessage } from "formik"
import { useEffect, useState } from "react"
import * as Yup from "yup"
import { useNavigate, useParams } from "react-router-dom"
import { API_PROTECTED } from "../../util/api"

const PayForm = () => {
    const navigate = useNavigate()
    const [loggedIn, ] = useState(localStorage.getItem("userinfo") !== null)
    const [failed, setFailed] = useState(false)
    const { transactionId } = useParams()
    const [total, setTotal] = useState(0)

    useEffect(() => {
        if (!loggedIn) {
            localStorage.setItem("userinfo", null)
            navigate("/auth/logout", { replace: true })
        }
        API_PROTECTED.get("/user/transactions/" + transactionId + "/total").then(response => {
            console.log(response)
            setTotal(response.data.total)
        }).catch(() => {})
    }, [loggedIn, navigate])

    return (
        <div className="text-center">
            <div className="form-signin">
                <Formik
                    initialValues={{cardnumber: ''}}
                    validationSchema={Yup.object({
                        cardnumber: Yup.string().required("Required"),
                    })}
                    onSubmit={(values, { setSubmitting, resetForm }) => {
                        setSubmitting(true);
                        API_PROTECTED.post("/user/payments", {paymenttype: 1, transactionid: Number(transactionId)}).then((response) => {
                            if (response.status === 201) {
                                resetForm()
                                setSubmitting(false);
                                navigate("/profile/transactions", { replace: true })
                            }
                        }).catch(() => {
                            setSubmitting(false);
                            setFailed(true);
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
                        <div className={failed ? "alert alert-danger" : "d-none"}>Failed to complete payment</div>
                            <Form
                                onChange={handleChange}
                                onBlur={handleBlur}
                                onSubmit={handleSubmit}
                            >
                                <div className="mb-1">
                                    <span className="fw-bold fs-2">Total: {total}</span>
                                </div>
                                <div className="mb-1">
                                    <label htmlFor="cardnumber" className="form-label">Card Number</label>
                                    <Field name="cardnumber" type="text" value={values.cardnumber} className="form-control" />
                                    <div className={touched.cardnumber && errors.cardnumber ? "alert alert-danger" : null}><ErrorMessage name="cardnumber" /></div>
                                </div>
                                <button type="submit" className="btn btn-lg btn-primary w-100" disabled={isSubmitting}>Pay</button>
                            </Form>
                        </>
                )}
                </Formik>
            </div>
        </div>
    )
}

export default PayForm