import { createContext, useContext } from "react";
import { useState, useEffect, useMemo } from "react";
import PropTypes from "prop-types";
import SessionStorageHelper from "../helpers/SessionStorageHelper";
import StorageConstants from "../constants/StorageConstants";
import {
  fetchSchools,
  updateSchool,
} from "../contexts/repositories/SchoolRepository";
import {
  fetchCourses,
  updateCourse,
  createCourse,
} from "../contexts/repositories/CourseRepository";
import {
  fetchPrograms,
  updateProgram,
  createProgram,
  fetchGradLevels,
} from "../contexts/repositories/ProgramRepository";

const DataContext = createContext();

/**
 * Custom hook to access the DataContext.
 * @returns {DataContextValue} The context value.
 */
export function useData() {
  return useContext(DataContext);
}

/**
 * Provides data context to its children components.
 * Fetches and manages courses, programs, schools, and grad levels data.
 * @param {object} props - The component props.
 * @param {React.ReactNode} props.children - The child components.
 * @returns {JSX.Element} The DataProvider component.
 */
export const DataProvider = ({ children }) => {
  const storageHelper = useMemo(() => new SessionStorageHelper(), []);

  const [courses, setCourses] = useState(
    storageHelper.getItem(StorageConstants.coursesKey) ?? []
  );
  const [programs, setPrograms] = useState(
    storageHelper.getItem(StorageConstants.programsKey) ?? []
  );

  const [schools, setSchools] = useState(
    storageHelper.getItem(StorageConstants.schoolsKey) ?? []
  );

  const [gradLevels, setGradLevels] = useState(
    storageHelper.getItem(StorageConstants.gradLevelsKey) ?? []
  );

  useEffect(() => {
    fetchPrograms().then((data) => {
      setPrograms(data);
      storageHelper.setItem(StorageConstants.programsKey, data);
    });

    fetchCourses().then((data) => {
      setCourses(data);
      storageHelper.setItem(StorageConstants.coursesKey, data);
    });

    fetchSchools().then((data) => {
      setSchools(data);
      storageHelper.setItem(StorageConstants.schoolsKey, data);
    });

    fetchGradLevels().then((data) => {
      setGradLevels(data);
      storageHelper.setItem(StorageConstants.gradLevelsKey, data);
    });

    Promise.all([
      fetchCourses(),
      fetchPrograms(),
      fetchSchools(),
      fetchGradLevels(),
    ]);
  }, [storageHelper]);

  /**
   * @typedef {object} CourseActions
   * @property {function(CourseForm): Promise<Course>} updateCourse - Updates a course.
   * @property {function(CourseForm): Promise<Course>} createCourse - Creates a course.
   */

  const courseActions = useMemo(
    () => ({
      updateCourse: async (course) => {
        const updatedCourse = await updateCourse(course);
        const updatedCourses = courses.map((c) =>
          c.id === course.id ? updatedCourse : c
        );
        setCourses(updatedCourses);
        storageHelper.setItem(StorageConstants.coursesKey, updatedCourses);
        return updatedCourse != null;
      },

      createCourse: async (course) => {
        const c = await createCourse(course);
        const updatedCourses = courses.push(c);
        setCourses(updatedCourses);
        storageHelper.setItem(StorageConstants.coursesKey, updatedCourses);
        return c != null;
      },
    }),
    [courses, storageHelper]
  );

  /**
   * @typedef {object} ProgramActions
   * @property {function(ProgramForm): Promise<Program>} updateProgram - Updates a program.
   * @property {function(ProgramForm): Promise<Program>} createProgram - Creates a program
   */
  const programActions = useMemo(
    () => ({
      updateProgram: async (id, program) => {
        const prog = await updateProgram(id, program);
        const updatedPrograms = programs.map((p) => (p.id === id ? prog : p));
        setPrograms(updatedPrograms);
        storageHelper.setItem(StorageConstants.programsKey, updatedPrograms);
        return prog != null;
      },

      createProgram: async (program) => {
        const p = await createProgram(program);
        const updatedPrograms = [...programs, p];

        setPrograms(updatedPrograms);
        storageHelper.setItem(StorageConstants.programsKey, updatedPrograms);
        return p != null;
      },
    }),
    [programs, storageHelper]
  );

  /**
   * @typedef {object} SchoolActions
   * @property {function(SchoolForm): Promise<School>} updateSchool - Updates a school.
   */
  const schoolActions = useMemo(
    () => ({
      updateSchool: async (school) => {
        const updatedSchool = await updateSchool(school);
        const updatedSchools = schools.map((s) =>
          s.id === s.id ? updatedSchool : s
        );
        setPrograms(updatedSchools);
        storageHelper.setItem(StorageConstants.programsKey, updatedSchools);
        return updatedSchool;
      },
    }),
    [schools, storageHelper]
  );

  /**
   * @typedef {object} DataContextValue
   * @property {Course[]} courses - The array of courses.
   * @property {Program[]} programs - The array of programs.
   * @property {School[]} schools - The array of schools.
   * @property {GradLevel[]} gradLevels - The array of grad levels.
   * @property {CourseActions} courseActions - The course actions.
   * @property {ProgramActions} programActions - The program actions.
   * @property {SchoolActions} schoolActions - The school actions.
   */

  /** @type {DataContextValue}  */
  const exported = {
    courses,
    programs,
    schools,
    gradLevels,
    courseActions,
    programActions,
    schoolActions,
  };

  return (
    <DataContext.Provider value={exported}>{children}</DataContext.Provider>
  );
};

DataProvider.propTypes = {
  children: PropTypes.node.isRequired,
};
