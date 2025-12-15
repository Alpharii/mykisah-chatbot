import { useLoaderData, type LoaderFunctionArgs } from "react-router";


export async function loader({request, params}: LoaderFunctionArgs) {
    console.log('params',params.id)

    return params.id
}

export default function ChatDetails() {
    const chat = useLoaderData<typeof loader>()

    console.log(chat)

    return (
        <div>
            <h1>{chat}</h1>
        </div>
    )
}
