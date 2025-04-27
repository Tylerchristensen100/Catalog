import { useState, useEffect } from "react";
import { setAuthToken } from "./axios_client";
import AuthContext from "./auth_context";
import PropTypes from "prop-types";
import { API } from "./axios_client";
import { getCookies, deleteAllCookies } from "./../helpers/CookieHelpers";

const cookieName = "access_token";
export const AuthProvider = ({ children }) => {
  const [auth, setAuth] = useState(false);
  const [token, setToken] = useState(null);
  const [user, setUserInfo] = useState(null);
  const [isLoading, setIsLoading] = useState(false);

  const login = () => {
    window.location.replace(window.location.origin + "/login");
  };

  const logout = async () => {
    await API.post("/api/logout", {
      access_token: token,
      logout: true,
    }).then((res) => res.data);

    deleteAllCookies();
    setAuth(false);
    setUserInfo(null);
    window.location.replace(window.location.origin + "/admin/");
  };

  const getAuthToken = () => {
    const cookies = getCookies();
    return cookies[cookieName];
  };
  const getUser = async () => {
    setIsLoading(true);

    const data = await API.get("/api/admin/user").then((res) => res.data);
    setUserInfo(data);
    setAuth(data != null ? true : false);
    setIsLoading(false);
  };

  useEffect(() => {
    const accessToken = getAuthToken();

    if (accessToken) {
      setAuthToken(accessToken);
      setToken(accessToken);
      getUser();
    } else {
      console.log("no accessToken");
    }
  }, []);

  useEffect(() => {
    cookieStore.addEventListener("change", ({ changed }) => {
      for (const { name, value } of changed) {
        if (name == "access_token") {
          if (!value || value.length < 10) {
            setAuth(false);
            setAuthToken(null);
            setUserInfo(null);
            return;
          }
          setAuthToken(value);
          setAuth(true);
        }
      }
    });
  }, []);

  const isAuthed = () => {
    const cookies = getCookies();

    const { value } = cookies[cookieName];
    if (!value || value.length < 10) {
      setAuth(false);
      setAuthToken(null);
      setUserInfo(null);
      return;
    } else {
      setAuthToken(value);
      setAuth(true);
    }

    return auth;
  };

  const exported = {
    auth: isAuthed,
    setAuth,
    user,
    login,
    logout,
    isLoading,
  };

  return (
    <AuthContext.Provider value={exported}>{children}</AuthContext.Provider>
  );
};

AuthProvider.propTypes = {
  children: PropTypes.node.isRequired,
};
