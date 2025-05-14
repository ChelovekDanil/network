import { Navigate, Outlet, useLocation } from 'react-router-dom';
import CheckAuth from './scripts'

function PrivateRouter() {
    const location = useLocation()

    return (
        CheckAuth() ? <Outlet/> : <Navigate to="/login" state={{from: location}} replace/>
    )
}

export default PrivateRouter;