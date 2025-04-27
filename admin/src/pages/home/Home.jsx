import { Grid, Skeleton } from "@mantine/core";
import { useData } from "../../global/contexts/DataContext";
import styles from "./Home.module.css";
const Home = () => {
  const { courses, programs, loading } = useData();

  return (
    <div>
      <h1>Catalog Admin</h1>
      <Grid className={styles.stats} gutter="md">
        <Grid.Col span={{ base: 12, md: 6 }}>
          <h3>Total Programs</h3>
          {loading ? (
            <Skeleton height={20} width="100%" />
          ) : (
            <span>{programs?.length ?? "NULL"}</span>
          )}
        </Grid.Col>

        <Grid.Col span={{ base: 12, md: 6 }}>
          <h3>Total Courses</h3>
          {loading ? (
            <Skeleton height={20} width="100%" />
          ) : (
            <span>{courses?.length ?? "NULL"}</span>
          )}
        </Grid.Col>
      </Grid>
    </div>
  );
};

export default Home;
