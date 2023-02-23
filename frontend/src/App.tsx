import React, { Fragment, useEffect, useState } from "react";

import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

import Navbar from "./components/Navbar";

import Manage_Show from "./components/Employee/Manage_Show";
import Manage_Save from "./components/Employee/Manage_Save";
import Manage_Edit from "./components/Employee/Manage_Edit";

import CheckInOutShow from "./components/CheckInOut/CheckInOutShow";
import CheckInOutCreate from "./components/CheckInOut/CheckInOutCreate";
import CheckInOutEdit from "./components/CheckInOut/CheckInOutEdit";

import RepRqShow from "./components/RepReq/RepRqShow";

import Bookings from "./components/Booking/Bookings";
import BookingCreate from "./components/Booking/BookingCreate";
import BookingUpdate from "./components/Booking/BookingUpdate";
import BookingDelete from "./components/Booking/BookingDelete";

import CHK_Payments from "./components/CHK_Payment/CHK_Payment";
import CHK_PaymentCreate from "./components/CHK_Payment/CHK_PaymentCreate";
import CHK_PaymentUpdate from "./components/CHK_Payment/CHK_PaymentUpdate";

import ServiceShow from "./components/Services/ServiceShow";
import ServiceAdd from "./components/Services/ServiceAdd";
import ServiceUpdate from "./components/Services/ServiceUpdate";
import ServiceDelete from "./components/Services/ServiceDelete";

import PaymentAddIn from "./components/Payment/PaymentAddIn";
import PaymentAddOut from "./components/Payment/PaymentAddOut";
import PaymentShow from "./components/Payment/PaymentShow";
import PaymentUpdate from "./components/Payment/PaymentUpdate";

import Home from "./components/Home";
import SignIn from "./components/Login";
import Review_Show from "./components/Review/Review_Show";
import Review_Save from "./components/Review/Review_list";
import Review_list from "./components/Review/Review_list";
//import Review_Save from "./components/Review/Review_Save";

import RoomShow from "./components/Room/RoomShow";
import RoomCreate from "./components/Room/RoomCreate";
import RoomEdit from "./components/Room/RoomEdit";

import StorageShow from "./components/Storage/StorageShow";
import StorageCreate from "./components/Storage/StorageCreate";
import StorageEdit from "./components/Storage/StorageEdit";

import RepRqCreate from "./components/RepReq/RepRqCreate";
import RepRqEdit from "./components/RepReq/RepRqEdit";

import Customer from "./components/Customer/CreateCustomer";
import ProfileCustomer from "./components/Customer/ProfileCustomer";
import Checkroomlist from "./components/Checkroom/Checkroomlist";
import CheckroomEdit from "./components/Checkroom/CheckroomEdit";
import Checkroom from "./components/Checkroom/Createcheckroom";
import EditCustomer from "./components/Customer/EditCustomer";
import CustomerlistforAdmin from "./components/Customer/ShowCustomerforAdmin";
import Homepage from "./components/Homepage";
import Homeshow from "./components/Homeshow";

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
    return (
      <Router>
        <Routes>
          <Route path="/" element={<Homepage />} />
          <Route path="/Homeshow" element={<Homeshow />} />
          <Route path="/RW" element={<Review_Show />} />
          <Route path="/home" element={<SignIn />} />
        </Routes>
      </Router>
    );
  }

  return (
    <Router>
      {token && (
        <Fragment>
          <Navbar />
          <Routes>
            {role === "Officer" && (
              <>
                <Route path="/home" element={<Home />} />
                <Route path="/Manage-Save" element={<Manage_Save />} />
                <Route path="/Manage-Show" element={<Manage_Show />} />
                <Route path="/Manage-Edit/:id" element={<Manage_Edit />} />
              </>
              )
            }
            {role === "Customer" && (
              <>
                <Route path="/home" element={<Home />} />
                <Route path="/Book" element={<Bookings />} />
                <Route path="/Book/Create" element={<BookingCreate />} />
                <Route path="/Book/Edit" element={<BookingUpdate />} />
                <Route path="/Book/Delete" element={<BookingDelete />} />

                <Route path="/Rep" element={<RepRqShow />} />
                <Route path="/Rep/Create" element={<RepRqCreate />} />
                <Route path="/Rep/Edit/:id" element={<RepRqEdit />} />


                <Route path="/Reviewlist" element={<Review_list />} />

                <Route path="/ss" element={<ServiceShow />} />
                <Route path="/sa" element={<ServiceAdd />} />
                <Route path="/su/:id" element={<ServiceUpdate />} />
                <Route path="/sd" element={<ServiceDelete />} />

                <Route path="/ps" element={<PaymentShow />} />
                <Route path="/pai" element={<PaymentAddIn />} />
                <Route path="/pao" element={<PaymentAddOut />} />
                <Route path="/pu/:id" element={<PaymentUpdate />} />

                <Route path="/customer/create" element={<Customer />} />
                <Route path="/customer/profile" element={<ProfileCustomer />} /> 
                <Route path="/customer/edit" element={<EditCustomer />} /> 
                <Route path="/customer/showforadmin" element={<CustomerlistforAdmin />} />
                </>
            )
            }

              {role === "Employee" && (
              <>
              <Route path="/home" element={<Home />} />
              
              <Route path="/RT" element={<RoomShow />} />
              <Route path="/RT/Create" element={<RoomCreate />} />
              <Route path="/RT/Edit" element={<RoomEdit />} />

              <Route path="/CPM" element={<CHK_Payments />} />
              <Route path="/CPM/Create" element={<CHK_PaymentCreate />} />
              <Route path="/CPM/Edit" element={<CHK_PaymentUpdate />} />

              <Route path="/RoomW" element={<StorageShow />} />
              <Route path="/RoomW/Create" element={<StorageCreate />} />
              <Route path="/RoomW/Edit" element={<StorageEdit />} />

              <Route path="/CNCO" element={<CheckInOutShow />} />
              <Route path="/CNCO/Create" element={<CheckInOutCreate />}/> 
              <Route path="/CNCO/Edit/:id" element={<CheckInOutEdit />}/> 

              <Route path="/checkroom/create" element={<Checkroom />} />
              <Route path="/checkroom/list" element={<Checkroomlist />} />
              <Route path="/checkroom/edit" element={<CheckroomEdit />} />

              <Route path="/customer/showforadmin" element={<CustomerlistforAdmin />} />
              </>
            )
            }
          </Routes>
        </Fragment>
      )
    }
    </Router>
  );
}
