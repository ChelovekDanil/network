function CheckAuth() {
    return localStorage.getItem("isAuth") === "true" ? true : false
}

export default CheckAuth