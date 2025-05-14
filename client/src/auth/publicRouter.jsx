import { Navigate, Outlet, useLocation } from 'react-router-dom';
import CheckAuth from './scripts'

function PublicRouter() {
    const location = useLocation()

    return (
        CheckAuth() ? <Navigate to="/" state={{from: location}} replace /> : <Outlet/>
    )
}

export default PublicRouter;