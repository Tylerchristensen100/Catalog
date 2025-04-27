import { FormHeaders, API } from "../../auth/axios_client";


/**
 * @typedef {object} Program
 * @property {number} id - The program ID.
 * @property {string} name - The program name.
 * @property {string} gradLevel - The program grad level.
 * @property {string} programType - The program type.
 * @property {string} school - The program school.
 * @property {string} majorCode - The program major code.
 * @property {string} campus - The program campus.
 * @property {boolean} online - The program online status.
 * @property {number} cip - The program CIP code.
 * @property {string} description - The program description.
 */


/**
 * @typedef {object} GradLevel
 * @property {number} id - The grad level ID.
 * @property {string} level - The grad level name.
 */


export const fetchPrograms = async () => {
    return API.get("/api/programs").then((res) => res.data);
};



export const updateProgram = async (id, program) => {
    try {
        return await API.put(`/api/admin/program/${id}`, program).then((res) => res.data);
    } catch (error) {
        console.error(error);
    }
};

export const createProgram = async (program) => {
    try {
        return  await API.post(`/api/admin/program`, program, FormHeaders).then((res) => res.data);

    } catch (error) {
        console.error(error);
    }
};

export const fetchGradLevels = async () => {
    return await API.get("/api/grad-levels").then((res) => res.data);
};