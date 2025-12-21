import {
  Box,
  Paper,
  TextField,
  IconButton,
  Typography,
} from "@mui/material"
import { Send } from "lucide-react"
import { useLoaderData, useOutletContext, useParams, type LoaderFunctionArgs } from "react-router"
import { useEffect, useRef, useState } from "react"
import { apiClient } from "~/lib/Axios"

type ChatMessage = {
  ID?: number
  SessionID: number
  Role: "user" | "assistant"
  Content: string
  CreatedAt?: string
  Error?: boolean
  ClientOnly?: boolean
}

export async function loader({request, params}: LoaderFunctionArgs) {
    const chatId = params.id
    const chat = await apiClient.get(`/chat/session/${chatId}`)

    return chat.data
}

export default function ChatDetails() {
    const params = useParams()
    const dbMessages = useLoaderData() as ChatMessage[]
    const [messages, setMessages] = useState<ChatMessage[]>(dbMessages)
    const [input, setInput] = useState("")
    const bottomRef = useRef<HTMLDivElement>(null)
    const sessionId = Number(params.id)

    useEffect(() => {
      setMessages(prev => {
        const clientOnlyMessages = prev.filter(m => m.ClientOnly)
        return [...dbMessages, ...clientOnlyMessages]
      })
    }, [dbMessages])



    useEffect(() => {
        bottomRef.current?.scrollIntoView({ behavior: "smooth" })
    }, [messages])

    async function handleSend() {
      if (!input.trim()) return

      const userMessage: ChatMessage = {
        SessionID: sessionId,
        Role: "user",
        Content: input,
      }

      setMessages(prev => [...prev, userMessage])
      setInput("")

      const res = await apiClient.post("/chat/send", {
        sessionId,
        message: userMessage.Content,
      })

      // ✅ placeholder HARUS ClientOnly
      const tempAiMessage: ChatMessage = {
        SessionID: sessionId,
        Role: "assistant",
        Content: "",
        ClientOnly: true,
      }

      setMessages(prev => [...prev, tempAiMessage])

      const aiRes = await apiClient.get(res.data.stream_url)

      setMessages(prev => {
        const copy = [...prev]
        const lastIndex = copy
          .map(m => m.ClientOnly)
          .lastIndexOf(true)

        if (lastIndex !== -1) {
          copy[lastIndex] = {
            ...aiRes.data.message,
            ClientOnly: true,
          }
        }

        return copy
      })
    }


    useEffect(() => {
      const last = messages.at(-1)
      if (last?.Error) {
        const timer = setTimeout(() => {
          setMessages(prev => prev.filter(m => m !== last))
        }, 5000)

        return () => clearTimeout(timer)
      }
    }, [messages])



  return (
    <Box className="flex flex-col h-full bg-gray-950 text-gray-100">
      {/* Header */}
      <Box className="px-6 py-4 border-b border-gray-800">
        <Typography variant="h6">Chat</Typography>
      </Box>

      {/* Messages */}
      <Box className="flex-1 overflow-y-auto px-6 py-4 space-y-4">
        {messages.map((msg, i) => {
          const isUser = msg.Role === "user"
          return (
            <Box
              key={i}
              className={`flex ${isUser ? "justify-end" : "justify-start"}`}
            >
              <Paper
                className={`
                  max-w-[75%] px-4 py-3 rounded-2xl text-sm
                  ${isUser
                    ? "bg-blue-600 text-white"
                    : "bg-gray-800 text-gray-100"}
                `}
              >
                {msg.Content || (
                  msg.Error
                    ? null
                    : <span className="opacity-50">AI sedang mengetik…</span>
                )}
              </Paper>
            </Box>
          )
        })}
        <div ref={bottomRef} />
      </Box>

      {/* Input */}
      <Box className="border-t border-gray-800 p-4">
        <Paper className="flex items-center gap-2 bg-gray-900 px-3 py-2 rounded-xl">
          <TextField
            fullWidth
            placeholder="Ketik pesan..."
            variant="standard"
            value={input}
            onChange={e => setInput(e.target.value)}
            onKeyDown={e => e.key === "Enter" && handleSend()}
            InputProps={{ disableUnderline: true }}
          />
          <IconButton onClick={handleSend}>
            <Send size={20} />
          </IconButton>
        </Paper>
      </Box>
    </Box>
  )
}