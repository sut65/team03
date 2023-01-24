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
import { UsersInterface } from "../../models/IUser";
import {
  createTheme,
  FormControlLabel,
  FormLabel,
  MenuItem,
  Radio,
  RadioGroup,
  Select,
  SelectChangeEvent,
  ThemeProvider,
} from "@mui/material";
import {
  DepartmentInterface,
  EmployeeInterface,
  LocationInterface,
  OfficerInterface,
  PositionInterface,
} from "../../models/IEmployee";
import { grey } from "@mui/material/colors";

const bgbutton = createTheme({
  palette: {
    primary: {
      // Purple and grey play nicely together.
      main: grey[800],
    },
    secondary: {
      // Purple and grey play nicely together.
      main: grey[50],
    },

  },
});

const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
  props,

  ref
) {
  return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function Manage_Save() {
  const [date, setDate] = React.useState<Date | null>(null);
  const [employee, setEmployee] = React.useState<Partial<EmployeeInterface>>(
    {
      DepartmentID: 0,
      PositionID: 0,
      LocationID: 0,
    }
  );
  const [location, setLocation] = React.useState<LocationInterface[]>([]);
  const [department, setDepartment] = React.useState<DepartmentInterface[]>([]);
  const [position, setPosition] = React.useState<PositionInterface[]>([]);
  const [user, setUser] = React.useState<OfficerInterface>();
  const [gender, setGender] = React.useState<string>("");
  const [success, setSuccess] = React.useState(false);
  const [error, setError] = React.useState(false);
  const [offID,setOffID] = React.useState<Number | null>(0);
  const [dateOfBirth, setDateOfBirth] = React.useState<Date | null>(new Date());
  const [yearOfStart, setYearOfStart] = React.useState<Date | null>(new Date());


  //-----------เริ่มดึงข้อมูล-----------//
//---------------------Department-------------------------------------
const getDepartment = async () => {
  const apiUrl = `http://localhost:8080/Departments`;

  const requestOptions = {
    method: "GET",

    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
    },
  };
  //การกระทำ //json
  fetch(apiUrl, requestOptions)
    .then((response) => response.json()) //เรียกได้จะให้แสดงเป็น json ซึ่ง json คือ API

    .then((res) => {
      console.log(res.data); //show ข้อมูล

      if (res.data) {
        setDepartment(res.data);
      } else {
        console.log("else");
      }
    });
};
//---------------------Position-------------------------------------
const getPosition = async () => {
  const apiUrl = `http://localhost:8080/Positions`;

  const requestOptions = {
    method: "GET",

    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
    },
  };
  //การกระทำ //json
  fetch(apiUrl, requestOptions)
    .then((response) => response.json()) //เรียกได้จะให้แสดงเป็น json ซึ่ง json คือ API

    .then((res) => {
      console.log(res.data); //show ข้อมูล

      if (res.data) {
        setPosition(res.data);
      } else {
        console.log("else");
      }
    });
};
//---------------------Position-----------------------------
const getLocation = async () => {
  const apiUrl = `http://localhost:8080/Locations`;

  const requestOptions = {
    method: "GET",

    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
    },
  };
  //การกระทำ //json
  fetch(apiUrl, requestOptions)
    .then((response) => response.json()) //เรียกได้จะให้แสดงเป็น json ซึ่ง json คือ API

    .then((res) => {
      console.log(res.data); //show ข้อมูล

      if (res.data) {
        setLocation(res.data);
      } else {
        console.log("else");
      }
    });
};


// const getOfficer = async () => {
//   const apiUrl = `http://localhost:8080/Officers/`;

//   const requestOptions = {
//     method: "GET",

//     headers: {
//       Authorization: `Bearer ${localStorage.getItem("token")}`,
//       "Content-Type": "application/json",
//     },
//   };
//   //การกระทำ //json
//   fetch(apiUrl, requestOptions)
//     .then((response) => response.json()) //เรียกได้จะให้แสดงเป็น json ซึ่ง json คือ API

//     .then((res) => {
//       //console.log(res.data); //show ข้อมูล

