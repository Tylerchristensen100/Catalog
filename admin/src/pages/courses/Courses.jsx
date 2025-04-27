import { useData } from "../../global/contexts/DataContext";
import Course from "./components/Course";
import { Skeleton, Button, Accordion, Flex, Autocomplete } from "@mantine/core";
import { useNavigate } from "react-router-dom";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faAdd } from "@fortawesome/free-solid-svg-icons";
import { useState, useEffect } from "react";

const CHUNK_SIZE = 500;

const Courses = () => {
  const { courses } = useData();
  const navigate = useNavigate();
  const [renderedChunks, setRenderedChunks] = useState(1);
  const totalChunks = Math.ceil(courses.length / CHUNK_SIZE);

  useEffect(() => {
    function isScrolledOverThreshold() {
      const THRESHOLD = 0.75;
      const scrollHeight = Math.max(
        document.documentElement.scrollHeight,
        document.body.scrollHeight
      );
      const viewportHeight = window.innerHeight;
      const totalScrollableHeight = scrollHeight - viewportHeight;

      const eightyPercentThreshold = totalScrollableHeight * THRESHOLD;
      const currentScrollY =
        window.scrollY || document.documentElement.scrollTop;
      return currentScrollY >= eightyPercentThreshold;
    }

    document.addEventListener("scroll", () => {
      if (isScrolledOverThreshold() && renderedChunks < totalChunks) {
        setRenderedChunks(renderedChunks + 1);
      }
    });
  }, [renderedChunks, totalChunks]);

  const renderCourses = () => {
    let renderedItems = [];
    for (let i = 0; i < renderedChunks; i++) {
      const chunk = courses.slice(i * CHUNK_SIZE, (i + 1) * CHUNK_SIZE);
      renderedItems = renderedItems.concat(
        chunk.map((course) => <Course key={course.id} {...course} />)
      );
    }
    return renderedItems;
  };

  const handleAutoCompleteChange = (value) => {
    if (value.length > 2) {
      const [majorCode, courseCode] = value.split(" ")[0].split("-");
      const selectedCourse = courses.find(
        (course) =>
          course.major_code === majorCode && course.code === courseCode
      );
      if (selectedCourse) {
        navigate(
          `/courses/${selectedCourse.major_code}-${selectedCourse.code}`
        );
      }
    }
  };

  return (
    <div>
      <Flex justify="space-between" align="center">
        <h1>Courses</h1>
        <Flex justify="space-between" align="center" gap="md">
          <Button
            onClick={() => navigate(`/courses/upload`)}
            variant="outline"
          >
            Upload CSV
          </Button>
          <Button
            onClick={() => navigate(`/courses/create`)}
            variant="outline"
          >
            New <FontAwesomeIcon icon={faAdd} />
          </Button>
        </Flex>
      </Flex>

      {courses.length === 0 && (
        <Skeleton height={50} width="100%" radius="md" />
      )}
      {courses.length > 0 && (
        <>
          <Autocomplete
            placeholder="Search Courses"
            clearable
            data={courses.map((c) => `${c.major_code}-${c.code} ${c.name}`)}
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
              {renderCourses()}
              {renderedChunks < totalChunks && (
                <Flex align="center" justify="center">
                  <Button onClick={() => setRenderedChunks(renderedChunks + 1)}>
                    Load More
                  </Button>
                </Flex>
              )}
            </Accordion>
          </ul>
        </>
      )}
    </div>
  );
};

export default Courses;
