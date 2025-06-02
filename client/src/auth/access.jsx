import { useEffect } from "react";
import { useNavigate } from "react-router-dom";

const CheckToken = () => {    
    console.log("test");
    
    const navigate = useNavigate();

    useEffect(() => {
        const token = localStorage.getItem("accessToken");

        const checkAuth = async () => {
            try {
                const response = await fetch("http://localhost:8080/auth/check/", {
                    method: 'GET',
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${token}`
                    }
                });

                if (!response.ok) {                    
                    throw new Error(`Ошибка: ${response.status}`);
                }
            } catch (error) {
                try {
                    const response = await fetch("http://localhost:8080/auth/refresh/", {
                        method: 'POST',
                        headers: {
                            'Content-Type': 'application/json',
                        }
                    });
                    
                    if (!response.ok) {
                        throw new Error(`Ошибка: ${response.status}`);
                        
                    } 

                    const data = await response.json(); 
                    const accessToken = data.accessToken; 
                    const refreshToken = data.refreshToken;
                    localStorage.setItem("accessToken", accessToken)
                    localStorage.setItem("refreshToken", refreshToken)
                    console.log("uraaa");
                    
                } catch(error) {
                    console.error(error);
                    localStorage.setItem("isAuth", false);
                    navigate("/");    
                }
                console.error(error);
                localStorage.setItem("isAuth", false);
                navigate("/");
            }
        };

        checkAuth();
    }, [navigate]); 

    return null; 
};

export default CheckToken;
