import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { useAuth } from "../../auth/use_auth";
import { showNotification } from "@mantine/notifications";
import { Skeleton } from '@mantine/core';



const Callback = () => {
  const navigate = useNavigate();
  const { auth, userManager, user, setUser, logout } = useAuth();

  useEffect(() => {
    if (auth === false) {
      userManager
        .signinRedirectCallback()
        .then((newUser) => {
          if (newUser) {
            showNotification({
              title: "Welcome",
              message: `Welcome ${newUser.profile.name}`,
              color: "blue",
            });

            setUser({
              name: newUser.profile.name,
              email: newUser.profile.email,
              username: newUser.profile.username,
            });
            navigate("/")
          }
        })
        .catch((err) => {
          console.error(err);
        });
    }
  }, [auth, userManager, user, navigate, setUser]);
  if (auth === true && user) {
    return (
      <div className="user">
        <h2>Welcome, {user.name}!</h2>
        <p className="description">Your ZITADEL Profile Information</p>
        <p>Name: {user.name}</p>
        <p>Email: {user.email}</p>
        <p>Email Verified: {user.email ? "Yes" : "No"}</p>
        <p>Roles: </p>

        <button onClick={logout}>Log out</button>
      </div>
    );
  } else {
    return <Skeleton />;
  }
};

export default Callback;