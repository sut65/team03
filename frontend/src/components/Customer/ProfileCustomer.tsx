import React, { useState, useEffect } from "react";
import Paper from "@mui/material/Paper";
import Grid from "@mui/material/Unstable_Grid2";
import TextField from "@mui/material/TextField";
import Container from "@mui/material/Container";
import FormLabel from "@mui/material/FormLabel";
import FormHelperText from "@mui/material/FormHelperText";
import Button from "@mui/material/Button";
import Snackbar from "@mui/material/Snackbar";
import MuiAlert, { AlertProps } from "@mui/material/Alert";
import IconButton from "@mui/material/IconButton";
import OutlinedInput from "@mui/material/OutlinedInput";
import { Link as RouterLink, useNavigate } from "react-router-dom";
import DeleteIcon from '@mui/icons-material/Delete';

import Stack from '@mui/material/Stack';
import HouseIcon from '@mui/icons-material/House';
import EditIcon from '@mui/icons-material/Edit';
import { CustomerInterface } from "../../models/modelCustomer/ICustomer";
import { NametitleInterface } from "../../models/modelCustomer/INametitle";
import { DeleteCustomer, GetNametitleByUID } from "./service/servicecus";
import { FormControl, Select, SelectChangeEvent } from "@mui/material";
import { ProvinceInterface } from "../../models/modelCustomer/IProvince";
import { GenderInterface } from "../../models/modelCustomer/IGender";



