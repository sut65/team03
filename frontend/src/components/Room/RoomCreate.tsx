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


import { EmployeeInterface } from "../../models/IEmployee"; 

import { RoomInterface, RoomTypeInterface, RoomZoneInterface, StateInterface } from "../../models/IRoom";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import { DesktopDateTimePicker } from "@mui/x-date-pickers";
import { InputLabel, Stack } from "@mui/material";
import { GetEmployees, GetRoomTypes, GetRoomZones, GetStates, CreateRoom } from "./service/RoomHttpClientService";

import { grey } from '@mui/material/colors';
import { Message } from "@mui/icons-material";


const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
  props,
  ref
) {
  return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function RoomCreate() {
  const [roomtypes, setRoomTypes] = useState<RoomTypeInterface[]>([]);
  const [states, setStates] = useState<StateInterface[]>([]);
  const [employees, setEmployees] = useState<EmployeeInterface[]>([]);
  const [roomzone, setRoomZones] = useState<RoomZoneInterface[]>([]);

  const [room, setRoom] = useState<RoomInterface>({
    Time: new Date(),
    Amount: 0,});
  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);
  const [message, setAlertMessage] = useState("");
  const [number, setNumber] = useState("");

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
    const name = event.target.name as keyof typeof room;
    setRoom({
      ...room,
      [name]: event.target.value,
    });
  }; 

  const handleInputChangenumber = (

    event: React.ChangeEvent<{ id?: string; value: any }>

  ) => {

    const id = event.target.id as keyof typeof room;

    const { value } = event.target;

    setRoom({ ...room, [id]: value });

  };

  // const handleInputChangenum = (

  //   event: React.ChangeEvent<{ id?: string; value: any }>

  // ) => {

  //   const id = event.target.id as keyof typeof room;

  //   const { value } = event.target;

  //   setRoom({ ...room, [id]: value  === "" ? "" : Number(value)  });
  //   console.log(`id: ${id} value:${value}`)

  // };

  const getRoomType =  async () => {
    let res = await GetRoomTypes();
    if (res) {
      setRoomTypes(res);
    }
  };

  const getRoomZone =  async () => {
    let res = await GetRoomZones();
    if (res) {
      setRoomZones(res);
    }
  };

  const getState =  async () => {
    let res = await GetStates();
    if (res) {
      setStates(res);
    }
  };

  // const getEmployee =  async () => {
  //   let res = await GetEmployees();
  //   if (res) {
  //     setEmployees(res);
  //   }
  // };

  useEffect(() => {
    getRoomType();
    getRoomZone();
    //getEmployee();
    getState();
  }, []);

  const convertType = (data: string | number | undefined | null) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  async function submit() {
    let data = {
      RoomTypeID: convertType(room.RoomTypeID),
      RoomZoneID: convertType(room.RoomZoneID),
      StateID: convertType(room.StateID),
      EmployeeID: convertType(localStorage.getItem("id")),
      Room_No: room.Room_No,
      Amount: convertType(number)||0,
      //Amount: convertType(room.Amount),
      Time: room.Time,
    };

    let res = await CreateRoom(data);
    if (res.status) {
      setAlertMessage("บันทึกข้อมูลสำเร็จ");
      setSuccess(true);
    } else {
      setAlertMessage(res.message);
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
          {/* บันทึกข้อมูลสำเร็จ */}
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
          {/* บันทึกข้อมูลไม่สำเร็จ */}
          {message}
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
              color="grey"
              gutterBottom
            >
              เพิ่มข้อมูลห้อง
            </Typography>
          </Box>
        </Box>
        <Divider />
        
        <Grid container spacing={3} sx={{ padding: 2 }}>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
            <TextField
          id="Room_No" label="หมายเลขห้อง" type="string" 
          InputLabelProps={{ shrink: true,}} value={room?.Room_No} 
          onChange={handleInputChangenumber}    
          inputProps={{name: "Room_No"}}    
          />
            </FormControl>
        </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
            <InputLabel id="demo-simple-select-label">ประเภทของห้องพัก</InputLabel>
              <Select
                labelId="demo-simple-select-label"
                native
                value={room.RoomTypeID + ""}
                label="ประเภทของห้องพัก."
                onChange={handleChange}
                inputProps={{
                  name: "RoomTypeID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกประเภทห้องพัก
                </option>
                {roomtypes.map((item: RoomTypeInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Size}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
            <InputLabel id="demo-simple-select-label">โซนของห้องพัก</InputLabel>
              <Select
                labelId="demo-simple-select-label"
                native
                value={room.RoomZoneID + ""}
                label="โซนของห้องพัก."
                onChange={handleChange}
                inputProps={{
                  name: "RoomZoneID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกโซนห้องพัก
                </option>
                {roomzone.map((item: RoomZoneInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
            <InputLabel id="demo-simple-select-label">สถานะของห้องพัก</InputLabel>
              <Select
                labelId="demo-simple-select-label"
                native
                value={room.StateID + ""}
                label="สถานะของห้องพัก."
                onChange={handleChange}
                inputProps={{
                  name: "StateID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกสถานะห้อง
                </option>
                {states.map((item: StateInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={6}>
          <FormControl fullWidth variant="outlined">
          <TextField
              type="string"
              value={number}
              id="Amount"
              label="ราคาห้องพัก"
              variant="outlined"
              inputProps={{
                  name: "Amount",
                  // min: 0
              }}
              onChange={(e)=>setNumber(e.target.value)} 
          />
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
                <LocalizationProvider dateAdapter={AdapterDayjs}>
                    <Stack spacing={3}>
                        <DesktopDateTimePicker
                        label="วันที่และเวลาที่บันทึกข้อมูล"
                        value={room.Time}
                        onChange={(newValue) => {
                            setRoom({
                                ...room,
                                Time: newValue,
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
              to="/RT"
              variant="contained"
              color="inherit"
            >
              กลับ
            </Button>
            <Button
              style={{ float: "right" }}
              onClick={submit}
              variant="contained"
              color="success"
            >
              บันทึก
            </Button>
          </Grid>
        </Grid>
      </Paper>
    </Container>
  );

}
export default RoomCreate;