//       if (res.data) {
//         setOfficer(res.data);
//       } else {
//         console.log("else");
//       }
//     });
// };
//----------------------------------จบการดึงข้อมูล------------------------
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

  const handleInputChange = (
    event: React.ChangeEvent<{ id?: string; value: any }>
  ) => {
    const id = event.target.id as keyof typeof Manage_Save;

    const { value } = event.target;

    setEmployee({ ...employee, [id]: value });
  };

  const handleChange = (event: SelectChangeEvent<number>) => {
    const name = event.target.name as keyof typeof employee;
    setEmployee({
      ...employee,
      [name]: event.target.value,
    });
  };

  function submit() {
    let data = {
      PersonalID: typeof employee.PersonalID === "string" ? parseInt(employee.PersonalID) : 0,
      Employeename:  employee.Employeename ,
      Email: employee.Email ,
      Eusername: employee.Eusername ?? "",
      Password: employee.Password ?? "",
      Salary: typeof employee.Salary === "string" ? parseInt(employee.Salary) : 0,
      Phonenumber: employee.Phonenumber ?? "",
      Gender: gender,
      DateOfBirth: dateOfBirth,
      YearOfStart: yearOfStart,
      Address: employee.Address ?? "",
      OfficerID: user?.ID ?? "",
      DepartmentID: employee.DepartmentID ,
      PositionID: employee.PositionID ,
      LocationID: employee.LocationID,
      Signin: {
        Username: employee.Eusername ?? "",
        Password: employee.Password ?? "",
      }
    };

    const apiUrl = "http://localhost:8080/Employees";

    const requestOptions = {
      method: "POST",

      headers:       
      {  
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json" 
      },

      body: JSON.stringify(data),
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
    const getToken = localStorage.getItem("token");
    if (getToken) {
      setUser(JSON.parse(localStorage.getItem("user") || ""));
    }
    getDepartment();
    getPosition();
    getLocation();
    // getOfficer();

  }, []);

  return (
  <ThemeProvider theme={bgbutton}>
    <Container maxWidth="xl" >
      <Snackbar
        open={success}
        autoHideDuration={6000}
        onClose={handleClose}
        anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
      >
        <Alert onClose={handleClose} severity="success">
          บันทึกข้อมูลสำเร็จ
        </Alert>
      </Snackbar>

      <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
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
              Record Employee Information
            </Typography>
          </Box>
        </Box>

        <Divider />

        <Grid container spacing={3} sx={{ padding: 2 }} style={{ marginLeft: "6.5%"}}>
          <Grid item xs={5}>
            <FormControl fullWidth variant="outlined">
              <FormLabel>Personal ID</FormLabel>

              <TextField
                id="PersonalID"
                variant="outlined"
                type="string"
                size="medium"
                value={employee.PersonalID || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>

          <Grid item xs={5}>
            <FormControl fullWidth variant="outlined">
              <FormLabel>Name</FormLabel>
              <TextField
                id="Employeename"
                variant="outlined"
                type="string"
                size="medium"
                value={employee.Employeename || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>
        </Grid>

        <Grid container spacing={3} sx={{ padding: 2 }} style={{marginLeft: "10.5%"}}>
          {/* ComboboxDepartment */}
          <Grid item xs={3}>
            <FormLabel>Department</FormLabel>
            <FormControl fullWidth variant="outlined">
              <Select
                value={employee.DepartmentID}
                onChange={handleChange}
                inputProps={{
                  name: "DepartmentID",
                }}
              >
                <MenuItem value={0} key={0}>
                  เลือกแผนก
                </MenuItem>
                {department.map((item: DepartmentInterface) => (
                  <MenuItem value={item.ID}>{item.Name}</MenuItem>
                ))}
              </Select>
            </FormControl>
          </Grid>

          {/* ComboboxPosition */}
          <Grid item xs={3}>
            <FormLabel>Position</FormLabel>
            <FormControl fullWidth variant="outlined">
              <Select
                value={employee.PositionID}
                onChange={handleChange}
                inputProps={{
                  name: "PositionID",
                }}
              >
                <MenuItem value={0} key={0}>
                  เลือกตำแหน่ง
                </MenuItem>
                {position.map((item: PositionInterface) => (
                  <MenuItem value={item.ID}>{item.Name}</MenuItem>
                ))}
              </Select>
            </FormControl>
          </Grid>

          {/* ComboboxLocation */}
          <Grid item xs={3}>
            <FormLabel>Location</FormLabel>

            <FormControl fullWidth variant="outlined">
              <Select
                value={employee.LocationID}
                onChange={handleChange}
                inputProps={{
                  name: "LocationID",
                }}
              >
                <MenuItem value={0} key={0}>
                  เลือกสถานที่
                </MenuItem>
                {location.map((item: LocationInterface) => (
                  <MenuItem value={item.ID}>{item.Name}</MenuItem>
                ))}
              </Select>
            </FormControl>
          </Grid>
        </Grid>

        <Grid container spacing={3} sx={{ padding: 2 }} style={{marginLeft: "14.5%"}}>
          <Grid item xs={4}>
            <FormControl fullWidth variant="outlined">
              <FormLabel>Username</FormLabel>

              <TextField
                id="Eusername"
                variant="outlined"
                type="string"
                size="medium"
                value={employee.Eusername || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>

          <Grid item xs={4}>
            <FormControl fullWidth variant="outlined">
              <FormLabel>Password</FormLabel>

              <TextField
                id="Password"
                variant="outlined"
                type="string"
                size="medium"
                value={employee.Password || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>
        </Grid>

        <Grid container spacing={3} sx={{ padding: 2 }} style={{marginLeft: "14.5%"}}>
          <Grid item xs={4}>
            <FormControl fullWidth variant="outlined">
              <FormLabel>Email</FormLabel>

              <TextField
                id="Email"
                variant="outlined"
                type="string"
                size="medium"
                value={employee.Email || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>

          <Grid item xs={4}>
            <FormControl fullWidth variant="outlined">
              <FormLabel>Tel</FormLabel>

              <TextField
                id="Phonenumber"
                variant="outlined"
                type="string"
                size="medium"
                value={employee.Phonenumber || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>
        </Grid>

        <Grid container spacing={3} sx={{ padding: 2 }} style={{marginLeft: "14.5%"}}>
          <Grid item xs={4}>
            <FormControl fullWidth variant="outlined">
              <FormLabel>Salary</FormLabel>

              <TextField
                id="Salary"
                variant="outlined"
                type="string"
                size="medium"
                value={employee.Salary || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>
        </Grid>

        <Grid container spacing={3} sx={{ padding: 2 }} style={{marginLeft: "14.5%"}}>
          <Grid item xs={6}>
            <FormControl>
              <FormLabel>Gender</FormLabel>
              <RadioGroup
                row
                aria-labelledby="demo-row-radio-buttons-group-label"
                name="row-radio-buttons-group"
                onChange={(event) => {
                  setGender(event.target.value);
                }}
              >
                <FormControlLabel
                  value="Female"
                  control={<Radio />}
                  label="Female"
                />
                <FormControlLabel
                  value="Male"
                  control={<Radio />}
                  label="Male"
                />
              </RadioGroup>
            </FormControl>
          </Grid>
        </Grid>

        <Grid container spacing={3} sx={{ padding: 2 }}>
          <Grid item xs={12}>
            <FormControl fullWidth variant="outlined">
              <FormLabel>Address</FormLabel>

              <TextField
                id="Address"
                variant="outlined"
                type="string"
                size="medium"
                value={employee.Address || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>
        </Grid>

        {/* วันเกิด */}
      <Grid container spacing={3} sx={{ padding: 2 }} style={{marginLeft: "14.5%"}}>
        <Grid item xs={4}>
          <FormControl fullWidth variant="outlined">
            <FormLabel>BirthDay</FormLabel>
            <LocalizationProvider dateAdapter={AdapterDateFns}>
              <DatePicker
                value={dateOfBirth}
                onChange={setDateOfBirth}
                renderInput={(params) => <TextField {...params} />}
              />
            </LocalizationProvider>
          </FormControl>
        </Grid>

        {/* วันที่ทำงาน */}

        <Grid item xs={4}>
          <FormControl fullWidth variant="outlined">
            <FormLabel>Start Date</FormLabel>

            <LocalizationProvider dateAdapter={AdapterDateFns}>
              <DatePicker
                value={yearOfStart}
                onChange={setYearOfStart}
                renderInput={(params) => <TextField {...params} />}
              />
            </LocalizationProvider>
          </FormControl>
        </Grid>
      </Grid>

        <Grid container spacing={3} sx={{ padding: 2 }}>
          <Grid item xs={4}>
            <FormControl fullWidth variant="outlined">
              <FormLabel>Officer</FormLabel>

              <TextField
                fullWidth
                disabled
                id="OfficerID"
                value={user?.Officername}
                //value={officer?.ID || ""}
                // onChange={(event) => setOffID(Number(event.target.value)) }
              />
            </FormControl>
          </Grid>

          <Grid item xs={12}>
            <Button component={RouterLink} to="/ManageShow" variant="contained">
              Back
            </Button>

            <Button
              style={{ float: "right" }}
              onClick={submit}
              variant="contained"
              color="primary"
            >
              Submit
            </Button>
          </Grid>
        </Grid>
      </Paper>
    </Container>
</ThemeProvider>
  );
}

export default Manage_Save;