import { LogOut } from "lucide-react"
import { Form, Link, redirect, useLocation, type ActionFunctionArgs } from "react-router"
import { apiClient } from "~/lib/Axios"

type ChatSession = {
  ID: number
  Title: string
}

export async function action({ request }: ActionFunctionArgs) {
  console.log('action', import.meta.env.API_URL)

  try {
    const res = await apiClient.post("/chat/new",);
    const sessionId = res.data

    console.log(sessionId)

    return redirect(`/login/${sessionId}`)
  } catch (error: any) {
    console.log('err', error)
    const message =
      error.response?.data?.error || "Login gagal, periksa kembali email/password";
    return { success: false, error: message, status: 401 }
  }
}

export default function Sidebar(chatSession: any) {
  const location = useLocation()
  const chatMenu: ChatSession[] = chatSession.chatSession

  return (
    <aside className="w-64 h-screen bg-gray-900 text-gray-100 flex flex-col border-r border-gray-800">
      <Form method="post" action="/chat/new">
        {/* Header */}
        <div className="p-4 border-b border-gray-800">
          <Link to="/chat">
            <h1 className="text-lg font-semibold">My Kisah Chatbot</h1>
          </Link>
          <button
            type="submit"
            className="mt-3 block w-full text-center text-sm bg-gray-800 hover:bg-gray-700 rounded-md py-2"
          >
            + New Chat
          </button>
        </div>
      </Form>

      {/* Chat Session List */}
      <div className="flex-1 overflow-y-auto p-2 space-y-1">
        {chatMenu.map(session => {
          const active = location.pathname === `/chat/${session.ID}`
          return (
            <Link
              key={session.ID}
              to={`/chat/${session.ID}`}
              className={`block px-3 py-2 rounded-md text-sm truncate
                ${active 
                  ? "bg-gray-700" 
                  : "hover:bg-gray-800"
                }`}
            >
              {session.Title || "New Chat"}
            </Link>
          )
        })}
      </div>

      {/* Footer */}
      <div className="p-4">
        <form method="post" action="/login">
          <button
            type="submit"
            className="flex items-center gap-2 text-sm text-red-500 hover:text-red-600 transition"
          >
            <LogOut className="w-5 h-5" />
            Logout
          </button>
        </form>
      </div>

    </aside>
  )
}
