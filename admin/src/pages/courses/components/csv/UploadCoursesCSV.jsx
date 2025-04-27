import { Button, Modal } from "@mantine/core";
import UploadCSV from "../../../../global/components/UploadCSV";
import { useState, useRef } from "react";
import styles from "./CSV.module.css";
import CourseRow from "./CourseRow";
import { useData } from "../../../../global/contexts/DataContext.jsx";
import { useDisclosure } from "@mantine/hooks";

const UploadCoursesCSV = () => {
  const [courses, setCourses] = useState(null);
  const [opened, { open, close }] = useDisclosure(false);
  const modalRef = useRef({ open });
  const [errors, setErrors] = useState(null);
  const { courseActions } = useData();

  const validate = (course, i) => {
    let errors = [];
    if (!course.name) {
      errors.push("Name is required");
    }

    if (errors.length > 0) {
      modalRef.current.open();
      setErrors({
        row: i + 1,
        errors: errors,
      });
      return false;
    }

    return true;
  };

  const handleCSV = (arr) => {
    const courses = arr.slice(1).map((row, i) => {
      const c = {
        name: row[0],
        major_code: row[1],
        code: row[2],
        credit_hours: row[3],
        prerequisites: row[4],
        description: row[5],
        uploaded: null,
      };
      if (validate(c, i)) {
        return c;
      }
    });

    setCourses(courses);
  };

  const uploadCourse = async (course) => {
    await courseActions.createCourse(course).then((res) => {
      if (res.status < 300 && res.status >= 200) {
        return true;
      }
    });
  };

  const submit = async () => {
    const c = [...courses];
    c.forEach(async (course) => {
      const success = await uploadCourse(course);
      course.uploaded = success ? "success" : "failure";
      setCourses(c);
    });
  };
  return (
    <>
      <UploadCSV setData={handleCSV} />

      {courses && courses.length > 0 ? (
        <>
          <table>
            <thead>
              <tr>
                <th>Name</th>
                <th>Major Code</th>
                <th>Code</th>
                <th>Credit Hours</th>
                <th>Prerequisites</th>
                <th>Description</th>
              </tr>
            </thead>
            <td>
              {courses.map((course, i) => (
                <CourseRow key={i} course={course} />
              ))}
            </td>
          </table>

          <Button onclick={submit}>Submit</Button>
        </>
      ) : (
        <>
          <h2>Example CSV</h2>
          <table className={styles.table}>
            <thead>
              <tr>
                <th>Name</th>
                <th>Major Code</th>
                <th>Code</th>
                <th>Credit Hours</th>
                <th>Prerequisites</th>
                <th>Description</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td>Course Name</td>
                <td>TEST</td>
                <td>1055</td>
                <td>3</td>
                <td>None</td>
                <td>This is a sample description for the Course.</td>
              </tr>
            </tbody>
          </table>
        </>
      )}

      <Modal opened={opened} onClose={close} title="CSV Errors">
        {errors ? (
          <>
            <p>Errors found on row {errors.row}</p>
            <ul>
              {errors.errors.map((error, i) => (
                <li key={i}>{error}</li>
              ))}
            </ul>
          </>
        ) : (
          <p>No errors found</p>
        )}
      </Modal>
    </>
  );
};
export default UploadCoursesCSV;
