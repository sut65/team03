//import React, { useEffect } from "react";

//import { Link as RouterLink } from "react-router-dom";

import Typography from "@mui/material/Typography";

//import Button from "@mui/material/Button";

//import Container from "@mui/material/Container";

import Box from "@mui/material/Box";

import { Alert, Snackbar } from "@mui/material";

// //import { UsersInterface } from "../models/IUser";

import { DataGrid, GridColDef, GridRowParams } from "@mui/x-data-grid";

import { RoomInterface, RoomTypeInterface, RoomZoneInterface,StateInterface } from "../../models/IRoom";

// import { createTheme, ThemeProvider } from "@mui/material/styles";

// import { grey } from '@mui/material/colors';

// import { EmployeeInterface } from "../../models/IEmployee";
// import { createEmitAndSemanticDiagnosticsBuilderProgram } from "typescript";

import { ButtonGroup, Grid, Table, TableBody, TableCell, TableContainer, TableHead, TableRow } from "@mui/material";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import { Link as RouterLink } from "react-router-dom";
import { useEffect, useState } from "react";

import DeleteOutlineIcon from '@mui/icons-material/DeleteOutline';

import { grey } from '@mui/material/colors';
import Container from "@mui/material/Container";
import Button from "@mui/material/Button";
import moment from "moment";
import { GetRooms, Room, DeleteRoom } from "./service/RoomHttpClientService";


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

  const [success, setSuccess] = useState(false);
  const [successDel, setSuccessDel] = useState(false);
  const [error, setError] = useState(false);
  const [errorDel, setErrorDel] = useState(false);
  const [id, setId] = useState(0);

  

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

  const handleClose = (
    event?: React.SyntheticEvent | Event,
    reason?: string
  ) => {
    if (reason === "clickaway") {
      return;
    }
    setSuccess(false);
    setSuccessDel(false);
    setError(false);
    setErrorDel(false);
  };

  const getList = async () => {
    let res = await GetRooms();
    if (res) {
      setRoom(res);
    }
  };

  const onDelete = async (id: number) => {
    let res = await DeleteRoom(id);
    if (res) {
      setSuccessDel(true);
    } else {
      setErrorDel(true);
    }
    getList()
  }



 const columns: GridColDef[] = [

   { field: "Room_No", headerName: "หมายเลขห้อง", width: 150 , valueFormatter: (params) => params?.value?.Room_No,},
   
   { field: "Employee", headerName: "ชื่อ-นามสกุล", width: 150 , valueFormatter: (params) => params?.value?.Employeename,},

   { field: "RoomType", headerName: "ประเภทของห้อง", width: 150 , valueFormatter: (params) => params?.value?.Size,},

   { field: "RoomZone", headerName: "โซนของห้องพัก", width: 150 , valueFormatter: (params) => params?.value?.Name,},

   { field: "State", headerName: "สถานะของห้อง", width: 150 , valueFormatter: (params) => params?.value?.Name,},

   { field: "Amount", headerName: "ราคาของห้อง", width: 150 , valueFormatter: (params) => params?.value?.Amount,},

   { field: "Time", headerName: "วันที่และเวลา", width: 150, valueFormatter: (params) => moment(params.value).format('DD-MM-yyyy เวลา hh:mm') },

   {
    field: "delete",
    headerName: "",
    sortable: true,
    width: 120,
    align:"center",
    headerAlign: "center",
    renderCell: ({ row }: Partial<GridRowParams>) =>
        <Button 
            size="small"
            //variant="contained"
            color="error"
            onClick={() => {
               onDelete(row.ID);
            }}
            sx={{borderRadius: 20,'&:hover': {color: '#FC0000', backgroundColor: '#F9EBEB'}}}
            endIcon={<DeleteOutlineIcon />}
        >
            DELETE
        </Button>,
  },
 ];

 useEffect(() => {
  getList();
}, []);


//  useEffect(() => {

//    getRooms();

//  }, []);


 return (
  <div>
      <Container maxWidth="lg">
        <Snackbar
          open={successDel}
          autoHideDuration={3000}
          onClose={handleClose}
          anchorOrigin={{ vertical: "top", horizontal: "center" }}
        >
          <Alert onClose={handleClose} severity="success">
            ลบข้อมูลสำเร็จ
          </Alert>
        </Snackbar>
        <Snackbar
          open={errorDel}
          autoHideDuration={6000}
          onClose={handleClose}
          anchorOrigin={{ vertical: "top", horizontal: "center" }}
        >
          <Alert onClose={handleClose} severity="error">
            ไม่สามารถลบข้อมูลได้
          </Alert>
        </Snackbar>
        <Snackbar
          open={success}
          autoHideDuration={3000}
          onClose={handleClose}
          anchorOrigin={{ vertical: "top", horizontal: "center" }}
        >
          <Alert onClose={handleClose} severity="success">
            Add สำเร็จ
          </Alert>
        </Snackbar>
        <Snackbar
          open={error}
          autoHideDuration={6000}
          onClose={handleClose}
          anchorOrigin={{ vertical: "top", horizontal: "center" }}
        >
          <Alert onClose={handleClose} severity="error">
            ไม่สามารถ Add ได้
          </Alert>
        </Snackbar>
        <Box
          display="flex"
          sx={{
            marginTop: 2,
          }}
        >
          <Box flexGrow={1}>
            <Typography
              component="h2"
              variant="h6"
              color="grey"
              gutterBottom
            >
              ข้อมูลห้อง
            </Typography>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/RT/Create"
              variant="contained"
              color="info"
            >
              Add
            </Button>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/RT/Edit"
              variant="contained"
              color="info"
            >
              Edit
            </Button>
          </Box>
          </Box>
        <div style={{ height: 400, width: "100%", marginTop: "20px" }}>
          <DataGrid
            rows={room}
            getRowId={(row) => row.ID}
            columns={columns}
            pageSize={5}
            rowsPerPageOptions={[5]}
          />
        </div>
      </Container>
    </div>


   
 );

}


export default RoomShow