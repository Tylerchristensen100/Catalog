import { FileInput } from "@mantine/core";
import { notifications } from "@mantine/notifications";
import { faInfoCircle } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { convertCSVToArray } from "../../global/helpers/CSVHelpers";

const UploadCSV = ({ setData }) => {
  async function handleFileUpload(file) {
    if (file) {
      if (file.type !== "text/csv") {
        notifications.show({
          color: "red",
          title: "Invalid File Type",
          icon: <FontAwesomeIcon icon={faInfoCircle} />,
          message: `File must be a CSV not ${file.type}`,
        });
        file = null;
        return;
      }
      const text = await file.text();
      if (typeof text !== "string") return;

      const arr = convertCSVToArray(text);

      setData(arr);
    }
  }

  return (
    <>
      <FileInput
        label="Upload a CSV file"
        placeholder="No file selected"
        accept=".csv"
        onChange={handleFileUpload}
        required
        clearable
      />
    </>
  );
};

export default UploadCSV;
