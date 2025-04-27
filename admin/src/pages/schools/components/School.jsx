import { useNavigate } from "react-router-dom";
import AccordionControl from "../../../global/components/accordion_control/AccordionControl";
import { faEdit } from "@fortawesome/free-solid-svg-icons";
import { Accordion } from "@mantine/core";

const School = ({ id, name, code }) => {
  const navigate = useNavigate();
  return (
    <li>
      <Accordion.Item key={id} value={name}>
        <AccordionControl
          icon={faEdit}
          onClick={() => navigate(`/schools/${code}`)}
        >
          {name}
        </AccordionControl>

        <Accordion.Panel>
          <p>
            <strong>College Code</strong> {code}
          </p>
        </Accordion.Panel>
      </Accordion.Item>
    </li>
  );
};

export default School;
