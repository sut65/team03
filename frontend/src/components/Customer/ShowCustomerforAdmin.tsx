import React, { useState, useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { Alert, Snackbar } from "@mui/material";
import DeleteOutlineIcon from '@mui/icons-material/DeleteOutline';
import { CustomerInterface } from "../../models/modelCustomer/ICustomer"; 
import { DataGrid, GridApi, GridColDef } from "@mui/x-data-grid";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import { grey } from '@mui/material/colors';
import moment from "moment";
import {GetCustomerlist, DeleteCustomer}  from "./service/servicecus";


function CustomerlistforAdmin() {
    const [customer, setCustomer] = useState<CustomerInterface[]>([]);
    const [success, setSuccess] = useState(false);
    const [successDelete, setSuccessDelete] = useState(false);
    const [error, setError] = useState(false);
    const [errorDelete, setErrorDelete] = useState(false);
  


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

    const handleClose = (
      event?: React.SyntheticEvent | Event,
      reason?: string
    ) => {
      if (reason === "clickaway") {
        return;
      }
      setSuccess(false);
      setSuccessDelete(false);
      setError(false);
      setErrorDelete(false);
    };
  
    const getCustomerlist = async () => {
      let res = await GetCustomerlist();
      if (res) {
          setCustomer(res);
          console.log(res);
      }
    };
    
    const onDelete = async (id: number) => {
      let res = await DeleteCustomer(id);
      if (res) {
        setSuccessDelete(true);
        window.location.reload();
      } else {
        setErrorDelete(true);
      }
      GetCustomerlist()
    }
   
    const columns: GridColDef[] = [
        { field: "ID", headerName: "ลำดับ", width: 100 },
        { field: "Nametitle", headerName: "คำนำหน้า", width: 150 , valueFormatter: (params) => params?.value?.NT_Name,},
        { field: "FirstName", headerName: "ชื่อ", width: 100 },
        { field: "LastName", headerName: "สกุล", width: 100 },
        { field: "Age", headerName: "อายุ", width: 100 },
        { field: "Phone", headerName: "เบอร์โทรศัพท์", width: 150 },
        { field: "Email", headerName: "อีเมล", width: 200 },
        { field: "Gender", headerName: "เพศ", width: 100 , valueFormatter: (params) => params?.value?.G_Name,},
        { field: "Province", headerName: "จังหวัดที่อยู่", width: 150 , valueFormatter: (params) => params?.value?.P_Name,},
        { field: "delete",
          headerName: "ลบข้อมูลสมาชิก",
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
     
      ];
      useEffect(() => {
        getCustomerlist();
      }, []);
  
    return (
        <div>
      <Container maxWidth="lg">
        <Snackbar
          open={successDelete}
          autoHideDuration={3000}
          onClose={handleClose}
          anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
        >
          <Alert onClose={handleClose} severity="success">
            ลบข้อมูลสำเร็จ
          </Alert>
        </Snackbar>
        <Snackbar
          open={errorDelete}
          autoHideDuration={6000}
          onClose={handleClose}
          anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
        >
          <Alert onClose={handleClose} severity="error">
            ไม่สามารถลบข้อมูลได้
          </Alert>
        </Snackbar>
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
               ข้อมูลสมาชิก
              </Typography>
            </Box>
           </Box>
   
          <div style={{ height: 400, width: "100%", marginTop: '20px'}}>
   
            <DataGrid
              rows={customer}
              getRowId={(row) => row.ID}
              columns={columns}
              pageSize={5}
              rowsPerPageOptions={[5]}
            />
          </div>
        </Container>
        </ThemeProvider>
      </Container>
      </div>
    );
  }
  
  export default CustomerlistforAdmin;