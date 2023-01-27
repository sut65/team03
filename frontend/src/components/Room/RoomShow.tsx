//import React, { useEffect } from "react";

//import { Link as RouterLink } from "react-router-dom";

import Typography from "@mui/material/Typography";

//import Button from "@mui/material/Button";

//import Container from "@mui/material/Container";

import Box from "@mui/material/Box";

// //import { UsersInterface } from "../models/IUser";

 import { DataGrid, GridColDef } from "@mui/x-data-grid";

import { RoomInterface, RoomTypeInterface, RoomZoneInterface,StateInterface } from "../../models/IRoom";

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
import { GetRooms } from "./service/RoomHttpClientService";


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


function RoomShow() {

  const [room, setRoom] = useState<RoomInterface[]>([]);

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
    let res = await GetRooms();
    if (res) {
      setRoom(res);
    }
  };

 const columns: GridColDef[] = [

   { field: "ID", headerName: "หมายเลขห้อง", width: 100 },

   { field: "Employee", headerName: "ชื่อ-นามสกุล", width: 150 , valueFormatter: (params) => params?.value?.Employeename,},

   { field: "RoomType", headerName: "ประเภทของห้อง", width: 150 , valueFormatter: (params) => params?.value?.Size,},

   { field: "RoomZone", headerName: "โซนของห้องพัก", width: 150 , valueFormatter: (params) => params?.value?.Name,},

   { field: "State", headerName: "สถานะของห้อง", width: 150 , valueFormatter: (params) => params?.value?.Name,},

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

            ข้อมูลห้อง

           </Typography>

         </Box>

         <Box>
         <Button //ตัวบันทึก
                component={RouterLink} //ลิ้งหน้าต่อไป
                to="/RT/Create"
                variant="contained"
                color="primary"
              >
                <Typography
                  color="second"
                  component="div"
                  sx={{ flexGrow: 1 }}
                >
             เพิ่มข้อมูลห้อง

             </Typography>
           </Button>
         </Box>
         <Box>
         <Button //ตัวบันทึก
                component={RouterLink} //ลิ้งหน้าต่อไป
                to="/RT/Edit"
                variant="contained"
                color="primary"
              >
                <Typography
                  color="second"
                  component="div"
                  sx={{ flexGrow: 1 }}
                >
             แก้ไขข้อมูลห้อง

             </Typography>
           </Button>
         </Box>
         <Box>
         <Button //ตัวบันทึก
                component={RouterLink} //ลิ้งหน้าต่อไป
                to="/RT/Delete"
                variant="contained"
                color="primary"
              >
                <Typography
                  color="second"
                  component="div"
                  sx={{ flexGrow: 1 }}
                >
             ลบข้อมูลห้อง

             </Typography>
           </Button>
         </Box>



       </Box>

       <div style={{ height: 400, width: "100%", marginTop: '20px'}}>

         <DataGrid

           rows={room}

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


export default RoomShow