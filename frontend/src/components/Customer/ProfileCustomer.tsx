import React, { useState, useEffect } from "react";
import Paper from "@mui/material/Paper";
import Grid from "@mui/material/Unstable_Grid2";
import TextField from "@mui/material/TextField";
import Container from "@mui/material/Container";
import FormLabel from "@mui/material/FormLabel";
import MenuItem from "@mui/material/MenuItem";
import FormHelperText from "@mui/material/FormHelperText";
import Select, { SelectChangeEvent } from "@mui/material/Select";
import Button from "@mui/material/Button";
import Snackbar from "@mui/material/Snackbar";
import MuiAlert, { AlertProps } from "@mui/material/Alert";
import IconButton from "@mui/material/IconButton";
import OutlinedInput from "@mui/material/OutlinedInput";
import InputLabel from "@mui/material/InputLabel";
import InputAdornment from "@mui/material/InputAdornment";
import { Navigate } from "react-router-dom";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import Typography from "@mui/material/Typography";
import Box from "@mui/material/Box";
import { Link as RouterLink } from "react-router-dom";
import DeleteIcon from '@mui/icons-material/Delete';
import SendIcon from '@mui/icons-material/Send';
import Stack from '@mui/material/Stack';
import HouseIcon from '@mui/icons-material/House';
import EditIcon from '@mui/icons-material/Edit';
import FolderIcon from '@mui/icons-material/Folder';


import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell, { tableCellClasses } from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import TablePagination from '@mui/material/TablePagination';
import TableFooter from '@mui/material/TableFooter';

import SearchIcon from '@mui/icons-material/Search';
import styled from "@emotion/styled";




// function ProfileCustomer() {
// //   const apiUrl = "http://localhost:8080";

// //   const navigate = useNavigate();
// //   const params = useParams();

// //   async function GetCustomerByUID () {
// //     let uid = localStrorage.getItem("uid");
// //     const requestOptions = {
// //         method: "GET",
// //         headers: { 
// //           Authorization: `Bearer ${localStorage.getItem("token")}`,
// //           "Content-Type": "application/json" },
// //     };
// //     fetch(`${apiUrl}/customers${customer_id}`, requestOptions)
// //         .then((response) => response.json())
// //         .then((res) => {
  
// //             if (res.data) {
// //               return res.data;
// //             }else{
// //               return false;
// //             }
// //         });
// //   };
  
// //   useEffect(() => {
// //     GetCustomerByUID;
// //   }, []);
  

 

//     return (
//         <div>
//             <Container maxWidth="sm" sx={{ marginTop: 6 }}>
//         <Paper
//           elevation={4}
//           sx={{
//             marginBottom: 2,
//             marginTop: 2,
//             padding: 1,
//             paddingX: 2,
//             display: "flex",
//             justifyContent: "flex-start",
//           }}
//         >
//           <h4 style={{ color: "#000000" }}>Profile</h4>
//         </Paper>
//         <form>
//           <Paper
//             variant="outlined"
//             sx={{ padding: 2, paddingTop: 1, marginBottom: 2 }}
            
//           >
//             <Grid container spacing={2} sx={{ marginBottom: 1.5 }}>

//             {/*=======================================(Title)===========================================================*/}
//               <Grid
//               xs={12}
//               md={8}
//               sx={{ display: "flex", alignItems: "center", margin: 1 }}
//               >
//               <FormLabel
//                 id="demo-simple-select-helper-label"
//                 sx={{ marginRight: 6.5, fontSize: 17, paddingBottom: 2 }}
//               >
//                 Title:
//               </FormLabel>
//               <TextField
//                 // labelId="demo-simple-select-helper-label"
//                 id="NametitleID"
//                 disabled
//                 value={}
                
//                 inputProps={{
//                   name: "Nametitle_ID",
//                 }}
//                 fullWidth
//               />
              
//               <FormHelperText disabled sx={{ width: 350, marginLeft: 2 }}>
//                 คำนำหน้าชื่อ
//               </FormHelperText>
//               </Grid>

//               {/*============================================(First name)======================================================*/}
//               <Grid xs={6} md={6}>
//               <p style={{ color: "grey", fontSize: 17 }}>First name</p>
//               <TextField
//                 id="Fristname"
//                 type="string"
//                 value={}
//                 variant="outlined"
//                 fullWidth
//                 required
//                 disabled
//               />
//               </Grid>
//               {/*=============================================(Last name)=====================================================*/}
//               <Grid xs={6} md={6}>
//               <p style={{ color: "grey", fontSize: 17 }}>Last name</p>
//               <TextField
//                 id="lastname"
//                 type="string"
//                 value={}
//                 variant="outlined"
//                 fullWidth
//                 required
//                 disabled
                
