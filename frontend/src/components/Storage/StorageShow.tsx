import React from "react";

import Typography from "@mui/material/Typography";

import Box from "@mui/material/Box";

import { Alert, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle, Snackbar } from "@mui/material";


import { DataGrid, GridColDef, GridRowParams } from "@mui/x-data-grid";

import { StorageInterface, ProductInterface, ProductTypeInterface } from "../../models/IStorage";

import { ButtonGroup, Grid, Table, TableBody, TableCell, TableContainer, TableHead, TableRow } from "@mui/material";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import { Link as RouterLink } from "react-router-dom";
import { useEffect, useState } from "react";

import DeleteOutlineIcon from '@mui/icons-material/DeleteOutline';

import { grey } from '@mui/material/colors';
import Container from "@mui/material/Container";
import Button from "@mui/material/Button";
import moment from "moment";
import { GetStorages, Storage, DeleteStorage } from "./service/StorageHttpClientService";


function StorageShow() {

  const [storage, setStorage] = useState<StorageInterface[]>([]);

  const id_cus = localStorage.getItem("id");

  const [success, setSuccess] = useState(false);
  const [successDel, setSuccessDel] = useState(false);
  const [error, setError] = useState(false);
  const [errorDel, setErrorDel] = useState(false);
  const [id, setId] = useState(0);

  //For Delete state 
   const [deleteID, setDeleteID] = React.useState<number>(0)

   // For Set dialog open
   const [openDelete, setOpenDelete] = React.useState(false);
 
   const handleDialogDeleteclose = () => {
     setOpenDelete(false)
     setTimeout(() => {
         setDeleteID(0)
     }, 500)
 }
 
 const handleDialogDeleteOpen = (ID: number) => {
   setDeleteID(ID)
   setOpenDelete(true)
 }
 
 const handleDelete = async () => {
   let res = await DeleteStorage(deleteID)
   if (res.status) {
       console.log(res.data)
   } else {
       console.log(res.data)
   }
   getList();
   setOpenDelete(false)
 
 }

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
    let res = await GetStorages();
    if (res) {
      setStorage(res);
    }
  };

  // const onDelete = async (id: number) => {
  //   let res = await DeleteStorage(id);
  //   if (res) {
  //     setSuccessDel(true);
  //   } else {
  //     setErrorDel(true);
  //   }
  //   getList()
  // }



 const columns: GridColDef[] = [

   { field: "ID", headerName: "รหัสสินค้า", width: 100 },

   { field: "Employee", headerName: "ชื่อ-นามสกุล", width: 150 , valueFormatter: (params) => params?.value?.Employeename,},

   { field: "Product", headerName: "ชื่อสินค้า", width: 150 , valueFormatter: (params) => params?.value?.Name,},

   { field: "ProductType", headerName: "ประเภทของสินค้า", width: 150 , valueFormatter: (params) => params?.value?.Name,},
   
   { field: "Quantity", headerName: "จำนวน", width: 150 , valueFormatter: (params) => params?.value?.Quantity,},

   { field: "Time", headerName: "วันที่และเวลาที่บันทึกข้อมูล", width: 170, valueFormatter: (params) => moment(params.value).format('DD-MM-yyyy เวลา hh:mm') },

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
               //onDelete(row.ID);
               handleDialogDeleteOpen(row.ID)
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
            เพิ่มข้อมูลสำเร็จ
          </Alert>
        </Snackbar>
        <Snackbar
          open={error}
          autoHideDuration={6000}
          onClose={handleClose}
          anchorOrigin={{ vertical: "top", horizontal: "center" }}
        >
          <Alert onClose={handleClose} severity="error">
            ไม่สามารถเพิ่มข้อมูลได้
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
              ข้อมูลคลังสินค้าห้องพัก
            </Typography>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/RoomW/Create"
              variant="contained"
              color="inherit"
            >
              เพิ่มข้อมูลคลังสินค้าห้องพัก
            </Button>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/RoomW/Edit"
              variant="contained"
              color="inherit"
            >
              แก้ไขข้อมูลคลังสินค้าห้อง
            </Button>
          </Box>
          </Box>
        <div style={{ height: 400, width: "100%", marginTop: "20px" }}>
          <DataGrid
            rows={storage}
            getRowId={(row) => row.ID}
            columns={columns}
            pageSize={5}
            rowsPerPageOptions={[5]}
          />
        </div>
        <Dialog
                open={openDelete}
                onClose={handleDialogDeleteclose}
                aria-labelledby="alert-dialog-title"
                aria-describedby="alert-dialog-description"
            >
                <DialogTitle id="alert-dialog-title" >
                    {`คุณต้องการลบข้อมูลคลังสินค้าที่รหัสสินค้า${storage.filter((storage) => (storage.ID === deleteID)).at(0)?.ID} ?`}
                </DialogTitle>
                <DialogContent>
                    <DialogContentText id="alert-dialog-description">
                    ถ้าคุณลบข้อมูลนี้ คุณจะไม่สามารถกู้ข้อมูลคืนได้ คุณแน่ใจหรือไม่ว่าต้องการลบข้อมูลนี้?
                    </DialogContentText>
                </DialogContent>
                <DialogActions>
                    <Button color="inherit" onClick={handleDialogDeleteclose}>ยกเลิก</Button>
                    <Button color="error" onClick={handleDelete} autoFocus>
                        ยืนยัน
                    </Button>
                </DialogActions>
            </Dialog>
      </Container>
    </div>


   
 );

}


export default StorageShow