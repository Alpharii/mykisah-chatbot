import type { LoaderFunctionArgs } from "react-router";


export async function loader({request, params}: LoaderFunctionArgs) {
    console.log('params',params.id)
}

export default function ChatDetails() {


    return (
        <div>
            <h1>testes</h1>
        </div>
    )
}
