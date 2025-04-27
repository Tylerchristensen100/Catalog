import styles from "./program.module.css";
import { Accordion, Grid } from "@mantine/core";
import AccordionControl from "../../../global/components/accordion_control/AccordionControl";
import { faEdit } from "@fortawesome/free-solid-svg-icons";
import { useNavigate } from "react-router-dom";

function Program({ id, name, description, program_type, school, grad_level, cip, online }) {
  const navigate = useNavigate();
  return (
    <li className={styles.program}>
      <Accordion.Item key={id} value={name} className={styles.accordion}>
        <AccordionControl
          icon={faEdit}
          onClick={() => navigate(`/programs/${name}`)}
        >
          {name}
        </AccordionControl>

        <Accordion.Panel>
          <p dangerouslySetInnerHTML={{ __html: description }}></p>
          <Grid align="center" justify="space-between">
            <Grid.Col span={6}>
            <p><strong>Program Type</strong> {program_type}</p>
            </Grid.Col>
            <Grid.Col span={3}>
              <p><strong>Program CIP</strong> {cip}</p>
            </Grid.Col>
            <Grid.Col span={3}>
              <p><strong>Online</strong> {online > 0 ? "Yes" : "No"}</p>
            </Grid.Col>


            <Grid.Col span={6}>
            <p><strong>School</strong> {school?.name}</p>
            </Grid.Col>
            <Grid.Col span={6}>
              <p><strong>Grad Level</strong> {grad_level?.level}</p>
            </Grid.Col>
          </Grid>
          
        </Accordion.Panel>
      </Accordion.Item>
    </li>
  );
}

export default Program;
