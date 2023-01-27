import React, { useEffect } from "react";

import { Link as RouterLink } from "react-router-dom";

import TextField from "@mui/material/TextField";

import Button from "@mui/material/Button";

import FormControl from "@mui/material/FormControl";

import Container from "@mui/material/Container";

import Paper from "@mui/material/Paper";

import Grid from "@mui/material/Grid";

import Box from "@mui/material/Box";

import Typography from "@mui/material/Typography";

import Divider from "@mui/material/Divider";

import Snackbar from "@mui/material/Snackbar";

import MuiAlert, { AlertProps } from "@mui/material/Alert";

import { AdapterDateFns } from "@mui/x-date-pickers/AdapterDateFns";

import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";

import { DatePicker } from "@mui/x-date-pickers/DatePicker";
//สี
import { createTheme, ThemeProvider } from "@mui/material/styles";
import { grey } from '@mui/material/colors';
//timedate
//import dayjs, { Dayjs } from "dayjs";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import { DateTimePicker } from "@mui/x-date-pickers/DateTimePicker";
//combobox
// import { ComboBoxComponent } from "@syncfusion/ej2-react-dropdowns";

import { RoomInterface, RoomTypeInterface, RoomZoneInterface, StateInterface } from "../../models/IRoom";

import Select, { SelectChangeEvent } from "@mui/material/Select";
import MenuItem from '@mui/material/MenuItem';
import { EmployeeInterface } from "../../models/IEmployee";
//import { getDecorators } from "typescript";


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

//เด้งขึ้นมาแจ้งเตือนว่าบันทึกสำเร็จ ไม่สำเร็จ 
const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
  props,

  ref
) {
  return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function RoomCreate() {

  //  const [date, setDate] = React.useState<Date | null>(null);

  const [room, setRoom] = React.useState<Partial<RoomInterface>>({
    EmployeeID: 0,
    RoomTypeID: 0,
    RoomZoneID: 0,
    StateID: 0,
  });


  //Partial คือเลือกค่า set ค่าได้เฉพาะตัวได้

  const [success, setSuccess] = React.useState(false);

  const [error, setError] = React.useState(false);

  const [roomtype, setroomtype] = React.useState<RoomTypeInterface[]>(
    []
  );
  const [employee, setEmployee] = React.useState<EmployeeInterface>(
  );

  //เราส่งมาในรูปแบบอาเรย์ ทำการดึงข้อมูล
  const [roomzone, setRoomZone] = React.useState<RoomZoneInterface[]>([]);
  const [state, setState] = React.useState<StateInterface[]>([]);
  const [selectedDate, setSelectedDate] = React.useState<Date | null>(
    new Date()
  );
  //const [user, setUser] = React.useState<DoctorsInterface>();

  const handleDateChange = (date: Date | null) => {
    setSelectedDate(date);
  };

  const getRoomZone = async () => {
    const apiUrl = `http://localhost:8080/room_zones`;

    const requestOptions = {
      method: "GET",

      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },
    };

    //การกระทำ //json
    fetch(apiUrl, requestOptions)
      .then((response) => response.json())

      .then((res) => {
        console.log(res.data); //show ข้อมูล

        if (res.data) {
          setRoomZone(res.data);
        } else {
          console.log("else");
        }
      });
  };
  //activity
  const getRoomType = async () => {
    const apiUrl = `http://localhost:8080/room_types`;

    const requestOptions = {
      method: "GET",

      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },
    };
    //การกระทำ
    fetch(apiUrl, requestOptions)
      .then((response) => response.json())

      .then((res) => {
        console.log(res.data);

        if (res.data) {
          setroomtype(res.data);
        } else {
          console.log("else");
        }
      });
  };
  //activity

   //activity
   const getState = async () => {
    const apiUrl = `http://localhost:8080/states`;

    const requestOptions = {
      method: "GET",

      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },
    };
    //การกระทำ
    fetch(apiUrl, requestOptions)
      .then((response) => response.json())

      .then((res) => {
        console.log(res.data);

        if (res.data) {
          setState(res.data);
        } else {
          console.log("else");
        }
      });
  };
  //activity

  const getEmployee = async () => {
    const uid = localStorage.getItem("uid")
    const apiUrl = `http://localhost:8080/employee/${uid}`;

    const requestOptions = {
      method: "GET",

      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },
    };
    //การกระทำ
    fetch(apiUrl, requestOptions)
      .then((response) => response.json())

      .then((res) => {
        console.log(res.data);

        if (res.data) {
          setEmployee(res.data);
          room.EmployeeID = res.data.ID
        } else {
          console.log("else");
        }
      });
  };
  //เปิดปิดตัว Alert
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

  console.log(room);

  //ทุกครั้งที่พิมพ์จะทำงานเป็น state เหมาะสำหรับกับคีย์ textfield
 
  
