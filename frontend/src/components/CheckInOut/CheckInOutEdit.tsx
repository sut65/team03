import React, { useEffect, useState } from "react";
import { Link as RouterLink, useParams } from "react-router-dom";
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
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import { AdapterDateFns } from "@mui/x-date-pickers/AdapterDateFns";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";

import { BookingsInterface } from "../../models/modelBooking/IBooking";
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
  GetCheckInOut,
  UpdateCheckInOut,
  UpdateCheckIn,
  UpdateCheckOut,
 } from "./service/CheckInOutHttpClientService";
import { DateTimePicker } from "@mui/x-date-pickers/DateTimePicker";
import { InputLabel, Stack } from "@mui/material";
import { DesktopDateTimePicker } from "@mui/x-date-pickers";

const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
  props,
  ref
) {
  return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function CheckInOutEdit() {
  const [bookings, setBookings] = useState<BookingsInterface[]>([]);
  const [statuses, setStatuses] = useState<CheckInOutStatusInterface[]>([]);
  //const [emps, setEmps] = useState<EmployeeInterface[]>([]);
  const [checkinout, setCheckinout] = useState<CheckInOutInterface>({});
  const [cio, setCio] = useState<CheckInOutInterface[]>([]);
  //const [date, setDate] = React.useState<Date | null>(null);
  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);
  const { id } = useParams();
  
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

  const getCheckInOut =  async () => {
    let res = await GetCheckInOut();
    if (res) {
      setCio(res);
    }
  };

  useEffect(() => {
    getBookings();
    getStatuses();
    getCheckInOut();
    //getEmps();
  }, []);

  const convertType = (data: string | number | undefined | null) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };
  
  const convertTypeNotNull = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  async function submitCheckIn() {
    let data = {
        ID: convertTypeNotNull(id),
        // ID: typeof checkinout.ID === "string" ? parseInt(checkinout.ID) : 0,
        //BookingID: convertType(checkinout.BookingID),
        BookingID: null,
        CheckInOutStatusID: null,
        //EmployeeID: convertType(checkinout.EmployeeID),
        EmployeeID: convertType(localStorage.getItem("id")),
        CheckInTime: checkinout.CheckInTime,
        CheckOutTime: null,
    };

    let res = await UpdateCheckIn(data);
    if (res) {
      setSuccess(true);
    } else {
      setError(true);
    }
  }

  async function submitCheckOut() {
    let data = {
        //ID: checkinout.ID,
        ID: convertTypeNotNull(id),
        BookingID: null,
        //BookingID: typeof checkinout.BookingID === "string" ? parseInt(checkinout.BookingID) : 0,
        CheckInOutStatusID: 2,
        //EmployeeID: convertType(checkinout.EmployeeID),
        EmployeeID: convertType(localStorage.getItem("id")),
        CheckInTime: null,
        CheckOutTime: checkinout.CheckOutTime,
    };

    let res = await UpdateCheckOut(data);
    if (res) {
      setSuccess(true);
    } else {
      setError(true);
    }
    console.log(data)
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
          อัพเดตข้อมูลสำเร็จ
        </Alert>
      </Snackbar>
      <Snackbar
        open={error}
        autoHideDuration={6000}
        onClose={handleClose}
        anchorOrigin={{ vertical: "top", horizontal: "center" }}
      >
        <Alert onClose={handleClose} severity="error">
          อัพเดตข้อมูลไม่สำเร็จ
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
              CHECK IN EDIT
            </Typography>
          </Box>
        </Box>
        <Divider />
        <Grid container spacing={3} sx={{ padding: 2 }}>
          {/* <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <InputLabel id="demo-simple-select-label">CheckInOut No.</InputLabel>
              <Select
                labelId="demo-simple-select-label"
                native
                value={checkinout.ID + ""}
                label="CheckInOut No."
                onChange={handleChange}
                inputProps={{
                  name: "ID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกหมายเลข CheckInOut
                </option>
                {cio.map((item: CheckInOutInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.ID}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid> */}
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
                <LocalizationProvider dateAdapter={AdapterDayjs}>
                    <Stack spacing={3}>
                        <DesktopDateTimePicker
                        label="Check-In Time"
                        value={checkinout.CheckInTime}
                        onChange={(newValue) => {
                            setCheckinout({
                                ...checkinout,
                                CheckInTime: newValue,
                              });
                        }}
                        renderInput={(params) => <TextField {...params} />}
                        />
                    </Stack>
                </LocalizationProvider>
            </FormControl>
          </Grid>
          <Grid item xs={12}>
            <Button
              style={{ float: "right" }}
              onClick={submitCheckIn}
              variant="contained"
              color="primary"
            >
              บันทึก
            </Button>
          </Grid>
        </Grid>
      </Paper>
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
              CHECK OUT EDIT
            </Typography>
          </Box>
        </Box>
        <Divider />
        <Grid container spacing={3} sx={{ padding: 2 }}>
          {/* <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <InputLabel id="demo-simple-select-label">CheckInOut No.</InputLabel>
              <Select
                labelId="demo-simple-select-label"
                native
                value={checkinout.ID + ""}
                label="CheckInOut No."
                onChange={handleChange}
                inputProps={{
                  name: "ID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกหมายเลข CheckInOut
                </option>
                {cio.map((item: CheckInOutInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.ID}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid> */}
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
                <LocalizationProvider dateAdapter={AdapterDayjs}>
                    <Stack spacing={3}>
                        <DesktopDateTimePicker
                        label="Check-Out Time"
                        value={checkinout.CheckOutTime}
                        onChange={(newValue) => {
                            setCheckinout({
                                ...checkinout,
                                CheckOutTime: newValue,
                              });
                        }}
                        renderInput={(params) => <TextField {...params} />}
                        />
                    </Stack>
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
              onClick={submitCheckOut}
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
export default CheckInOutEdit;