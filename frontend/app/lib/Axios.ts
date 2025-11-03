import axios from "axios";
import { createCookie } from "react-router";

let authToken: string | null = null

export const tokenCookie = createCookie("token", {
  httpOnly: true,
  path: "/",
  sameSite: "lax",
  maxAge: 60 * 60 * 24,
})

export const apiClient = axios.create({
    baseURL: `${import.meta.env.VITE_API_URL}/api`,
    timeout: 10_000,
    withCredentials: true,
})

apiClient.interceptors.request.use(
    (config) => {
        if(authToken){
            config.headers["Authorization"] = `Bearer ${authToken}`
        }
        return config
    },
    (error) => Promise.reject(error)
)

export default function setApiToken(token: string){
    authToken = token
}