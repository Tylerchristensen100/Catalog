import { useData } from "../../global/contexts/DataContext";
import { Skeleton, Accordion, Button, Flex, Autocomplete } from "@mantine/core";
import Program from "./components/Program";
import { useNavigate } from "react-router-dom";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faAdd } from "@fortawesome/free-solid-svg-icons";

const Programs = () => {
  const { programs } = useData();
  const navigate = useNavigate();

  const handleAutoCompleteChange = (value) => {
    if (value.length > 1) {
      const selectedProgram = programs.find(
        (program) => program.name === value || program.code === value
      );
      if (selectedProgram) {
        navigate(`/programs/${selectedProgram.name}`);
      }
    }
  };
  return (
    <div>
      <Flex justify="space-between" align="center">
        <h1>Programs</h1>
        <Flex justify="space-between" align="center" gap="md">
          <Button
            onClick={() => navigate(`/programs/upload`)}
            variant="outline"
          >
            Upload CSV
          </Button>

          <Button
            onClick={() => navigate(`/programs/create`)}
            variant="outline"
          >
            New <FontAwesomeIcon icon={faAdd} />
          </Button>
        </Flex>
      </Flex>
      <Autocomplete
        placeholder="Search Programs"
        clearable
        data={programs.map((p) => `${p?.name}`)}
        onChange={handleAutoCompleteChange}
        limit={10}
        my={50}
      />
      <ul>
        <Accordion
          variant="filled"
          radius="lg"
          chevronPosition="left"
          disableChevronRotation
        >
          {programs.length > 0 ? (
            programs.map((program) => <Program key={program?.id} {...program} />)
          ) : (
            <Skeleton height={1000} radius="md" animate={true} />
          )}
        </Accordion>
      </ul>
    </div>
  );
};

export default Programs;
