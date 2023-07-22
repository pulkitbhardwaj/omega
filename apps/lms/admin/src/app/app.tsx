// eslint-disable-next-line @typescript-eslint/no-unused-vars
import { Route, Routes } from "react-router-dom";

import Dashboard from "../components/pages/dashboard/dashboard";
import Login from "../components/pages/login/login";
import Register from "../components/pages/register/register";
import Header from "../components/templates/header/header";

export function App() {
  return (
    <>
      <Header />
      <Routes>
        <Route path="/" element={<Dashboard />} />
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />
      </Routes>
    </>
  );
}

export default App;
