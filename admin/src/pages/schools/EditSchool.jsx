import { useParams } from "react-router-dom";
import { useData } from "../../global/contexts/DataContext.jsx";
import { TextInput, Button, Grid, Center } from "@mantine/core";
import { notifications } from "@mantine/notifications";
import { useState } from "react";

const EditCourse = () => {
  const { schoolCode } = useParams();

  const { schools, createSchool, updateSchool } = useData();
  const school = schools.find((s) => s.code === schoolCode);

  const isCreating = schoolCode == null || school == null;

  const [name, setName] = useState(school?.name ?? "");
  const [code, setCode] = useState(school?.code);

  async function submit(e) {
    e.preventDefault();
    if (!name || !code) {
      notifications.show({
        title: "Missing Information",
        message: "Please fill out all fields",
        color: "red",
      });
      return;
    }

    const s = new FormData();
    s.name = name;
    s.code = code;

    if (isCreating) {
      await createSchool(s);
    } else {
      s.id = school.id;
      await updateSchool(s);
    }

    notifications.show({
      title: isCreating ? "School Created" : name + " Updated",
      message: `${name} Saved!`,
    });
  }

  return (
    <>
      <h1>{school ? `Edit ${school.name}` : "Create School"}</h1>

      <form onSubmit={submit}>
        <TextInput
          label="Name"
          description="The Name of the School"
          placeholder="The College of something educational"
          value={name}
          onChange={(e) => setName(e.currentTarget.value)}
          required
          name="name"
        />

        <Grid
          grow
          justify="center"
          align="flex-end"
          my={50}
          type="container"
          breakpoints={{
            xs: "100px",
            sm: "200px",
            md: "300px",
            lg: "400px",
            xl: "500px",
          }}
        >
          <Grid.Col span={3}>
            <TextInput
              label="School Code"
              description="The acronym for the school"
              placeholder="CHSS"
              maw={350}
              value={code}
              onChange={(e) => setCode(e.currentTarget.value)}
              miw={100}
              required
              max-length="4"
              name="code"
            />
          </Grid.Col>
        </Grid>

        <Center mt={50}>
          <Button type="submit" mt="sm">
            Submit
          </Button>
        </Center>
      </form>
    </>
  );
};

export default EditCourse;
