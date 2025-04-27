import { useEffect } from "react";
import { useAuth } from "../../auth/use_auth";


const Logout = () => {
    const { logout } = useAuth();

    useEffect(() => {
        logout();
    }, [logout]);

    return (
        <div>
            Logged Out
        </div>
    );
}

export default Logout;