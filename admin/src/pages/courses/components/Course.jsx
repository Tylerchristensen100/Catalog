import React from "react";
import styles from "./course.module.css";
import { useNavigate } from "react-router-dom";
import AccordionControl from "../../../global/components/accordion_control/AccordionControl";
import { faEdit } from "@fortawesome/free-solid-svg-icons";
import { Accordion } from "@mantine/core";

const Course = React.memo(({
  id,
  name,
  major_code,
  major,
  code,
  credit_hours,
  description,
}) => {
  const navigate = useNavigate();
  return (
    <li className={styles.course}>
      <Accordion.Item key={id} value={name.length > 0 ? name : `${major_code}-${code}`} className={styles.accordion}>
        <AccordionControl
          icon={faEdit}
          onClick={() => navigate(`/courses/${major_code}-${code}`)}
        >
          {name.length > 0 ? name : `${major_code}-${code}`}
        </AccordionControl>

        <Accordion.Panel>
          <p>
            <strong>Course Code</strong> {major_code}-{code}
          </p>
          <p>
            <strong>Major</strong> {major.name}
          </p>

          <p>
            <strong>Credit Hours</strong> {credit_hours}
          </p>
          <p>{description}</p>
        </Accordion.Panel>
      </Accordion.Item>
    </li>
  );
});
Course.displayName = "Course";
export default Course;
