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

import { GetBookings } from "../Booking/services/BookingHttpClientService";
import { 
  GetIOStatus,
  GetCheckInOut,
  UpdateCheckIn,
  UpdateCheckOut,
 } from "./service/CheckInOutHttpClientService";

import { Stack } from "@mui/material";
import { DesktopDateTimePicker } from "@mui/x-date-pickers";
import Tabs from '@mui/material/Tabs';
import Tab from '@mui/material/Tab';

interface TabPanelProps {
  children?: React.ReactNode;
  index: number;
  value: number;
}


const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
  props,
  ref
) {
  return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function TabPanel(props: TabPanelProps) {
    const { children, value, index, ...other } = props;
    return (
        <div
        role="tabpanel"
        hidden={value !== index}
        id={`simple-tabpanel-${index}`}
        aria-labelledby={`simple-tab-${index}`}
        {...other}
        >
        {value === index && (
            <Box sx={{ p: 3 }}>
            <Typography>{children}</Typography>
            </Box>
        )}
        </div>
    );
}

function a11yProps(index: number) {
  return {
    id: `simple-tab-${index}`,
    'aria-controls': `simple-tabpanel-${index}`,
  };
}

function CheckInOutEdit() {
  const [checkinout, setCheckinout] = useState<CheckInOutInterface>({});
  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);
  const { id } = useParams();
  const [message, setAlertMessage] = useState("");
  const [value, setValue] = React.useState(0);

  const handleChangeTab = (event: React.SyntheticEvent, newValue: number) => {
    setValue(newValue);
  };

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
    if (res.status) {
      setAlertMessage("Edit Check In Success")
      setSuccess(true);
      setInterval(() => {
        window.location.assign("/CNCO");
    }, 2000);
    } else {
      setAlertMessage(res.message);
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
    if (res.status) {
      setAlertMessage("Edit Check Out Success")
      setSuccess(true);
      setInterval(() => {
        window.location.assign("/CNCO");
    }, 2000);
    } else {
      setAlertMessage(res.message);
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
         {message}
        </Alert>
      </Snackbar>
      <Snackbar
        open={error}
        autoHideDuration={6000}
        onClose={handleClose}
        anchorOrigin={{ vertical: "top", horizontal: "center" }}
      >
        <Alert onClose={handleClose} severity="error">
          {message}
        </Alert>
      </Snackbar>
      <Divider />
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
              CHECK IN/OUT EDIT
            </Typography>
          </Box>
        </Box>
      <Box sx={{ width: '100%' }}>
            <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
                <Tabs value={value} onChange={handleChangeTab} aria-label="basic tabs example">
                <Tab label="Check In" {...a11yProps(0)} />
                <Tab label="Check Out" {...a11yProps(1)} />
                </Tabs>
            </Box>
                <TabPanel value={value} index={0}>
                    <Grid container spacing={3} sx={{ padding: 2 }}>
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
                            component={RouterLink}
                            to="/CNCO"
                            variant="contained"
                            color="inherit"
                            >
                            BACK
                            </Button>
                        <Button
                        style={{ float: "right" }}
                        onClick={submitCheckIn}
                        variant="contained"
                        color="primary"
                        >
                        SAVE
                        </Button>
                    </Grid>
                    </Grid>
                </TabPanel>
                <TabPanel value={value} index={1}>
                    <Grid container spacing={3} sx={{ padding: 2 }}>
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
                            BACK
                            </Button>
                            <Button
                            style={{ float: "right" }}
                            onClick={submitCheckOut}
                            variant="contained"
                            color="primary"
                            >
                            SAVE
                            </Button>
                        </Grid>
                        </Grid>
                </TabPanel>
            </Box>
      </Paper>
    </Container>
  );

}
export default CheckInOutEdit;