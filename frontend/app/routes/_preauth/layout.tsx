import { Outlet } from "react-router";

export default function PreAuthLayout() {
  return (
    <div className="bg-black">
        <h1>preauth</h1>
        <Outlet />
    </div>
  )
}
