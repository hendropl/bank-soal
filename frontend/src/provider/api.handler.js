import axios from "axios";
import { API_BASE_URL } from "../core/constant";

const ApiHandler = axios.create({
    baseURL: API_BASE_URL,
    timeout: 10000,
    withCredentials: true
});

ApiHandler.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('token');
    if (token) {
      config.headers["Authorization"] = `Bearer ${token}`
    }
    return config
  },
  (error) => Promise.reject(error)
)

ApiHandler.interceptors.response.use(
  (response) => {
    return response
  },
  async (error) => {
    const originalRequest = error.config;
    console.log(error)
    if (error.response.status === 401 && !originalRequest._retry) {
      originalRequest._retry = true
      
      try {
        const refreshResponse = await axios.post(
          `${API_BASE_URL}/user/refresh`,
          {},
          { withCredentials: true }
        )
        // console.log("Refresh Response", refreshResponse)

        const newAccessToken = refreshResponse.data.data
        // console.log("New", newAccessToken)
        localStorage.setItem('token', newAccessToken)

        originalRequest.headers["Authorization"] = `Bearer ${newAccessToken}`;

        return ApiHandler(originalRequest);
      } catch (error) {
        console.error("Refresh token expired or invalid:", refreshError);
        localStorage.removeItem('token')
      }
    }

    return Promise.reject(error.response.data.error);
  }
)

export default ApiHandler