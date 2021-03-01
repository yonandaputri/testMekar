import axios from "axios";

const baseURL = "http://localhost:3000";

export const getUsers = async () => {
    const res = await axios.get(`${baseURL}/users`);
    return await res;
};

export const getUserById = async (id) => {
    const res = await axios.get(`${baseURL}/user/${id}`);
    return await res;
};

export const createUser = async (user) => {
    const res = await axios.post(`${baseURL}/user`, user);
    return await res;
};

export const updateUser = async (id, user) => {
    const res = await axios.put(`${baseURL}/user/${id}`, user);
    return await res;
};

export const deleteUser = async (id) => {
    const res = await axios.delete(`${baseURL}/user/$id`);
    return await res;
};

export const getJobs = async () => {
    const res = await axios.get(`${baseURL}/job`);
    return await res;
}

export const getEducation = async () => {
    const res = await axios.get(`${baseURL}/education`);
    return await res;
}

