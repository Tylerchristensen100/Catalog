import { FormHeaders, API } from "../../auth/axios_client";


/**
 * @typedef {object} School
 * @property {number} id - The school ID.
 * @property {string} code - The school code.
 * @property {string} name - The school name.
 */

export const fetchSchools = async () => {
    return API.get("/api/schools").then((res) => res.data);
};






export const updateSchool = async (school) => {
    try {
        return await API
            .put(`/api/admin/school/${school.code}`, school).then((res) => res.data)
            .then((res) => res.data);

    } catch (error) {
        console.error(error);
    }
};

export const createSchool = async (school) => {
    try {
        return await API.post(`/api/admin/school`, school, FormHeaders).then((res) => res.data);
    } catch (error) {
        console.error(error);
    }
};