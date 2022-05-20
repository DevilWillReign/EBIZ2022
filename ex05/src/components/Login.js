const Login = () => {
    return (
        <div>
            <a href="http://localhost:9000/api/v1/auths/google/login?redirect_url=http://localhost:9001/login">Google Login</a>
            <a href="http://localhost:9000/api/v1/auths/github/login?redirect_url=http://localhost:9001/login">Github Login</a>
        </div>
    )
}

export default Login