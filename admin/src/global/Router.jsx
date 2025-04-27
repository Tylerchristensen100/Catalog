import { Route, Routes } from "react-router-dom";
import Home from "../pages/home/Home";
import Login from "./components/login/Login";
import Logout from "./components/login/Logout";
import Callback from "./components/login/Callback";
import ProtectedRoute from "./components/login/ProtectedRoute";
import PageNotFound from "../pages/PageNotFound";
import Programs from "../pages/programs/Programs";
import Courses from "../pages/courses/Courses";
import EditCourse from "../pages/courses/EditCourse";
import EditProgram from "../pages/programs/EditProgram";
import Schools from "../pages/schools/Schools";
import EditSchool from "../pages/schools/EditSchool";
import UploadProgramsCSV from "../pages/programs/components/csv/UploadProgramsCSV";
import UploadCoursesCSV from "../pages/courses/components/csv/UploadCoursesCSV";


const Router = () => {
  return (
    <Routes>
      <Route path="/" element={<Home />} handle={{ crumb: () => "Home" }} />
      <Route
        path="/login"
        element={<Login />}
        handle={{ crumb: () => "Login" }}
      />
      <Route
        path="/logout"
        element={<Logout />}
        handle={{ crumb: () => "Logout" }}
      />
      <Route
        path="/callback"
        element={<Callback />}
        handle={{ crumb: () => "Callback" }}
      />
      <Route
        path="/programs"
        element={
          <ProtectedRoute>
            <Programs />
          </ProtectedRoute>
        }
        handle={{ crumb: () => "Programs" }}
      />
      <Route
        path="/programs/:programName"
        element={
          <ProtectedRoute>
            <EditProgram />
          </ProtectedRoute>
        }
        handle={{ crumb: () => "Edit Program" }}
      />

      <Route
        path="/programs/new"
        element={
          <ProtectedRoute>
            <EditProgram />
          </ProtectedRoute>
        }
        handle={{ crumb: () => "Create Program" }}
      />

<Route
        path="/programs/upload"
        element={
          <ProtectedRoute>
            <UploadProgramsCSV />
          </ProtectedRoute>
        }
        handle={{ crumb: () => "Upload Programs" }}
      />

      <Route
        path="/courses"
        element={
          <ProtectedRoute>
            <Courses />
          </ProtectedRoute>
        }
        handle={{ crumb: () => "Courses" }}
      />

      <Route
        path="/courses/:id"
        element={
          <ProtectedRoute>
            <EditCourse />
          </ProtectedRoute>
        }
        handle={{ crumb: () => "Edit Course" }}
      />

      <Route
        path="/courses/new"
        element={
          <ProtectedRoute>
            <EditCourse />
          </ProtectedRoute>
        }
        handle={{ crumb: () => "Create Course" }}
      />

<Route
        path="/courses/upload"
        element={
          <ProtectedRoute>
            <UploadCoursesCSV />
          </ProtectedRoute>
        }
        handle={{ crumb: () => "Upload Courses" }}
      />




<Route
        path="/schools"
        element={
          <ProtectedRoute>
            <Schools />
          </ProtectedRoute>
        }
        handle={{ crumb: () => "Schools" }}
      />

      <Route
        path="/schools/:schoolCode"
        element={
          <ProtectedRoute>
            <EditSchool />
          </ProtectedRoute>
        }
        handle={{ crumb: () => "Edit School" }}
      />

      <Route
        path="/schools/new"
        element={
          <ProtectedRoute>
            <EditSchool />
          </ProtectedRoute>
        }
        handle={{ crumb: () => "Create School" }}
      />

      <Route path="*" element={<PageNotFound />} />
    </Routes>
  );
};

export default Router;
