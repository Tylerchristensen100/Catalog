import { useDisclosure } from "@mantine/hooks";
import { Modal, Button } from "@mantine/core";
import UploadCSV from "../../../../global/components/UploadCSV";
import { useState, useRef } from "react";
import styles from "./CSV.module.css";
import ProgramRow from "./ProgramRow";
import { useData } from "../../../../global/contexts/DataContext.jsx";

const UploadProgramsCSV = () => {
  const [programs, setPrograms] = useState(null);
  const [opened, { open, close }] = useDisclosure(false);
  const modalRef = useRef({ open });
  const [errors, setErrors] = useState(null);
  const { schools, gradLevels, programActions } = useData();

  const validate = (program, i) => {
    let errors = [];
    if (!program.name) {
      errors.push("Name is required");
    }

    if (!program.program_type) {
      errors.push("Program Type is required");
    }

    if (!program.grad_level) {
      errors.push("Grad Level is required");
    }

    if (!program.school) {
      errors.push("School is required");
    }
    if (!program.major_code) {
      errors.push("Major Code is required");
    }
    if (!program.cip) {
      errors.push("CIP is required");
    }
    if (!program.description) {
      errors.push("Description is required");
    }
    if (program.online === undefined || program.online === null) {
      errors.push("Online is required");
    }

    if (Number.isNaN(Number(program.cip)) || Number(program.cip) < 1) {
      errors.push(`CIP must be a number greater than 0, got ${program.cip}`);
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

  const uploadProgram = async (program) => {
    await programActions.createProgram(program).then((res) => {
      if (res.status < 300 && res.status >= 200) {
        return true;
      }
    });
  };

  const handleCSV = (arr) => {
    const programs = arr.slice(1).map((row, i) => {
      const p = {
        name: row[0],
        program_type: row[2],
        grad_level: gradLevels.find(
          (l) => l.level.toLowerCase() === row[1].toLowerCase()
        ),
        school: schools.find(
          (s) => s.name.toLowerCase() === row[3].toLowerCase()
        ),
        major_code: row[4],
        cip: row[7],
        online:
          row[5].toLowerCase() == "true" ||
          row[5].toLowerCase() == "yes" ||
          row[5].toLowerCase() == "1",
        description: row[6],
        uploaded: null,
      };
      if (validate(p, i)) {
        return p;
      }
    });

    setPrograms(programs);
  };

  const submit = async () => {
    const p = [...programs];
    p.forEach(async (program) => {
      const success = await uploadProgram(program);
      program.uploaded = success ? "success" : "failure";
      setPrograms(p);
    });
  };
  return (
    <>
      <UploadCSV setData={handleCSV} />

      {programs && programs.length > 0 ? (
        <>
          <table>
            <thead>
              <tr>
                <th>Name</th>
                <th>Program Type</th>
                <th>Grad Level</th>
                <th>School</th>
                <th>Major Code</th>
                <th>CIP</th>
                <th>Online</th>
                <th>Description</th>
              </tr>
            </thead>
            <td>
              {programs.map((program, i) => (
                <ProgramRow
                  key={i}
                  program={program}
                  isChecked={program.online}
                  setChecked={() => (program.online = !program.online)}
                />
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
                <th>Program Type</th>
                <th>Grad Level</th>
                <th>School</th>
                <th>Major Code</th>
                <th>CIP</th>
                <th>Online</th>
                <th>Description</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td>Program Name</td>
                <td>Masters</td>
                <td>Graduate</td>
                <td>College of Engineering & Technology</td>
                <td>TEST</td>
                <td>11.0101</td>
                <td>This is a sample description for the program.</td>
                <td>
                  <input type="checkbox" defaultChecked={false} disabled />{" "}
                  Online
                </td>
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
export default UploadProgramsCSV;
