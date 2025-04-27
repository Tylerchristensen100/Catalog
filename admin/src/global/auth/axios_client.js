import axios from 'axios';

export const API = axios.create();

export const setAuthToken = (token) => {
  if (token) {
    API.defaults.headers.common['Authorization'] = `Bearer ${token}`;
  } else {
    delete API.defaults.headers.common['Authorization'];
  }
};

export const FormHeaders = {
  headers: {
    "Content-Type": "multipart/form-data",
  },
};