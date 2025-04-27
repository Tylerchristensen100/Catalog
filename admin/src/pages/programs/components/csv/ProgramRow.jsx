const ProgramRow = ({ id, program, isChecked, setChecked }) => {
  return (
    <tr key={id} className={program?.uploaded != null ? program.uploaded : ""}>
      <td>{program?.name}</td>
      <td>{program?.description}</td>
      <td> {program.program_type}</td>
      <td>{program?.school}</td>
      <td>{program.grad_level}</td>
      <td>{program?.cip}</td>
      <td>
        <input
          type="checkbox"
          checked={isChecked}
          onChange={(e) => setChecked(e.target.checked)}
        />
        Online
      </td>
    </tr>
  );
};

export default ProgramRow;
