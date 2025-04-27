const CourseRow = ({ id, course }) => {
  return (
    <tr key={id} className={course?.uploaded != null ? course.uploaded : ""}>
      <td>{course?.name}</td>
      <td>{course?.majorCode}</td>
      <td>{course?.major}</td>
      <td>{course?.code}</td>
      <td>{course?.credit_hours}</td>
      <td>{course?.description}</td>
    </tr>
  );
};

export default CourseRow;
