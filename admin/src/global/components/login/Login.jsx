import { useNavigate } from "react-router-dom";
import { useAuth } from "../../auth/use_auth";
import { useEffect } from "react";

const Login = () => {
  const navigate = useNavigate();
  const { auth, login } = useAuth();

  useEffect(() => {
    if (auth) {
      navigate("/programs");
    } else {
      login();
    }
  }, [navigate, auth, login]);


  return (
    <div>
      <h1>Login</h1>
    </div>
  );
};

export default Login;