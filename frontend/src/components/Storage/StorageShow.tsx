//import React, { useEffect } from "react";

//import { Link as RouterLink } from "react-router-dom";

import Typography from "@mui/material/Typography";

//import Button from "@mui/material/Button";

//import Container from "@mui/material/Container";

import Box from "@mui/material/Box";

// //import { UsersInterface } from "../models/IUser";

 import { DataGrid, GridColDef } from "@mui/x-data-grid";

import { StorageInterface, ProductTypeInterface, ProductInterface } from "../../models/IStorage";

// import { createTheme, ThemeProvider } from "@mui/material/styles";

// import { grey } from '@mui/material/colors';

// import { EmployeeInterface } from "../../models/IEmployee";
// import { createEmitAndSemanticDiagnosticsBuilderProgram } from "typescript";

import { ButtonGroup, Grid, Table, TableBody, TableCell, TableContainer, TableHead, TableRow } from "@mui/material";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import { Link as RouterLink } from "react-router-dom";
import { useEffect, useState } from "react";

import { grey } from '@mui/material/colors';
import Container from "@mui/material/Container";
import Button from "@mui/material/Button";
import moment from "moment";
import { GetStorages } from "./service/StorageHttpClientService";


const theme = createTheme({
  palette: {
      primary: {
          main: grey[800],
      },
      secondary: {
          main: grey[50],
      },
  },
});


function StorageShow() {

  const [storage, setStorage] = useState<StorageInterface[]>([]);

  const id_cus = localStorage.getItem("id");

  // const getRooms = async () => {
  //     const apiUrl = `http://localhost:8080/service/rooms`;
  //     const requestOptions = {
  //         method: "GET",
  //         headers: {
  //             Authorization: `Bearer ${localStorage.getItem("token")}`,
  //             "Content-Type": "application/json",
  //         },
  //     };

  //     fetch(apiUrl, requestOptions)
  //         .then((response) => response.json())
  //         .then((res) => {
  //             if (res.data) {
  //                 setRoom(res.data);
  //             }
  //         });
  // };

  const getList = async () => {
    let res = await GetStorages();
    if (res) {
      setStorage(res);
    }
  };

 const columns: GridColDef[] = [

   { field: "ID", headerName: "รหัสสินค้า", width: 100 },

   { field: "Employee", headerName: "ชื่อ-นามสกุล", width: 150 , valueFormatter: (params) => params?.value?.Employeename,},

   { field: "Product", headerName: "ชื่อสินค้า", width: 150 , valueFormatter: (params) => params?.value?.Name,},

   { field: "ProductType", headerName: "ประเภทของสินค้า", width: 150 , valueFormatter: (params) => params?.value?.Name,},

   { field: "Time", headerName: "วันที่และเวลา", width: 170, valueFormatter: (params) => moment(params.value).format('DD-MM-yyyy เวลา hh:mm') },

 ];

 useEffect(() => {
  getList();
}, []);


//  useEffect(() => {

//    getRooms();

//  }, []);


 return (

   <div>
     <ThemeProvider theme={theme}>
     <Container maxWidth="md">

       <Box

         display="flex"

         sx={{

           marginTop: 2,

         }}

       >

        <Box flexGrow={1}>
            <Typography // ตาราง
              component="h1"
              variant="h6"
              color="grey"
              gutterBottom
            >

            ข้อมูลคลังสินค้าห้องพัก

           </Typography>

         </Box>

         <Box>
         <Button //ตัวบันทึก
                component={RouterLink} //ลิ้งหน้าต่อไป
                to="/RoomW/Create"
                variant="contained"
                color="primary"
              >
                <Typography
                  color="second"
                  component="div"
                  sx={{ flexGrow: 1 }}
                >
             Add

             </Typography>
           </Button>
         </Box>
         <Box>
         <Button //ตัวบันทึก
                component={RouterLink} //ลิ้งหน้าต่อไป
                to="/RoomW/Edit"
                variant="contained"
                color="primary"
              >
                <Typography
                  color="second"
                  component="div"
                  sx={{ flexGrow: 1 }}
                >
             Edit

             </Typography>
           </Button>
         </Box>
         <Box>
         <Button //ตัวบันทึก
                component={RouterLink} //ลิ้งหน้าต่อไป
                to="/RoomW/Delete"
                variant="contained"
                color="primary"
              >
                <Typography
                  color="second"
                  component="div"
                  sx={{ flexGrow: 1 }}
                >
             Delete

             </Typography>
           </Button>
         </Box>



       </Box>

       <div style={{ height: 400, width: "100%", marginTop: '20px'}}>

         <DataGrid

           rows={storage}

           getRowId={(row) => row.ID}

           columns={columns}

           pageSize={5}

           rowsPerPageOptions={[5]}

         />

       </div>

     </Container>
     </ThemeProvider>
   </div>

 );

}


export default StorageShow