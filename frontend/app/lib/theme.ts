import { createTheme } from "@mui/material";

export const darkTheme = createTheme({
  palette: {
    mode: "dark",
    background: {
      default: "#0f1117",
      paper: "#1a1d24",
    },
    primary: {
      main: "#3b82f6",
    },
  },
  typography: {
    fontFamily: "Inter, sans-serif",
  },
});
