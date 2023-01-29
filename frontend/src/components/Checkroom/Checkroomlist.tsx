import React, { useState, useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { Alert, Snackbar } from "@mui/material";
import DeleteOutlineIcon from '@mui/icons-material/DeleteOutline';
import { CheckroomInterface } from "../../models/ICheckroom" 
import { DataGrid, GridApi, GridColDef } from "@mui/x-data-grid";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import { grey } from '@mui/material/colors';
import moment from "moment";
import {GetCheckroomslist}  from "./service/service";


function Checkroomlist() {
    const [checkroom, setCheckroom] = useState<CheckroomInterface[]>([]);
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
  
    const getCheckroomlist = async () => {
      let res = await GetCheckroomslist();
      if (res) {
          setCheckroom(res);
          console.log(res);
      }
    };
   
  
    // const onDelete = async (id: number) => {
    //   let res = await DeleteRepairReq(id);
    //   if (res) {
    //     setSuccessDel(true);
    //   } else {
    //     setErrorDel(true);
    //   }
    //   getList()
    // }
  
    const columns: GridColDef[] = [

        { field: "RoomID", headerName: "หมายเลขห้อง", width: 100 },
     
        { field: "Product", headerName: "อุปกรณ์", width: 150 , valueFormatter: (params) => params?.value?.Name,},
     
        { field: "Damage", headerName: "ความเสียหาย", width: 150 , valueFormatter: (params) => params?.value?.Description,},
     
        { field: "Status", headerName: "สถานะของห้อง", width: 150 , valueFormatter: (params) => params?.value?.S_Name,},
     
        { field: "Date", headerName: "วันที่และเวลา", width: 170, valueFormatter: (params) => moment(params.value).format('DD-MM-yyyy เวลา hh:mm') },

        { field: "Employee", headerName: "ชื่อ-นามสกุล", width: 150 , valueFormatter: (params) => params?.value?.Employeename,},
     
      ];
      useEffect(() => {
        getCheckroomlist();
      }, []);
   
     
   
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
   
               ข้อมูลการตรวจสอบห้องพัก
   
              </Typography>
   
            </Box>
   
            <Box>
            <Button //ตัวบันทึก
                   component={RouterLink} //ลิ้งหน้าต่อไป
                   to="/checkroom/create"
                   variant="contained"
                   color="primary"
                 >
                   <Typography
                     color="second"
                     component="div"
                     sx={{ flexGrow: 1 }}
                   >
                เพิ่มข้อมูล
   
                </Typography>
              </Button>
            </Box>
            <Box>
            <Button //ตัวบันทึก
                   component={RouterLink} //ลิ้งหน้าต่อไป
                   to="/"//รอเพิ่ม
                   variant="contained"
                   color="primary"
                 >
                   <Typography
                     color="second"
                     component="div"
                     sx={{ flexGrow: 1 }}
                   >
                แก้ไขข้อมูล
   
                </Typography>
              </Button>
            </Box>
            <Box>
            <Button //ตัวบันทึก
                   component={RouterLink} //ลิ้งหน้าต่อไป
                   to="/"//รอเพิ่ม
                   variant="contained"
                   color="primary"
                 >
                   <Typography
                     color="second"
                     component="div"
                     sx={{ flexGrow: 1 }}
                   >
                ลบข้อมูล
   
                </Typography>
              </Button>
            </Box>
   
   
   
          </Box>
   
          <div style={{ height: 400, width: "100%", marginTop: '20px'}}>
   
            <DataGrid
   
              rows={checkroom}
   
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
  
  export default Checkroomlist;