//กดเลือกคอมโบไม่ได้
  const handleChange = (
    event: SelectChangeEvent<number>
  ) => {
    const name = event.target.name as keyof typeof room;
    setRoom({
      ...room,
      [name]: event.target.value,
    });
  };

  const handleInputChange = (

    event: React.ChangeEvent<{ id?: string; value: any }>

  ) => {

    const id = event.target.id as keyof typeof room;

    const { value } = event.target;

    setRoom({ ...room, [id]: value });

  };


  const handleInputChangenumber = (

    event: React.ChangeEvent<{ id?: string; value: any }>

  ) => {

    const id = event.target.id as keyof typeof room;

    const { value } = event.target;

    setRoom({ ...room, [id]: value  === "" ? "" : Number(value)  });

  };

  function submit() {
    let data = {
      //แค่ข้างหน้า ชื่อต้องตรง!!!!!!!
      EmployeeID: room.EmployeeID,

      RoomTypeID: room.RoomTypeID,

      RoomZoneID: room.RoomZoneID,

      StateID: room.StateID,

      Time: selectedDate,
      
      // Num: typeof overtime?.Num === "string" ? (overtime?.Num === "" ? 0 : overtime?.Num) : overtime?.Num,
    
    };

    console.log(data)

    const apiUrl = "http://localhost:8080/rooms";

    const requestOptions = {
      method: "POST",

      headers: 
      {  Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json" 
      },

      body: JSON.stringify(data),
      //แปลงข้อมูล
    };

    fetch(apiUrl, requestOptions)
      .then((response) => response.json())

      .then((res) => {
        if (res.data) {
          setSuccess(true);
        } else {
          setError(true);
        }
      });
  }

  useEffect(() => {
    getEmployee();
    getRoomType();
    getRoomZone();
    getState();
  }, []);

  return (
    <ThemeProvider theme={theme}>
      <Container maxWidth="md">
        <Snackbar
          open={success}
          autoHideDuration={6000}
          onClose={handleClose}
          anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
        >
          <Alert onClose={handleClose} severity="success">
            Success
          </Alert>
        </Snackbar>

        <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
          <Alert onClose={handleClose} severity="error">
            Error
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
                ลบข้อมูลห้อง
              </Typography>
            </Box>
          </Box>

          <Divider />

          <Grid container spacing={3} sx={{ padding: 2 }}>
          <Grid item xs={6}>
                <p>หมายเลขห้อง</p>
                <TextField 
                fullWidth
                id="ID" InputProps={{inputProps: {min: 10}}} type="number" variant="outlined" value={room?.ID} onChange={handleInputChangenumber} 
                />
                  </Grid>
          {/*<Grid item xs={6}>
              <p>หมายเลขห้อง</p>
              <FormControl fullWidth variant="outlined"> 
                <TextField
                  value = {room?.ID || "" }
                   //onChange = {handleChange}
                  inputProps={{
                    // name: "DoctorID",
                    //readOnly: true
                  }}
                  // defaultValue={0}
                >
                     {/* <MenuItem value={0} key={0}>
                    กรุณาเลือกชื่อ
                  </MenuItem> */}
                  {/* <option aria-label="None" value="">
                  กรุณาเลือกความละอียด
                </option>
                {doctor.map((item: DoctorsInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Name}
                  </option>
                ))} */}
                  {/* {doctor.map((item: DoctorsInterface) => (
                    <MenuItem value={item.ID}>{item.Name}</MenuItem>
                  ))} */}
                {/*</TextField>
              </FormControl>
            </Grid>*/}
            {/*<Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>ประเภทของห้องพัก</p>
                <Select
                  value = {room.RoomTypeID}
                  onChange = {handleChange}
                  inputProps={{
                    size: "RoomTypeID",
                  }}
                  // defaultValue={0}
                >
                  <MenuItem value={0} key={0}>
                    กรุณาเลือกประเภทห้องพัก
                  </MenuItem>
                  {roomtype.map((item: RoomTypeInterface) => (
                    <MenuItem value={item.ID}>{item.Size}</MenuItem>
                  ))}
                </Select>
              </FormControl>
            </Grid>
            <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>สถานะของห้อง</p>
                <Select
                  value = {room.RoomTypeID}
                  onChange = {handleChange}
                  inputProps={{
                    name: "RoomTypeID",
                  }}
                  // defaultValue={0}
                >
                  <MenuItem value={0} key={0}>
                    กรุณาเลือกสถานะห้อง
                  </MenuItem>
                  {state.map((item: StateInterface) => (
                    <MenuItem value={item.ID}>{item.Name}</MenuItem>
                  ))}
                </Select>
              </FormControl>
            </Grid>
            {/*<Grid item xs={6}>
                <p>จำนวนชั่วโมงที่ทำ</p>
                <TextField 
                fullWidth
                id="Num" InputProps={{inputProps: {min: 1}}} type="number" variant="outlined" value={overtime?.Num} onChange={handleInputChangenumber} 
                />
                  </Grid>*/}
            {/*<Grid item xs={6}>
              <FormControl fullWidth variant="outlined">
                <p>โซนของห้องพัก</p>

                <Select
                  value={room.RoomZoneID}
                  onChange={handleChange}
                  inputProps={{
                    name: "RoomZoneID",
                  }}
                  // defaultValue={0}
                >
                  <MenuItem value={0} key={0}>
                    กรุณาเลือกโซนห้องพัก
                  </MenuItem>
                  {roomzone.map((item: RoomZoneInterface) => (
                    <MenuItem value={item.ID}>{item.Name}</MenuItem>
                  ))}
                </Select>
              </FormControl>
            </Grid>
            {/* //วันที่และเวลา */}
            {/*<Grid item xs={7}>
             <FormControl fullWidth variant="outlined">
              <p>วันที่และเวลา</p>
              <LocalizationProvider dateAdapter={AdapterDateFns}>
                <DatePicker
                  value={overtime.Time}
                  onChange={(newValue) => {
                    setOvertime({
                      ...overtime,
                      Time: newValue,
                    });
                  }}
                  renderInput={(params) => <TextField {...params} />}
                />
              </LocalizationProvider>
            </FormControl> */}
              {/*<FormControl fullWidth variant="outlined">
                <p>วันที่และเวลา</p>

                <LocalizationProvider dateAdapter={AdapterDayjs}>
                  <DateTimePicker
                    renderInput={(props) => <TextField {...props} />}
                    //label="กรุณาเลือกวันและเวลา"
                    value={selectedDate} //แก้
                    // onChange={(newValue) => {
                    // setDate(newValue);

                    // }}
                    onChange={setSelectedDate}
                  />
                </LocalizationProvider>
              </FormControl>
            </Grid>*/}

            <Grid item xs={12}>
              <Button component={RouterLink} to="/RT" variant="contained">
                <Typography
                  color="secondary"
                  component="div"
                  sx={{ flexGrow: 1 }}
                >
                  BACK
                </Typography>
              </Button>

              <Button
                style={{ float: "right" }}
                onClick={submit}
                variant="contained"
                color="primary"
              >
                <Typography
                  color="secondary"
                  component="div"
                  sx={{ flexGrow: 1 }}
                >
                 DELETE
                </Typography>
              </Button>
            </Grid>
          </Grid>
        </Paper>
      </Container>
    </ThemeProvider>
  );
}

export default RoomCreate;

