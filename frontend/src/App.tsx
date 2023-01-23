import React from "react";

import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

import Navbar from "./components/Navbar";

import Manage_Search from "./components/Employee/Manage_Search";
import Manage_Show from "./components/Employee/Manage_Show";
import Manage_Save from "./components/Employee/Manage_Save";

import CheckInOutShow from "./components/CheckInOut/CheckInOutShow";
import Home from "./components/Home";



export default function App() {

return (

  <Router>
    <div>
      <Navbar/>
   <Routes>

       <Route path="/" element={<Home />} />

       <Route path="/ManageSave" element={<Manage_Save />} />

       <Route path="/ManageShow" element={<Manage_Show />} />
       
       <Route path="/CheckInOutShow" element={<CheckInOutShow />} />
       
   </Routes>

   </div>

  </Router>

);

}