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

import { RoomInterface } from "../../models/IRoom";
import { CustomerInterface } from "../../models/modelCustomer/ICustomer";
import { RepairTypeInterface,RepairReqInterface } from "../../models/IRepairReq";

import { GetRooms } from "../Room/service/RoomHttpClientService";
import { GetCustomers } from "./service/RepRqHttpClientService";
import { 
    GetRepairTypes,
    GetRepairReqs,
    CreateRepairReq,
} from "./service/RepRqHttpClientService";

import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import { DesktopDateTimePicker } from "@mui/x-date-pickers";
import { InputLabel, Stack } from "@mui/material";

const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
  props,
  ref
) {
  return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function RepRqCreate() {
  const [rooms, setRooms] = useState<RoomInterface[]>([]);
  const [types, setTypes] = useState<RepairTypeInterface[]>([]);
  const [customers, setCustomers] = useState<CustomerInterface[]>([]);
  const [rep, setRep] = useState<RepairReqInterface>({});

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
    const name = event.target.name as keyof typeof rep;
    setRep({
      ...rep,
      [name]: event.target.value,
    });
  }; 


  const getRooms =  async () => {
    let res = await GetRooms();
    if (res) {
      setRooms(res);
    }
  };

  const getCustomers =  async () => {
    let res = await GetCustomers();
    if (res) {
      setCustomers(res);
    }
  };

  const getTypes =  async () => {
    let res = await GetRepairTypes();
    if (res) {
      setTypes(res);
    }
  };

  useEffect(() => {
    getRooms();
    getCustomers();
    getTypes();
  }, []);

  const convertType = (data: string | number | undefined | null) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  const handleInputChange = (
    event: React.ChangeEvent<{ id?: string; value: any }>
  ) => {
    const id = event.target.id as keyof typeof rep;
    const { value } = event.target;
    setRep({ ...rep, [id]: value });  
  };

  async function submit() {
    let data = {
        RoomID: convertType(rep.RoomID),
        RepairTypeID: convertType(rep.RepairTypeID),
        CustomerID: convertType(localStorage.getItem("id")),
        Note: rep.Note,
        Time: new Date(),
    };
    console.log(data)
    let res = await CreateRepairReq(data);
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
                CREATE REQUEST
            </Typography>
          </Box>
        </Box>
        <Divider />
        <Grid container spacing={3} sx={{ padding: 2 }}>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
            <InputLabel id="demo-simple-select-label">Room No.</InputLabel>
              <Select
                labelId="demo-simple-select-label"
                native
                value={rep.RoomID + ""}
                label="Room No."
                onChange={handleChange}
                inputProps={{
                  name: "RoomID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือก Room No.
                </option>
                {rooms.map((item: RoomInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.ID}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
            <InputLabel id="demo-simple-select-label">Type Of Problem</InputLabel>
              <Select
                native
                value={rep.RepairTypeID + ""}
                label="Type Of Problem"
                onChange={handleChange}
                inputProps={{
                  name: "RepairTypeID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือก Type Of Problem
                </option>
                {types.map((item: RepairTypeInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={15}>
          <FormControl fullWidth variant="outlined">
                <TextField
                    id="Note"
                    variant="outlined"
                    label="Note"
                    multiline
                    maxRows={4}
                    value={rep.Note}
                    onChange={handleInputChange}
                />
            </FormControl>
          </Grid>
          <Grid item xs={12}>
            <Button
              component={RouterLink}
              to="/Rep"
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
export default RepRqCreate;