import {
  type RouteConfig,
  route,
  layout,
  index,
  prefix,
} from "@react-router/dev/routes";

export default [
    layout("routes/_preauth/layout.tsx", [
    route("login", "routes/_preauth/login.tsx"),
    route("register", "routes/_preauth/register.tsx"),
  ]),
  ...prefix("", [
    index("routes/index.tsx"),
    layout("routes/_postauth/layout.tsx", [
      route("chat", "routes/_postauth/chat.tsx"),
      route("chat/:id", "routes/_postauth/chat.$id.tsx"),
    ]),
  ]),
] satisfies RouteConfig;