//               />
//               </Grid>
//               </Grid>

//               <Grid container spacing={2} sx={{ marginBottom: 1.5 }}>
//               {/*============================================(Age)======================================================*/}
//               <Grid xs={6} md={6}>
//               <p style={{ color: "grey", fontSize: 17 }}>Age</p>
//               <TextField
//                 id="Age"
//                 type="number"
//                 value = {}
//                 InputProps={{inputProps:{min:0, max:100}}} 
//                 disabled
//                 fullWidth
//                 required
                
//               />
//               </Grid>
//               {/*=============================================(Phone)=====================================================*/}
//               <Grid xs={6} md={6}>
//               <p style={{ color: "grey", fontSize: 17 }}>Phone number</p>
//               <TextField
//                 id="Phone"
//                 type="string"
//                 value = {}
//                 variant="outlined"
//                 fullWidth
//                 required
//                 disabled
//               />
//               </Grid>
//               </Grid>


//               {/*===========================================(email)=======================================================*/}
//               <Grid container spacing={1}/>
//               <Grid
//               xs={12}
//               md={12}
//               sx={{ display: "flex", alignItems: "center", margin: 1 }}
//               >
//               <FormLabel sx={{ marginRight: 7, fontSize: 17 }}>
//                 Email:
//               </FormLabel>
//               <TextField
//                 type="email"
//                 id="outlined-basic"
//                 value = {}
//                 variant="outlined"
//                 required
//                 disabled
                
                  
//                 fullWidth
//               />
//               </Grid>
//               {/*==============================================(password)====================================================*/}
//               <Grid
//               xs={12}
//               md={9}
//               sx={{ display: "flex", alignItems: "center", margin: 1 }}
//               >
//               <InputLabel
//                 htmlFor="outlined-adornment-password"
//                 sx={{ marginRight: 3, fontSize: 17 }}
//               >
//                 Password:
//               </InputLabel>
//               <OutlinedInput
//                 id="outlined-adornment-password"
//                 disabled
//                 value = {}
                
//                 endAdornment={
//                   <InputAdornment position="end">
//                     <IconButton
//                       aria-label="toggle password visibility"
                      
                      
//                       edge="end"
//                     >
                      
//                     </IconButton>
//                   </InputAdornment>
//                 }
//                 inputProps={{ maxLength: 10 }}
//               />
//               </Grid>
//               {/*=======================================(select Gender)===========================================================*/}
//               <Grid
//               xs={12}
//               md={9}
//               sx={{ display: "flex", alignItems: "center", margin: 1 }}
//               >
//               <FormLabel
//                 id="demo-simple-select-helper-label"
//                 sx={{ marginRight: 5.5, fontSize: 17, paddingBottom: 2 }}
//               >
//                 Gender:
//               </FormLabel>
//               <TextField
//                 required
//                 id="Gender_ID"
//                 value = {}
//                 disabled
//                 fullWidth
//                 inputProps={{
//                   name: "Gender_ID",
//                 }}
//               />
                
              
//               <FormHelperText disabled sx={{ width: 350, marginLeft: 2 }}>
//                 เพศ
//               </FormHelperText>
//               </Grid>

//               {/*=======================================(Province)===========================================================*/}
//               <Grid
//               xs={12}
//               md={8}
//               sx={{ display: "flex", alignItems: "center", margin: 1 }}
//               >
//               <FormLabel
//                 id="demo-simple-select-helper-label"
//                 sx={{ marginRight: 4.5, fontSize: 17, paddingBottom: 2 }}
//               >
//                 Province:
//               </FormLabel>
//               <TextField
//                 // labelId="demo-simple-select-helper-label"
//                 id="Province_ID"
//                 value = {}
//                 disabled
                
//                 inputProps={{
//                   name: "Province_ID",
//                 }}
//                 fullWidth
//               />
              
              
//               <FormHelperText disabled sx={{ width: 350, marginLeft: 2 }}>
//                 จังหวัดที่อยู่
//               </FormHelperText>
//               </Grid>
                

//               <Stack direction="row"  spacing={16} >
//                 <Button variant="outlined" color = "error" startIcon={<DeleteIcon />}>
//                 Delete
//               </Button>
//               <Button variant="outlined" startIcon={<EditIcon />}component={RouterLink}
//               to="/customers">
//                   Edit
//               </Button>

//               <Button variant="outlined" color="success" startIcon={<HouseIcon />}>
//                   Home
//               </Button>
             
//             </Stack>
                
//             </Paper></form></Container>

          

//         </div>
//     );

// }

// export default ProfileCustomer;