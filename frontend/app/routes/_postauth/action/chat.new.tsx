import { redirect, type ActionFunctionArgs, type LoaderFunctionArgs } from "react-router"
import { apiClient } from "~/lib/Axios"

export async function action({ request }: ActionFunctionArgs) {
  const res = await apiClient.post("/chat/session/new")
  const sessionId = res.data.ID

  return redirect(`/chat/${sessionId}`)
}
