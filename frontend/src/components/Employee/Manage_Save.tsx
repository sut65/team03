import React from "react";

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
  FormControlLabel,
  FormLabel,
  MenuItem,
  Radio,
  RadioGroup,
  Select,
  SelectChangeEvent,
} from "@mui/material";
import {
  DepartmentInterface,
  EmployeeInterface,
  LocationInterface,
  OfficerInterface,
  PositionInterface,
} from "../../models/IEmployee";

const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
  props,

  ref
) {
  return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function Manage_Save() {
  const [date, setDate] = React.useState<Date | null>(null);

  const [user, setUser] = React.useState<Partial<UsersInterface>>({});

  const [employee, setEmployee] = React.useState<Partial<EmployeeInterface>>(
    {}
  );

  const [location, setLocation] = React.useState<LocationInterface[]>([]);

  const [department, setDepartment] = React.useState<DepartmentInterface[]>([]);

  const [position, setPosition] = React.useState<PositionInterface[]>([]);

  const [officer, setOfficer] = React.useState<OfficerInterface[]>([]);

  const [gender, setGender] = React.useState<string>("");

  const [success, setSuccess] = React.useState(false);

  const [error, setError] = React.useState(false);

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

    setUser({ ...user, [id]: value });
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
      FirstName: user.FirstName ?? "",

      LastName: user.LastName ?? "",

      Email: user.Email ?? "",

      Age: typeof user.Age === "string" ? parseInt(user.Age) : 0,

      BirthDay: date,
    };

    const apiUrl = "http://localhost:8080/users";

    const requestOptions = {
      method: "POST",

      headers: { "Content-Type": "application/json" },

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

  return (
    <Container maxWidth="xl"  >
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
              บันทึกข้อมูลพนักงาน
            </Typography>
          </Box>
        </Box>

        <Divider />

        <Grid container spacing={3} sx={{ padding: 2 }}>
          <Grid item xs={4} >
            <p>รหัสบัตรประชาชน</p>
            <FormControl fullWidth variant="outlined">
              <TextField
                id="FirstName"
                variant="outlined"
                type="string"
                size="medium"
                value={user.FirstName || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>

          <Grid item xs={4}>
            <FormControl fullWidth variant="outlined">
              <p>ชื่อ - นามสกุล</p>
              <TextField
                id="LastName"
                variant="outlined"
                type="string"
                size="medium"
                value={user.LastName || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>
        </Grid>

        <Grid container spacing={3} sx={{ padding: 2 }}>
          {/* Comboboxlocation */}
          <Grid item xs={3}>
            <p>แผนก</p>
            <FormControl fullWidth variant="outlined">
              <Select
                value={employee.LocationID}
                onChange={handleChange}
                inputProps={{
                  name: "LocationID",
                }}
              >
                <MenuItem value={0} key={0}>
                  เลือกแผนก
                </MenuItem>
                {location.map((item: LocationInterface) => (
                  <MenuItem value={item.ID}>{item.Name}</MenuItem>
                ))}
              </Select>
            </FormControl>
          </Grid>

          {/* Comboboxlocation */}
          <Grid item xs={3}>
            <p>ตำแหน่ง</p>
            <FormControl fullWidth variant="outlined">
              <Select
                value={employee.LocationID}
                onChange={handleChange}
                inputProps={{
                  name: "LocationID",
                }}
              >
                <MenuItem value={0} key={0}>
                  เลือกตำแหน่ง
                </MenuItem>
                {location.map((item: LocationInterface) => (
                  <MenuItem value={item.ID}>{item.Name}</MenuItem>
                ))}
              </Select>
            </FormControl>
          </Grid>

          {/* ComboboxDepartment */}
          <Grid item xs={3}>
            <p>สถานที่</p>

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
        

        <Grid container spacing={3} sx={{ padding: 2 }}>
          {/* วันเกิด */}
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>BirthDay</p>

              <LocalizationProvider dateAdapter={AdapterDateFns}>
                <DatePicker
                  value={date}
                  onChange={(newValue) => {
                    setDate(newValue);
                  }}
                  renderInput={(params) => <TextField {...params} />}
                />
              </LocalizationProvider>
            </FormControl>
          </Grid>

          {/* วันที่ทำงาน */}

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>Start Date</p>

              <LocalizationProvider dateAdapter={AdapterDateFns}>
                <DatePicker
                  value={date}
                  onChange={(newValue) => {
                    setDate(newValue);
                  }}
                  renderInput={(params) => <TextField {...params} />}
                />
              </LocalizationProvider>
            </FormControl>
          </Grid>
        </Grid>

        <Grid container spacing={3} sx={{ padding: 2 }}>
          <Grid item xs={4}>
            <FormControl fullWidth variant="outlined">
              <p>Username</p>

              <TextField
                id="LastName"
                variant="outlined"
                type="string"
                size="medium"
                value={user.LastName || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>

          <Grid item xs={4}>
            <FormControl fullWidth variant="outlined">
              <p>Password</p>

              <TextField
                id="LastName"
                variant="outlined"
                type="string"
                size="medium"
                value={user.LastName || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>
        </Grid>

        <Grid container spacing={3} sx={{ padding: 2 }}>
          <Grid item xs={4}>
            <FormControl fullWidth variant="outlined">
              <p>Email</p>

              <TextField
                id="LastName"
                variant="outlined"
                type="string"
                size="medium"
                value={user.LastName || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>

          <Grid item xs={4}>
            <FormControl fullWidth variant="outlined">
              <p>เบอร์โทร</p>

              <TextField
                id="LastName"
                variant="outlined"
                type="string"
                size="medium"
                value={user.LastName || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>
        </Grid>

        <Grid container spacing={3} sx={{ padding: 2 }}>
          <Grid item xs={4}>
            <FormControl fullWidth variant="outlined">
              <p>เงินเดือน</p>

              <TextField
                id="LastName"
                variant="outlined"
                type="string"
                size="medium"
                value={user.LastName || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>

          <Grid item xs={4}>
            <FormControl fullWidth variant="outlined">
              <p>กรุ๊ปเลือด</p>

              <TextField
                id="LastName"
                variant="outlined"
                type="string"
                size="medium"
                value={user.LastName || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>
        </Grid>


        <Grid container spacing={3} sx={{ padding: 2 }}>
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
          <Grid item xs={4}>
            <FormControl fullWidth variant="outlined">
              <p>ที่อยู่</p>

              <TextField
                id="LastName"
                variant="outlined"
                type="string"
                size="medium"
                value={user.LastName || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>
        </Grid>

          {/* วันเกิด */}
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>BirthDay</p>

              <LocalizationProvider dateAdapter={AdapterDateFns}>
                <DatePicker
                  value={date}
                  onChange={(newValue) => {
                    setDate(newValue);
                  }}
                  renderInput={(params) => <TextField {...params} />}
                />
              </LocalizationProvider>
            </FormControl>
          </Grid>

          {/* วันที่ทำงาน */}

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>Start Date</p>

              <LocalizationProvider dateAdapter={AdapterDateFns}>
                <DatePicker
                  value={date}
                  onChange={(newValue) => {
                    setDate(newValue);
                  }}
                  renderInput={(params) => <TextField {...params} />}
                />
              </LocalizationProvider>
            </FormControl>
          </Grid>

          <Grid item xs={4}>
            <FormControl fullWidth variant="outlined">
              <p>เจ้าหน้าที่</p>

              <TextField
                id="LastName"
                variant="outlined"
                type="string"
                size="medium"
                value={user.LastName || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>

          <Grid item xs={12}>
            <Button component={RouterLink} to="/" variant="contained">
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
  );
}

export default Manage_Save;
