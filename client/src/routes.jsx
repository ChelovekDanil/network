import { Routes, Route } from "react-router-dom";
import Registration from "./auth/registration";
import PublicRouter from "./auth/publicRouter";
import PrivateRouter from "./auth/privateRouter";
import Login from "./auth/login";
import Main from "./main/main";

function useRoutes() {
    return (
        <Routes>
            {/* <Route element={<PublicRouter/>}> */}
                <Route path='/' element={<Login/>}></Route>
                <Route path='/login' element={<Login />} />
                <Route path='/registration' element={<Registration />} />
            {/* </Route> */}
            {/* <Route element={<PrivateRouter/>}> */}
                {/* <Route path='/' element={<Main/>}></Route> */}
            {/* </Route>             */}
        </Routes>
    )
}

export default useRoutes