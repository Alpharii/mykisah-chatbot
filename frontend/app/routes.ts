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
    // route("register", "routes/_preauth/register.tsx"),
  ]),
  ...prefix("", [
    // index("routes/_postauth/home.tsx"),
    layout("routes/_postauth/layout.tsx", [
      route("dashboard", "routes/_postauth/dashboard.tsx"),
    ]),
  ]),
] satisfies RouteConfig;
