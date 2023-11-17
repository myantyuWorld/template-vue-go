import axios from "axios";

const baseURL = "http://localhost:8080/";

const axiosInstance = axios.create({
  baseURL,
});
axiosInstance.interceptors.response.use(
  (response) => response,
  (error) => {
    return error.response.data;
  }
);

export default axiosInstance;
