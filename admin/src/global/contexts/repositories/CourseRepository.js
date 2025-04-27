import { FormHeaders, API } from "../../auth/axios_client";



/**
 * @typedef {object} Course
 * @property {number} id - The course ID.
 * @property {string} name - The course name.
 * @property {string} majorCode - The course major code.
 * @property {number} code - The course subject.
 * @property {number} CreditHours - The course credits.
 */


export const fetchCourses = async () => {
    return API.get("/api/courses").then((res) => res.data);
};



export const updateCourse = async (course) => {
    try {
        return await API.put(
            `/api/admin/course/${course.id}`, course).then((res) => res.data);
    } catch (error) {
        console.error(error);
    }
};



export const createCourse = async (course) => {
    try {
        return await API.post(`/api/admin/course`, course, FormHeaders).then((res) => res.data);
    } catch (error) {
        console.error(error);
    }
};
