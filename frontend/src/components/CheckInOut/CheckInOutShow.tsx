import React, { useState, useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";

import { CheckInOutInterface, ChecKInOutStatusInterface } from "../../models/ICheckInOut";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import { GetListCheckInOut } from "./service/CheckInOutHttpClientService";
import moment from "moment";

function CheckInOutShow() {
  const [checkInOuts, setCheckInOuts] = useState<CheckInOutInterface[]>([]);

  const getList = async () => {
    let res = await GetListCheckInOut();
    if (res) {
      setCheckInOuts(res);
    }
  };

  const columns: GridColDef[] = [
    { field: "ID", headerName: "ลำดับ", width: 60 },
    { field: "CheckInTime", headerName: "Check-In Time", width: 200, valueFormatter: (params) => moment(params.value).format('DD-MM-yyyy เวลา hh:mm:ss')},
    { field: "CheckOutTime", headerName: "Check-Out Time", width: 200, valueFormatter: (params) => moment(params.value).format('DD-MM-yyyy เวลา hh:mm:ss')},
    { field: "Booking", headerName: "Booking ID", width: 120, valueFormatter: (params) => params.value.ID},
    { field: "CheckInOutStatus", headerName: "Status", width: 120, valueFormatter: (params) => params.value.Name,},
    { field: "Employee", headerName: "Employee", width: 120, valueFormatter: (params) => params.value.Eusername,},
  ];

  useEffect(() => {
    getList();
  }, []);

  return (
    <div>
      <Container maxWidth="md">
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
              Check In/Out
            </Typography>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/user/create"
              variant="contained"
              color="primary"
            >
              สร้างข้อมูล
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