import React, { useState, useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { Alert, Snackbar } from "@mui/material";
import DeleteOutlineIcon from '@mui/icons-material/DeleteOutline';
import CreateOutlinedIcon from '@mui/icons-material/CreateOutlined';

import { RepairReqInterface } from "../../models/IRepairReq";
import { DataGrid, GridApi, GridCellValue, GridColDef, GridRowParams } from "@mui/x-data-grid";
import { GetRepairReqs,DeleteRepairReq, GetRepairReqByCID } from "./service/RepRqHttpClientService";
import moment from "moment";
import { error } from "console";

//color
import { createTheme, ThemeProvider } from '@mui/material/styles';
import { purple } from '@mui/material/colors';
import id from "date-fns/esm/locale/id/index.js";

function RepRqShow() {
  const [rprq, setRpRq] = useState<RepairReqInterface[]>([]);
  const [success, setSuccess] = useState(false);
  const [successDel, setSuccessDel] = useState(false);
  const [error, setError] = useState(false);
  const [errorDel, setErrorDel] = useState(false);
  const id_cus = localStorage.getItem("id");

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
    let res = await GetRepairReqs();
    if (res) {
        setRpRq(res);
    }
  };

  const getRepReqByCID = async () => {
    let res = await GetRepairReqByCID(id_cus + "");
    if (res) {
        setRpRq(res);
    }
  };

  const onDelete = async (id: number) => {
    let res = await DeleteRepairReq(id);
    if (res) {
      setSuccessDel(true);
    } else {
      setErrorDel(true);
    }
    //getRepReqByID()
    // getList()
    getRepReqByCID()
  }

  const columns: GridColDef[] = [
    { field: "ID", headerName: "No.", width: 80 },
    { field: "Room", headerName: "Room", width: 110, valueFormatter: (params) => params.value.ID},
    { field: "RepairType", headerName: "Type", width: 170, valueFormatter: (params) => params.value.Name,},
    { field: "Note", headerName: "Note", width: 300, valueFormatter: (params) => params.value.Note},
    { field: "Time", headerName: "Time", width: 180, valueFormatter: (params) => moment(params.value).format('DD-MM-yyyy เวลา hh:mm')},
    { field: "Customer", headerName: "Customer", width: 100, valueFormatter: (params) => params.value.FirstName},
    {
      field: "delete",
      headerName: "",
      sortable: true,
      width: 120,
      align:"center",
      headerAlign: "center",
      renderCell: ({ row }: Partial<GridRowParams>) =>
          <Button 
              //to="/prescription/edit"
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
    {
      field: "edit",
      headerName: " ",
      sortable: true,
      width: 90,
      align:"center",
      headerAlign: "center",
      renderCell: ({ row }: Partial<GridRowParams>) =>
          <Button component={RouterLink} 
              to={`/Rep/Edit/${row.ID}`}
              size="small"
              //variant="contained"
              color="warning"
              sx={{borderRadius: 20,'&:hover': {color: '#FC0000', backgroundColor: '#F9EBEB'}}}
              endIcon={<CreateOutlinedIcon />}
          >
              Edit
          </Button>,
    },
  ];

  useEffect(() => {
    //getRepReqByID()
    // getList()
    getRepReqByCID()
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
              REPAIR REQUEST
            </Typography>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/Rep/Create"
              variant="contained"
              color="primary"
            >
              Request
            </Button>
          </Box>
          {/* <Box>
            <Button
              component={RouterLink}
              to="/Rep/Edit"
              variant="contained"
              color="primary"
            >
              Edit
            </Button>
          </Box> */}
        </Box>
        <div style={{ height: 400, width: "100%", marginTop: "20px" }}>
          <DataGrid
            rows={rprq}
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

export default RepRqShow;