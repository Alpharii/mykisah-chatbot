import type { LoaderFunctionArgs } from "react-router";
import { Outlet, redirect, useLoaderData } from "react-router";
import Sidebar from "~/components/Sidebar";
import setApiToken, { apiClient, tokenCookie } from "~/lib/Axios";

export async function loader({request}: LoaderFunctionArgs){
  const cookie = request.headers.get("cookie")
  const token = await tokenCookie.parse(cookie)
  if (!token) return redirect("/login")

  setApiToken(token)
  const user = await apiClient.get("/auth/me")
  const chatSession = await apiClient.get("/chat/session")
  return {user: user.data, chatSession: chatSession.data, token}
}

export default function PostAuthLayout() {
  const {user, chatSession, token} = useLoaderData<typeof loader>()
  return (
    <div className="flex">
      <Sidebar chatSession={chatSession} />
      <main className="flex-1 h-screen overflow-y-auto p-4">
        <Outlet context={token}/>
      </main>
    </div>
  )
}