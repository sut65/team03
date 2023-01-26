import React, { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import Button from "@mui/material/Button";
import FormControl from "@mui/material/FormControl";
import Container from "@mui/material/Container";
import Paper from "@mui/material/Paper";
import Grid from "@mui/material/Grid";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import Divider from "@mui/material/Divider";
import Snackbar from "@mui/material/Snackbar";
import Select, { SelectChangeEvent } from "@mui/material/Select";
import MuiAlert, { AlertProps } from "@mui/material/Alert";
import TextField from "@mui/material/TextField";
import { AdapterDateFns } from "@mui/x-date-pickers/AdapterDateFns";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { DatePicker } from "@mui/x-date-pickers/DatePicker";

import { BookingsInterface } from "../../models/IBooking";
import { 
  CheckInOutInterface,
  CheckInOutStatusInterface, 
} from "../../models/ICheckInOut";
import { EmployeeInterface } from "../../models/IEmployee"; 

import { GetBookings } from "../Booking/services/BookingHttpClientService";
import { 
  GetIOStatus,
  CreateCheckInOut,
  GetEmps,
 } from "./service/CheckInOutHttpClientService";

const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
  props,
  ref
) {
  return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function CheckInOutCreate() {
  const [bookings, setBookings] = useState<BookingsInterface[]>([]);
  const [statuses, setStatuses] = useState<CheckInOutStatusInterface[]>([]);
  const [emps, setEmps] = useState<EmployeeInterface[]>([]);
  const [checkinout, setCheckinout] = useState<CheckInOutInterface>({});

  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);

  const handleClose = (
    event?: React.SyntheticEvent | Event,
    reason?: string
  ) => {
    if (reason === "clickaway") {
      return;
    }
    setSuccess(false);
    setError(false);
  };

  const handleChange = (event: SelectChangeEvent) => {
    const name = event.target.name as keyof typeof checkinout;
    setCheckinout({
      ...checkinout,
      [name]: event.target.value,
    });
  }; 


  const getBookings =  async () => {
    let res = await GetBookings();
    if (res) {
      setBookings(res);
    }
  };

  const getStatuses =  async () => {
    let res = await GetIOStatus();
    if (res) {
      setStatuses(res);
    }
  };

  const getEmps =  async () => {
    let res = await GetEmps();
    if (res) {
      setEmps(res);
    }
  };

  useEffect(() => {
    getBookings();
    getStatuses();
    getEmps();
  }, []);

  const convertType = (data: string | number | undefined | null) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  async function submit() {
    let data = {
      BookingID: convertType(checkinout.BookingID),
      //CheckInOutStatusID: convertType("checkinout.CheckInOutStatusID"),
      CheckInOutStatusID: 1,
      //EmployeeID: convertType(checkinout.EmployeeID),
      EmployeeID: convertType(localStorage.getItem("id")),
      CheckInTime: checkinout.CheckInTime,
      CheckOutTime: null,
    };

    let res = await CreateCheckInOut(data);
    if (res) {
      setSuccess(true);
    } else {
      setError(true);
    }
  }

  return (
    <Container maxWidth="md">
      <Snackbar
        open={success}
        autoHideDuration={3000}
        onClose={handleClose}
        anchorOrigin={{ vertical: "top", horizontal: "center" }}
      >
        <Alert onClose={handleClose} severity="success">
          บันทึกข้อมูลสำเร็จ
        </Alert>
      </Snackbar>
      <Snackbar
        open={error}
        autoHideDuration={6000}
        onClose={handleClose}
        anchorOrigin={{ vertical: "top", horizontal: "center" }}
      >
        <Alert onClose={handleClose} severity="error">
          บันทึกข้อมูลไม่สำเร็จ
        </Alert>
      </Snackbar>
      <Paper>
        <Box
          display="flex"
          sx={{
            marginTop: 2,
          }}
        >
          <Box sx={{ paddingX: 2, paddingY: 1 }}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              บันทึกการเข้าชมวีดีโอ
            </Typography>
          </Box>
        </Box>
        <Divider />
        <Grid container spacing={3} sx={{ padding: 2 }}>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>booking</p>
              <Select
                native
                value={checkinout.BookingID + ""}
                onChange={handleChange}
                inputProps={{
                  name: "BookingID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือก booking
                </option>
                {bookings.map((item: BookingsInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.ID}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          {/* <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>status</p>
              <Select
                native
                value={checkinout.CheckInOutStatusID + ""}
                onChange={handleChange}
                inputProps={{
                  name: "CheckInOutStatusID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือก status
                </option>
                {statuses.map((item: CheckInOutStatusInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid> */}
          {/* <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>employee</p>
              <Select
                native
                value={checkinout.EmployeeID + ""}
                onChange={handleChange}
                inputProps={{
                  name: "EmployeeID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือก Employee
                </option>
                {emps.map((item: EmployeeInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Eusername}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid> */}
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>วันที่และเวลา</p>
              <LocalizationProvider dateAdapter={AdapterDateFns}>
                <DatePicker
                  value={checkinout.CheckInTime}
                  onChange={(newValue) => {
                    setCheckinout({
                      ...checkinout,
                      CheckInTime: newValue,
                    });
                  }}
                  renderInput={(params) => <TextField {...params} />}
                />
              </LocalizationProvider>
            </FormControl>
          </Grid>
          <Grid item xs={12}>
            <Button
              component={RouterLink}
              to="/CNCO"
              variant="contained"
              color="inherit"
            >
              กลับ
            </Button>
            <Button
              style={{ float: "right" }}
              onClick={submit}
              variant="contained"
              color="primary"
            >
              บันทึก
            </Button>
          </Grid>
        </Grid>
      </Paper>
    </Container>
  );

}
export default CheckInOutCreate;