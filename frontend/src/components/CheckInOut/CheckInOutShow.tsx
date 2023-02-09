import React, { useState, useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { Alert, Snackbar } from "@mui/material";
import DeleteOutlineIcon from '@mui/icons-material/DeleteOutline';
import CheckCircleIcon from '@mui/icons-material/CheckCircle';

import { CheckInOutInterface } from "../../models/ICheckInOut";
import { DataGrid, GridApi, GridCellValue, GridColDef } from "@mui/x-data-grid";
import { GetCheckInOut,DeleteCheckInOut, CheckOut } from "./service/CheckInOutHttpClientService";
import moment from "moment";
import { error } from "console";

//color
import { createTheme, ThemeProvider } from '@mui/material/styles';
import { purple } from '@mui/material/colors';
import id from "date-fns/esm/locale/id/index.js";

function CheckInOutShow() {
  const [checkInOuts, setCheckInOuts] = useState<CheckInOutInterface[]>([]);
  //const [deleteIO, setDeleteIO] = useState<CheckInOutInterface[]>([]);
  const [success, setSuccess] = useState(false);
  const [successDel, setSuccessDel] = useState(false);
  const [error, setError] = useState(false);
  const [errorDel, setErrorDel] = useState(false);
  const [id, setId] = useState(0);
  //let strid:string = ""
  

  // const theme = createTheme({
  //   palette: {
  //     primary: {
  //       // Purple and green play nicely together.
  //       main: purple[500],
  //     },
  //     secondary: {
  //       // This is green.A700 as hex.
  //       main: '#11cb5f',
  //     },
  //     custom: {
  //       main: '#d32f2f'
  //     },
  //   },
  // });

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
    let res = await GetCheckInOut();
    if (res) {
      setCheckInOuts(res);
    }
  };

  const onDelete = async (id: number) => {
    let res = await DeleteCheckInOut(id);
    if (res) {
      setSuccessDel(true);
    } else {
      setErrorDel(true);
    }
    getList()
  }

  const onCheckOut = async (id: number) => {
    let res = await CheckOut(id);
    if (res) {
      setSuccess(true);
    } else {
      setError(true);
    }
    getList()
  }

  function getStringValue(value: any): string {
    return value.toString();
  }

  function formatTime (time : string){
    const data = new Date(time)
    console.log(data.getFullYear())
    if(String(data.getFullYear()) === "1"){
      return 
    }else{
      return moment(time).format('DD-MM-yyyy เวลา hh:mm')
    }
  }
  const columns: GridColDef[] = [
    { field: "ID", headerName: "ลำดับ", width: 100 },
    { field: "CheckInTime", headerName: "Check-In Time", width: 180, valueFormatter: (params) => moment(params.value).format('DD-MM-yyyy เวลา hh:mm')},
    { field: "CheckOutTime", headerName: "Check-Out Time", width: 180, valueFormatter: (params) => formatTime(params.value)},
    { field: "Booking", headerName: "Booking ID", width: 120, valueFormatter: (params) => params.value.ID},
    //{ field: "Booking_Name", headerName: "Customer Name", width: 120, valueFormatter: (params) => params.value.Name},
    { field: "CheckInOutStatus", headerName: "Status", width: 130, valueFormatter: (params) => params.value.Name,},
    { field: "Employee", headerName: "Employee", width: 110, valueFormatter: (params) => params.value.Eusername,},
    // { field: "Delete", headerName: "Delete", width: 80, valueFormatter: (params) => <Button onClick={() => deleteCheckInOut(params.value.ID)}>Edit</Button> },
    // { field: "Delete",
    //   renderCell: (cellValues) => {
    //     function handleClick(e: React.MouseEvent<HTMLButtonElement, MouseEvent>, cellValues: GridRenderCellParams<any, any, any>) {
    //       throw new Error("Function not implemented.");
    //     }

    //     return (
    //       <Button
    //         variant="contained"
    //         color="primary"
    //         onClick={(e) =>{
    //           handleClick(e, cellValues);
    //         }}
    //       >
    //       Delete
    //       </Button>
    //     )
    //   }}
    // {
    //   field: "test",
    //   headerName: "test",
    //   sortable: false,
    //   renderCell: (params) => {
    //     const onClick = (e: { stopPropagation: () => void; }) => {
    //       e.stopPropagation(); // don't select this row after clicking
  
    //       const api: GridApi = params.api;
    //       const thisRow: Record<string, GridCellValue> = {};
  
    //       api
    //         .getAllColumns()
    //         .filter((c) => c.field !== "__check__" && !!c)
    //         .forEach(
    //           (c) => (thisRow[c.field] = params.getValue(params.id, c.field))
    //         );
  
    //       return alert(JSON.stringify(thisRow, null, 4));
    //     };
  
    //     return <Button onClick={onClick}>CO</Button>;
    //   }
    // },
    {
      field: "checkout",
      headerName: "CHECK-OUT",
      width: 150,
      sortable: false,
      align:"center",
      renderCell: (params) => {
          const onClick = (e: { stopPropagation: () => void; }) => {
              e.stopPropagation();
              const id = params.getValue(params.id, "ID");
              onCheckOut(id);
          };

          return <Button onClick={onClick} color="success" endIcon={<CheckCircleIcon />}>CheckOut</Button>;
      }
    },
    {
      field: "delete",
      headerName: "DELETE",
      width: 150,
      sortable: false,
      align:"center",
      renderCell: (params) => {
          const onClick = (e: { stopPropagation: () => void; }) => {
              e.stopPropagation();
              const id = params.getValue(params.id, "ID");
              onDelete(id);
          };

          return <Button onClick={onClick} color="error" endIcon={<DeleteOutlineIcon />} >Delete</Button>;
      }
    },
    // {
    //   field: "edit",
    //   headerName: "EDIT",
    //   sortable: false,
    //   align:"center",
    //   renderCell: (params) => {
    //       // const onClick = (e: { stopPropagation: () => void; }) => {
    //       //     e.stopPropagation();
    //       //     const id = params.getValue(params.id, "ID");
    //       //     //onEdit(id);
    //       //     //alert(id)
    //       //     // let test = getStringValue(id)
    //       //     // alert(test)
              
    //       // };
    //       const onClick = (e: { stopPropagation: () => void; }) => {
    //         e.stopPropagation();
    //         setId(params.getValue(params.id, "ID"));
    //       };
        
    //      let strid = id.toString();
    //       return <Button onClick={onClick}  color="error" component={RouterLink} to={"/employee/update/" + id } endIcon={<DeleteOutlineIcon />} >EDIT</Button>;
    //   }
    // },
    
    
  ];
  // const deleteCheckInOut = (id: number) => {
  //   alert(id)
  // }

  //<Button component={RouterLink} to="/user/create" variant="contained" color="primary"> สร้างข้อมูล </Button>
  useEffect(() => {
    getList();
  }, []);

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
            Check Out สำเร็จ
          </Alert>
        </Snackbar>
        <Snackbar
          open={error}
          autoHideDuration={6000}
          onClose={handleClose}
          anchorOrigin={{ vertical: "top", horizontal: "center" }}
        >
          <Alert onClose={handleClose} severity="error">
            ไม่สามารถ Check Out ได้
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
              color="primary"
              gutterBottom
            >
              CHECK IN & CHECK OUT
            </Typography>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/CNCO/Create"
              variant="contained"
              color="primary"
            >
              check in
            </Button>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/CNCO/Edit"
              variant="contained"
              color="primary"
            >
              Edit
            </Button>
          </Box>
        </Box>
        <div style={{ height: 400, width: "100%", marginTop: "20px" }}>
          <DataGrid
            rows={checkInOuts}
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

export default CheckInOutShow;