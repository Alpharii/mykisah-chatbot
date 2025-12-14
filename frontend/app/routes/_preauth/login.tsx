import {
  Box,
  Typography,
  TextField,
  Button,
  InputAdornment,
  IconButton,
  Paper,
  ThemeProvider,
  CssBaseline,
  Alert,
} from "@mui/material";
import { Eye, EyeOff, Lock, LogIn, User } from "lucide-react";
import { useState } from "react";
import { Form, Link, redirect, useActionData, type ActionFunctionArgs } from "react-router";
import { apiClient, tokenCookie } from "~/lib/Axios";
import { darkTheme } from "~/lib/theme";


export async function action({ request }: ActionFunctionArgs) {
  console.log('action', import.meta.env.API_URL)
  const formData = await request.formData();
  const email = formData.get("email");
  const password = formData.get("password");

  if (typeof email !== "string" || typeof password !== "string") {
    return{
      success: false,
      error: "Email dan password wajib diisi",
      status: 400
    }
  }

  try {
    const res = await apiClient.post("/auth/login", { email, password });
    const token = res.data.token;

    return redirect("/chat", {
      headers: {
        "Set-Cookie": await tokenCookie.serialize(token),
      },
    });
  } catch (error: any) {
    console.log('err', error)
    const message =
      error.response?.data?.error || "Login gagal, periksa kembali email/password";
    return { success: false, error: message, status: 401 }
  }
}


// ✅ CLIENT COMPONENT (RENDER FORM)
export default function LoginPage() {
  const [showPassword, setShowPassword] = useState(false);
  const actionData = useActionData<typeof action>();

  return (
    <ThemeProvider theme={darkTheme}>
      <CssBaseline />
      <Box
        sx={{
          display: "flex",
          minHeight: "100vh",
          bgcolor: "background.default",
        }}
      >
        <Box
          sx={{
            flex: 1,
            display: { xs: "none", md: "flex" },
            backgroundImage:
              "url('https://images.unsplash.com/photo-1557683316-973673baf926?auto=format&fit=crop&w=1400&q=80')",
            backgroundSize: "cover",
            backgroundPosition: "center",
            borderRight: "1px solid rgba(255,255,255,0.1)",
          }}
        />

        <Box
          sx={{
            flex: 1,
            display: "flex",
            alignItems: "center",
            justifyContent: "center",
            px: 4,
          }}
        >
          <Paper
            elevation={0}
            sx={{
              width: "100%",
              maxWidth: 420,
              p: 5,
              borderRadius: 4,
              bgcolor: "background.paper",
              boxShadow: "0 0 40px rgba(0,0,0,0.4)",
            }}
          >
            <Box textAlign="center" mb={4}>
              <LogIn size={42} color="#3b82f6" />
              <Typography variant="h5" fontWeight="bold" mt={2}>
                Welcome Back
              </Typography>
              <Typography variant="body2" color="text.secondary">
                Sign in to your account
              </Typography>
            </Box>

            {actionData?.error && (
              <Alert severity="error" sx={{ mb: 2 }}>
                {actionData.error}
              </Alert>
            )}

            <Form method="post">
              <TextField
                label="Email"
                name="email"
                type="email"
                fullWidth
                required
                margin="normal"
                variant="outlined"
                InputProps={{
                  startAdornment: (
                    <InputAdornment position="start">
                      <User size={18} />
                    </InputAdornment>
                  ),
                }}
              />

              <TextField
                label="Password"
                name="password"
                type={showPassword ? "text" : "password"}
                fullWidth
                required
                margin="normal"
                variant="outlined"
                InputProps={{
                  startAdornment: (
                    <InputAdornment position="start">
                      <Lock size={18} />
                    </InputAdornment>
                  ),
                  endAdornment: (
                    <InputAdornment position="end">
                      <IconButton
                        onClick={() => setShowPassword(!showPassword)}
                        edge="end"
                      >
                        {showPassword ? (
                          <EyeOff size={18} />
                        ) : (
                          <Eye size={18} />
                        )}
                      </IconButton>
                    </InputAdornment>
                  ),
                }}
              />

              <Button
                type="submit"
                variant="contained"
                fullWidth
                sx={{
                  mt: 3,
                  py: 1.3,
                  fontWeight: "bold",
                  textTransform: "none",
                  borderRadius: 2,
                }}
              >
                Sign In
              </Button>
            </Form>

            <Typography
              variant="body2"
              color="text.secondary"
              textAlign="center"
              mt={3}
            >
              Don’t have an account?{" "}
              <Link to={"/register"}>
                <Typography
                  component="span"
                  color="primary"
                  sx={{ cursor: "pointer" }}
                >
                  Sign up
                </Typography>
              </Link>
            </Typography>
          </Paper>
        </Box>
      </Box>
    </ThemeProvider>
  );
}
