import { useParams } from "react-router-dom";
import { useData } from "../../global/contexts/DataContext.jsx";
import { TextInput, Textarea, Button, Grid, Center } from "@mantine/core";
import { notifications } from "@mantine/notifications";
import { useState } from "react";
import { useNavigate } from "react-router-dom";

const EditCourse = () => {
  const { id } = useParams();
  const navigator = useNavigate();

  const { courses, courseActions } = useData();
  const course = courses.find(
    (course) => `${course.major_code}-${course.code}` === id
  );

  const isCreating = id == null || course == null;

  const [name, setName] = useState(course?.name ?? "");
  const [majorCode, setMajorCode] = useState(course?.major_code ?? "");
  const [code, setCode] = useState(course?.code);
  const [creditHours, setCreditHours] = useState(course?.credit_hours);
  const [prerequisites, setPrereqs] = useState(course?.prerequisites ?? "");
  const [description, setDescription] = useState(course?.description ?? "");

  async function submit(e) {
    e.preventDefault();

    console.log({
      id,
      name,
      majorCode,
      code,
      creditHours,
      prerequisites,
      description,
    });
    if (!name || !majorCode || !code || !creditHours || !description) {
      notifications.show({
        title: "Missing Information",
        message: "Please fill out all fields",
        color: "red",
      });
      return;
    }

    const c = new FormData();
    c.append("name", name);
    c.append("major_code", majorCode);
    c.append("code", Number(code));
    c.append("credit_hours", Number(creditHours));
    c.append("prerequisites", prerequisites);
    c.append("description", description);

    let isSuccessful;
    if (isCreating) {
      isSuccessful = await courseActions.createCourse(c);
    } else {
      c.id = course.id;
      isSuccessful = await courseActions.updateCourse(c);
    }

    if (isSuccessful) {
      notifications.show({
        title: isCreating ? "Course Created" : "Course Updated",
        message: `${name} Saved!`,
      });
      navigator("/");
    } else {
      notifications.show({
        title: isCreating ? "Course Creation Failed" : "Course Update Failed",
        message: `${name} Not Saved!`,
        color: "red",
      });
    }
  }

  return (
    <>
      <h1>{id ? `Edit ${id}` : "Create Course"}</h1>

      <form onSubmit={submit}>
        <TextInput
          label="Name"
          description="The Name of the Course"
          placeholder=""
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
          <Grid.Col span={4}>
            <TextInput
              label="Major Code"
              description="The code for the major"
              placeholder="CS"
              maw={250}
              value={majorCode}
              onChange={(e) => setMajorCode(e.currentTarget.value)}
              miw={100}
              required
              max-length="3"
              name="major_code"
            />
          </Grid.Col>
          <Grid.Col span={3}>
            <TextInput
              label="Code"
              description="The code for the course"
              placeholder="1400"
              maw={250}
              value={code}
              onChange={(e) => setCode(e.currentTarget.value)}
              miw={100}
              required
              type="number"
              max-length="4"
              name="code"
            />
          </Grid.Col>
          <Grid.Col span={3}>
            <TextInput
              label="Credit Hours"
              description="How many credit hours the course is worth"
              placeholder="3"
              maw={250}
              value={creditHours}
              onChange={(e) => setCreditHours(e.currentTarget.value)}
              miw={100}
              required
              type="number"
              max-length="1"
              name="credit_hours"
            />
          </Grid.Col>
        </Grid>
        <TextInput
          label="Prerequisites"
          description="The prerequisites for the course"
          placeholder=""
          value={prerequisites}
          onChange={(e) => setPrereqs(e.currentTarget.value)}
          name="prerequisites"
        />
        <Textarea
          label="Description"
          value={description}
          onChange={(e) => setDescription(e.currentTarget.value)}
          autosize
          minRows={2}
          resize="vertical"
          required
          name="description"
          my={50}
        />

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
