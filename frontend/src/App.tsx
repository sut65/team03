import React from "react";

import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

import Navbar from "./components/Navbar";

import Manage_Search from "./components/Employee/Manage_Search";
import Manage_Show from "./components/Employee/Manage_Show";


export default function App() {

return (

  <Router>

   <div>

   <Navbar />

   <Routes>

       <Route path="/" element={<Manage_Search />} />

       <Route path="/create" element={<Manage_Show />} />

   </Routes>

   </div>

  </Router>

);

}