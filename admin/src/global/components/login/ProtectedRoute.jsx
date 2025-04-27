import { useNavigate } from "react-router-dom";
import { useAuth } from "../../auth/use_auth";
import { Skeleton } from '@mantine/core';
import { useEffect } from "react";
import PropTypes from 'prop-types';


const ProtectedRoute = ( { children } ) => {
  const navigate = useNavigate();

  const { auth, isLoading } = useAuth();

  useEffect(() => {
    if(!auth && !isLoading) {
      navigate("/login");
    }
  }, [auth, isLoading, navigate])
 

  if (isLoading) {
    return <Skeleton />;
  }
  return <>{children}</>;
}

ProtectedRoute.propTypes = {
  children: PropTypes.node.isRequired,
};

export default ProtectedRoute;