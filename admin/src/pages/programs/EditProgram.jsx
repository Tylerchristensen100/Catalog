import { useParams } from "react-router-dom";
import { useData } from "../../global/contexts/DataContext.jsx";
import {
  TextInput,
  Textarea,
  Switch,
  Select,
  Button,
  Grid,
  Center,
} from "@mantine/core";
import { useForm } from "@mantine/form";
import { notifications } from "@mantine/notifications";
import { useNavigate } from "react-router-dom";
import { useState } from "react";

const EditProgram = () => {
  const { programName } = useParams();
  const { programs, schools, gradLevels, programActions } = useData();
  const navigator = useNavigate();

  const program = programs?.find((program) => program?.name === programName);

  const isCreating = programName == null || program == null;

  const [name, setName] = useState(program?.name ?? "");
  const [gradLevel, setGradLevel] = useState(program?.grad_level ?? {});

  const [school, setSchool] = useState(program?.school ?? {});

  const [majorCode, setMajorCode] = useState(program?.major_code ?? "");
  const [online, setOnline] = useState(program?.online ?? false);
  const [description, setDescription] = useState(
    program?.description.trim() ?? ""
  );
  const [cip, setCip] = useState(program?.cip ?? 0);

  const form = useForm({
    initialValues: {
      name: program?.name ?? "",
      gradLevel: program?.grad_level ?? {},
      school: program?.school ?? {},
      majorCode: program?.major_code ?? "",
      online: program?.online ?? false,
      description: program?.description.trim() ?? "",
      cip: program?.cip ?? 0,
    },
    validationRules: {
      name: (value) => value.trim().length > 0,
      gradLevel: (value) => value.level.trim().length > 0,
      school: (value) => value.code.trim().length > 0,
      majorCode: (value) => value.trim().length > 0,
      online: (value) => value,
      description: (value) => value.trim().length > 0,
      cip: (value) => value.trim().length > 0 && !isNaN(value),
    },
    validationMessages: {
      name: "Name is required",
      gradLevel: "Grad Level is required",
      school: "School is required",
      majorCode: "Major Code is required",
      online: "Online is required",
      description: "Description is required",
      cip: "CIP is required",
    },
    onSubmit: submit,
    validateInputOnBlur: true,
  });

  async function submit(e) {
    e.preventDefault();

    const p = new FormData();
    p.append("name", name);
    p.append("grad_level", gradLevels.find((g) => g.level === gradLevel.value).id);
    console.log(gradLevel)
    p.append(
      "school",
      school?.value
        ? schools.find((s) => s.code === school.value).id
        : program.school.id
    );
    p.append("major_code", majorCode);
    p.append("online", online ? 1: 0);
    p.append("description", description);
    p.append("cip", cip);

    p.append(
      "program_type",
      gradLevel?.value == "Graduate Certificate" || gradLevel?.value == "Master"
        ? "Graduate"
        : "Undergraduate"
    );

    p.append("campus", 1);



    let isSuccessful;
    if (isCreating) {
      isSuccessful = await programActions.createProgram(p);
    } else {
      p.append("id", program.id);
      isSuccessful = await programActions.updateProgram(program.id, p);
    }

    if (isSuccessful) {
      notifications.show({
        title: isCreating ? "Program Created" : "Program Updated",
        message: `${name} Saved!`,
      });
      navigator("/");
    } else {
      notifications.show({
        title: isCreating ? "Program Creation Failed" : "Program Update Failed",
        message: `${name} Not Saved!`,
        color: "red",
      });
    }
  }

  return (
    <>
      <h1>{name ? `Edit ${name}` : "Create Program"}</h1>

      <form onSubmit={submit}>
        <Grid
          grow
          gutter={{ base: 5, xs: "md", md: "xl", xl: 50 }}
          justify="center"
          align="flex-end"
          type="container"
          breakpoints={{
            xs: "100px",
            sm: "200px",
            md: "300px",
            lg: "400px",
            xl: "500px",
          }}
        >
          <Grid.Col span={8}>
            <TextInput
              key={form.key("name")}
              label="Name"
              description="The Name of the program"
              placeholder=""
              value={name}
              onChange={(e) => setName(e.currentTarget.value)}
              required
              name="name"
            />
          </Grid.Col>
          <Grid.Col span={4}>
            <Switch
              key={form.key("online")}
              label="Online"
              checked={online}
              onChange={(e) => setOnline(e.currentTarget.checked)}
              onLabel="Yes"
              offLabel="No"
              size="lg"
              my={5}
              name="online"
            />
          </Grid.Col>
        </Grid>

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
            <Select
              key={form.key("school")}
              label="School"
              description="The school the program belongs to"
              placeholder="Select School"
              data={schools.map((school) => ({
                value: school.code,
                label: school.name,
              }))}
              value={school ? school.code : ""}
              onChange={(_value, option) => setSchool(option)}
              miw={450}
              required
              name="school"
            />
          </Grid.Col>
          <Grid.Col span={2}>
            <Select
              key={form.key("gradLevel")}
              label="Program Level"
              description="The level of the program"
              placeholder="Select Program Value"
              maw={250}
              data={gradLevels.map((gradLevel) => ({
                value: gradLevel.level,
                label: gradLevel.level,
              }))}
              value={gradLevel ? gradLevel.level : ""}
              onChange={(_value, option) => setGradLevel(option)}
              miw={200}
              required
              name="program_level"
            />
          </Grid.Col>
          <Grid.Col span={2.5}>
            <TextInput
              key={form.key("cip")}
              label="CIP Code"
              description="Classification of Instructional Programs Code"
              placeholder="12.0106"
              maw={250}
              value={cip}
              onChange={(e) => setCip(e.currentTarget.value)}
              miw={100}
              required
              type="number"
              max-length="7"
              name="cip"
            />
          </Grid.Col>
          <Grid.Col span={2.5}>
            <TextInput
              key={form.key("majorCode")}
              label="Major Code"
              description="The 2 letter prefix for the major"
              placeholder="CS"
              value={majorCode}
              onChange={(e) => setMajorCode(e.currentTarget.value)}
              maw={250}
              miw={50}
              required
              name="major_code"
            />
          </Grid.Col>
        </Grid>

        <Textarea
          key={form.key("description")}
          label="Description"
          description=""
          placeholder=""
          value={description}
          onChange={(e) => setDescription(e.currentTarget.value)}
          autosize
          minRows={2}
          resize="vertical"
          required
          name="description"
        />

        {/* <TextInput
        label="Campus"
        description=""
        placeholder=""
        value={campus}
        onChange={(e) => setCampus(e.currentTarget.value)}
      /> */}

        <Center mt={50}>
          <Button type="submit" mt="sm">
            Submit
          </Button>
        </Center>
      </form>
    </>
  );
};

export default EditProgram;
