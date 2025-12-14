import type { Route } from "./+types/home";
import { tokenCookie } from "~/lib/Axios";
import { redirect } from "react-router";

export function meta({}: Route.MetaArgs) {
  return [
    { title: "New React Router App" },
    { name: "description", content: "Welcome to React Router!" },
  ];
}

export async function loader({request}: Route.LoaderArgs){
  console.log('request', request)
  const cookie = request.headers.get("cookie")
  const token = await tokenCookie.parse(cookie)
  if(!token) return redirect("/login")
  
  return redirect("/chat")
}