import { useEffect } from "react";
import useRoutes from './routes.jsx'
import { useLocation, useNavigate } from "react-router-dom";

function App() {
  const routes = useRoutes()
  const location = useLocation();
  // const navigate = useNavigate();


  useEffect(() => {    
    // const token = localStorage.getItem("accessToken");
  }, [location]);

  return (
    <>
      {routes}      
    </>
  )
}

export default App
