import type { LoaderFunctionArgs } from "react-router";
import { Outlet, redirect, useLoaderData } from "react-router";
import setApiToken, { apiClient, tokenCookie } from "~/lib/Axios";

export async function loader({request}: LoaderFunctionArgs){
  const cookie = request.headers.get("cookie")
  const token = await tokenCookie.parse(cookie)
  if (!token) return redirect("/login")

  setApiToken(token)
  const user = await apiClient.get("/auth/me")
  return user.data
}

export default function PostAuthLayout() {
  const user = useLoaderData<typeof loader>()
  console.log('user', user)
  return (
    <div className="">
        <h1>postauth</h1>
        <Outlet />
    </div>
  )
}
