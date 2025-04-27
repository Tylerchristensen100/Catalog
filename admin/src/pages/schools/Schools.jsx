import { useData } from "../../global/contexts/DataContext";
import School from "./components/School";
import { Skeleton, Button, Accordion, Flex } from "@mantine/core";
import { useNavigate } from "react-router-dom";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faAdd } from "@fortawesome/free-solid-svg-icons";

const Schools = () => {
  const { schools } = useData();
  const navigate = useNavigate();

  return (
    <div>

      <Flex justify="space-between" align="center">
      <h1>Schools</h1>

      <Button onClick={() => navigate(`/schools/create`)} variant="outline" color="blue">
        New <FontAwesomeIcon icon={faAdd} />
      </Button>
            </Flex>
     
      <ul>
        <Accordion
          variant="filled"
          radius="lg"
          chevronPosition="left"
          disableChevronRotation
          defaultValue="Apples"
        >
          {schools.length > 0 ? (
            schools.map((school) => <School key={school.id} {...school} />)
          ) : (
            <Skeleton height={1000} radius="md" animate={true} />
          )}
        </Accordion>
      </ul>
    </div>
  );
};

export default Schools;
