import { Outlet } from "react-router";

export default function PreAuthLayout() {
  return (
    <div className="bg-white">
        <Outlet />
    </div>
  )
}
