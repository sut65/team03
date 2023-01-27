import React, { useEffect, useState } from "react";

import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

import Navbar from "./components/Navbar";

import Manage_Search from "./components/Employee/Manage_Search";
import Manage_Show from "./components/Employee/Manage_Show";
import Manage_Save from "./components/Employee/Manage_Save";

import CheckInOutShow from "./components/CheckInOut/CheckInOutShow";
import CheckInOutCreate from "./components/CheckInOut/CheckInOutCreate";

import Bookings from "./components/Booking/Bookings";
import BookingCreate from "./components/Booking/BookingCreate";
import BookingUpdate from "./components/Booking/BookingUpdate";
import BookingDelete from "./components/Booking/BookingDelete";

import CHK_Payments from "./components/CHK_Payment/CHK_Payment";
import CHK_PaymentCreate from "./components/CHK_Payment/CHK_PaymentCreate";
import CHK_PaymentUpdate from "./components/CHK_Payment/CHK_PaymentUpdate";

import Customer from "./components/Customer/CreateCustomer";

import ServiceShow from "./components/Services/ServiceShow";
import ServiceAdd from "./components/Services/ServiceAdd";
import ServiceUpdate from "./components/Services/ServiceUpdate";
import ServiceDelete from "./components/Services/ServiceDelete";

import Home from "./components/Home";
import SignIn from "./components/Login";
import Review_Show from "./components/Review/Review_Show";
import Review_Save from "./components/Review/Review_Save";
import CheckInOutEdit from "./components/CheckInOut/CheckInOutEdit";
import Manage_Edit from "./components/Employee/Manage_Edit";
import CheckInEdit from "./components/CheckInOut/CheckInEdit";
import CheckOutEdit from "./components/CheckInOut/CheckOutEdit";

import RoomShow from "./components/Room/RoomShow";
import RoomCreate from "./components/Room/RoomCreate";
import RoomDelete from "./components/Room/RoomDelete";
import RoomEdit from "./components/Room/RoomEdit";

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

  if (!token) {
    return <SignIn />;
  }

  return (
    <Router>
      <div>
        <Navbar />
        <Routes>

          <Route path="/" element={<Home />} />

       <Route path="/Manage-Save" element={<Manage_Save />} />
       <Route path="/Manage-Show" element={<Manage_Show />} />
       <Route path="/Manage-Edit/:id" element={<Manage_Edit />} />
       

          <Route path="/Book" element={<Bookings />} />
          <Route path="/Book/Create" element={<BookingCreate />} />
          <Route path="/Book/Edit" element={<BookingUpdate />} />
          <Route path="/Book/Delete" element={<BookingDelete />} />

          <Route path="/CPM" element={<CHK_Payments />} />
          <Route path="/CPM/Create" element={<CHK_PaymentCreate />} />
          <Route path="/CPM/Edit" element={<CHK_PaymentUpdate />} />

          <Route path="/CNCO" element={<CheckInOutShow />} />
          <Route path="/CNCO/Create" element={<CheckInOutCreate />}/> 
          <Route path="/CNCO/CN/Edit" element={<CheckInEdit />}/> 
          <Route path="/CNCO/CO/Edit" element={<CheckOutEdit />}/> 
          {/* <Route path="/CNCO/Edit" element={<CheckInOutEdit />}/>  */}


          <Route path="/RW" element={<Review_Show />} />
          <Route path="/ReviewSave" element={<Review_Save />} />

          <Route path="/ss" element={<ServiceShow />} />
          <Route path="/sa" element={<ServiceAdd />} />
          <Route path="/su/:id" element={<ServiceUpdate />} />
          <Route path="/sd" element={<ServiceDelete />} />

          <Route path="/RT" element={<RoomShow />} />
          <Route path="/RT/Create" element={<RoomCreate />} />
          <Route path="/RT/Delete" element={<RoomDelete />} />
          <Route path="/RT/Edit" element={<RoomEdit />} />


        </Routes>

      </div>

    </Router>

  );

}

///Ui สมัคร customer อันนี้ => <Customer />