function ProfileCustomer() {
    const [customer, setCustomer] = useState<CustomerInterface>({});
    const [nametitle, setNametitle] = useState<NametitleInterface[]>([]);
    const [gender, setGender] = useState<GenderInterface[]>([]);
    const [province, setProvince] = useState<ProvinceInterface[]>([]);
    let navigate = useNavigate();
    


    const handleChange = (event: SelectChangeEvent) => {
      const name = event.target.name as keyof typeof customer;
      console.log(event.target.name);
      console.log(event.target.value);
      
      setCustomer({
        ...customer,
        [name]: event.target.value,
      });
      console.log(customer);
    };

    const handleInputChange = (
      event: React.ChangeEvent<{ id?: string; value: any }>
    ) => {
      const id = event.target.id as keyof typeof customer;
  
      const { value } = event.target;
  
      setCustomer({ ...customer, [id]: value });
    };
    const Onclick = async (id: number) => {
      let res = await DeleteCustomer(id);
      if (res) {
        window.location.reload();
        localStorage.clear();
        window.location.href = "/";
      }
    }
    

    //-----------เริ่มดึงข้อมูล-----------//
//---------------------Department-------------------------------------
const getNametitle = async () => {
  const apiUrl = `http://localhost:8080/nametitles`;

  const requestOptions = {
    method: "GET",

    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
    },
  };
  //การกระทำ //json
  fetch(apiUrl, requestOptions)
    .then((response) => response.json()) //เรียกได้จะให้แสดงเป็น json ซึ่ง json คือ API

    .then((res) => {
    //   console.log(res.data); //show ข้อมูล

      if (res.data) {
        setNametitle(res.data);
      } else {
        // console.log("else");
      }
    });
};
//---------------------Position-------------------------------------
const getGender = async () => {
  const apiUrl = `http://localhost:8080/customers/genders`;

  const requestOptions = {
    method: "GET",

    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
    },
  };
  //การกระทำ //json
  fetch(apiUrl, requestOptions)
    .then((response) => response.json()) //เรียกได้จะให้แสดงเป็น json ซึ่ง json คือ API

    .then((res) => {
    //   console.log(res.data); //show ข้อมูล

      if (res.data) {
        setGender(res.data);
      } else {
        // console.log("else");
      }
    });
};
//---------------------Position-----------------------------
const getProvince = async () => {
  const apiUrl = `http://localhost:8080/provinces`;

  const requestOptions = {
    method: "GET",

    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
    },
  };
  //การกระทำ //json
  fetch(apiUrl, requestOptions)
    .then((response) => response.json()) //เรียกได้จะให้แสดงเป็น json ซึ่ง json คือ API

    .then((res) => {
    //   console.log(res.data); //show ข้อมูล

      if (res.data) {
        setProvince(res.data);
      } else {
        // console.log("else");
      }
    });
};

    const apiUrl = "http://localhost:8080";

  async function GetCustomer() {
    let uid = localStorage.getItem("id");
    const requestOptions = {
        method: "GET",
        headers: { 
          Authorization: `Bearer ${localStorage.getItem("token")}`,
          "Content-Type": "application/json" },
    };
    fetch(`${apiUrl}/customer/${uid}`, requestOptions)
        .then((response) => response.json())
        .then((res) => {
  
            if (res.data) {
                setCustomer(res.data);
                console.log(res.data);
            }else{
              return false
            }
        });
  }

  // const getNametitleByUID = async () => {
  //   let res = await GetNametitleByUID(customer);
  //   if (res) {
  //       setNametitle(res);
  //       console.log(res);
  //   }
  // };

  useEffect(() => {
    GetCustomer();
    getGender();
    getNametitle();
    getProvince();
    // getNametitleByUID();
  }, []);

  console.log(customer);

    return (
    <div>
        <Container maxWidth="sm" sx={{ marginTop: 6 }}>
    <Paper
      elevation={4}
      sx={{
        marginBottom: 2,
        marginTop: 2,
        padding: 1,
        paddingX: 2,
        display: "flex",
        justifyContent: "flex-start",
      }}
    >
      <h4 style={{ color: "#000000" }}>Profile</h4>
    </Paper>
    <form>
      <Paper
        variant="outlined"
        sx={{ padding: 2, paddingTop: 1, marginBottom: 2 }}
        
      >
        

        {/*=======================================(Title)===========================================================*/}
        <Grid container spacing={2} sx={{ marginBottom: 1.5 }}>
                <Grid
                xs={12}
                md={8}
                sx={{ display: "flex", alignItems: "center", margin: 1 }}
                >
                <FormLabel
                id="demo-simple-select-helper-label"
                sx={{ marginRight: 6.5, fontSize: 17, paddingBottom: 2 }}
                >
                Title:
                </FormLabel>
                <FormControl fullWidth variant="outlined">
                <Select
                native
                disabled
                value={customer.Nametitle_ID+""}
                onChange={handleChange}
                inputProps={{
                  name: "Nametitle_ID",
                }}
              >
                {nametitle.map((item: NametitleInterface) => (
                  <option value={item.ID}>{item.NT_Name}</option>
                ))}
              </Select>
                </FormControl>              
                <FormHelperText disabled sx={{ width: 350, marginLeft: 2 }}>
                คำนำหน้าชื่อ
                </FormHelperText>
                </Grid>

                {/*============================================(First name)======================================================*/}
                <Grid xs={6} md={6}>
                <p style={{ color: "grey", fontSize: 17 }}>First name</p>
                <TextField
                id="Fristname"
                disabled
                fullWidth
                required
                value={customer.FirstName}
                />
                </Grid>
                {/*=============================================(Last name)=====================================================*/}
                <Grid xs={6} md={6}>
                <p style={{ color: "grey", fontSize: 17 }}>Last name</p>
                <TextField
                id="lastname"
                disabled
                fullWidth
                required
                value={customer.LastName}
                />
                </Grid>
                </Grid>

                <Grid container spacing={2} sx={{ marginBottom: 1.5 }}>
                {/*============================================(Age)======================================================*/}
                <Grid xs={6} md={6}>
                <p style={{ color: "grey", fontSize: 17 }}>Age</p>
                <TextField
                id="Age"
                disabled              
                fullWidth
                required
                 value={customer.Age}
                />
                </Grid>
                {/*=============================================(Phone)=====================================================*/}
                <Grid xs={6} md={6}>
                <p style={{ color: "grey", fontSize: 17 }}>Phone number</p>
                <TextField
                id="Phone"
                disabled
                fullWidth
                required
                value={customer.Phone}
                />
                </Grid>
                </Grid>


                {/*===========================================(email)=======================================================*/}
                <Grid container spacing={1}>
                <Grid
                xs={12}
                md={12}
                sx={{ display: "flex", alignItems: "center", margin: 1 }}
                >
                <FormLabel sx={{ marginRight: 7, fontSize: 17 }}>
                Email:
                </FormLabel>
                <TextField
                type="email"
                id="outlined-basic"
                disabled
                required
                fullWidth
                value={customer.Email}
                />
                </Grid>
                
                {/*=======================================(select Gender)===========================================================*/}
                <Grid
                xs={12}
                md={9}
                sx={{ display: "flex", alignItems: "center", margin: 1 }}
                >
                <FormLabel
                id="demo-simple-select-helper-label"
                sx={{ marginRight: 5.5, fontSize: 17, paddingBottom: 2 }}
                >
                Gender:
                </FormLabel>
                <Select
                native
                disabled
                value={customer.Gender_ID+""}
                onChange={handleChange}
                inputProps={{
                  name: "Nametitle_ID",
                }}
              >
                {gender.map((item: GenderInterface) => (
                  <option value={item.ID}>{item.G_Name}</option>
                ))}
              </Select>
                
                </Grid>

                {/*=======================================(Province)===========================================================*/}
                <Grid
                xs={12}
                md={8}
                sx={{ display: "flex", alignItems: "center", margin: 1 }}
                >
                <FormLabel
                id="demo-simple-select-helper-label"
                sx={{ marginRight: 4.5, fontSize: 17, paddingBottom: 2 }}
                >
                Province:
                </FormLabel>
                
                <Select
                native
                disabled
                value={customer.Province_ID+""}
                onChange={handleChange}
                inputProps={{
                  name: "Province_ID",
                }}
              >
                {province.map((item: ProvinceInterface) => (
                  <option value={item.ID}>{item.P_Name}</option>
                ))}
              </Select>
                </Grid>
            

          <Stack direction="row"  spacing={40} >

          <Button variant="outlined" color="success" startIcon={<HouseIcon />} component={RouterLink} to="/">
              Home
          </Button>

          <Button variant="outlined" startIcon={<EditIcon />} onClick={() => navigate(`/customer/edit`)} >
              Edit
          </Button>
         
        </Stack>
            
        </Grid></Paper></form></Container></div>
    )
        
    

}

export default ProfileCustomer;