import React, { useEffect, useState } from "react";

import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

import Navbar from "./components/Navbar";

import Manage_Search from "./components/Employee/Manage_Search";
import Manage_Show from "./components/Employee/Manage_Show";
import Manage_Save from "./components/Employee/Manage_Save";
import Home from "./components/Home";
import SignIn from "./components/Login";



export default function App() {

  const [token, setToken] = useState<String>("");

  const [role, setRole] = useState<string>("");

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (token) {
      setToken(token);
      setRole(localStorage.getItem("role") || "");
    }
  }, []);

  if (token) {
    return <SignIn />;
  }

return (

  <Router>
    <div>
      <Navbar/>
   <Routes>

       <Route path="/" element={<Home />} />

       <Route path="/Man" element={<Manage_Save />} />

       <Route path="/ManageShow" element={<Manage_Show />} />

   </Routes>

   </div>

  </Router>

);

}