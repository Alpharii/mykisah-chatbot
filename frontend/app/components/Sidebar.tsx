import { LogOut } from "lucide-react"
import { Link, useLocation } from "react-router"

type ChatSession = {
  id: number
  title: string
}

const dummySessions: ChatSession[] = [
  { id: 1, title: "Chat tentang Golang Fiber" },
  { id: 2, title: "Belajar AI Streaming" },
  { id: 3, title: "Debug Remix Loader" },
]

export default function Sidebar(chatSession: any) {
  const location = useLocation()
  const chatMenu: any = chatSession.chatSession

  return (
    <aside className="w-64 h-screen bg-gray-900 text-gray-100 flex flex-col border-r border-gray-800">
      
      {/* Header */}
      <div className="p-4 border-b border-gray-800">
        <h1 className="text-lg font-semibold">My Kisah Chatbot</h1>
        <Link
          to="/chat/new"
          className="mt-3 block w-full text-center text-sm bg-gray-800 hover:bg-gray-700 rounded-md py-2"
        >
          + New Chat
        </Link>
      </div>

      {/* Chat Session List */}
      <div className="flex-1 overflow-y-auto p-2 space-y-1">
        {chatMenu.map(session => {
          const active = location.pathname === `/chat/${session.id}`
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
              {session.Title || "belum ada judul"}
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
