import React, { useState, useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import FormLabel from "@mui/material/FormLabel";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import FormControl from "@mui/material/FormControl";
import Container from "@mui/material/Container";
import Paper from "@mui/material/Paper";
import Grid from "@mui/material/Grid";
import Box from "@mui/material/Box";
import Select from "@mui/material/Select";
import MenuItem from "@mui/material/MenuItem";
//เพิ่ม
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import { DatePicker } from '@mui/x-date-pickers/DatePicker';
import { AdapterDateFns } from '@mui/x-date-pickers/AdapterDateFns';
//เพิ่ม
import { Message } from "@mui/icons-material";
import Typography from "@mui/material/Typography";
import Divider from "@mui/material/Divider";
import Snackbar from "@mui/material/Snackbar";
import { Link } from "react-router-dom";
import MuiAlert, { AlertProps } from "@mui/material/Alert";
import FormHelperText from "@mui/material/FormHelperText";
import { CheckroomInterface } from "../../models/ICheckroom";
import { DamageInterface } from "../../models/modelCheckroom/IDamage";
import { StatusInterface } from "../../models/modelCheckroom/IStatus"; 
import { ProductInterface } from "../../models/IStorage";
import { EmployeeInterface } from "../../models/IEmployee";
import { RoomInterface } from "../../models/IRoom";
 //เพิ่ม
import { SelectChangeEvent } from "@mui/material";
import moment from 'moment';
import { time } from "console";
import { GetRoom, GetProduct, GetDamage, GetStatus, Checkrooms } from "./service/service";
const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
  props,
  ref
) {
  return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function Checkroom() {
    const [checkroom, setCheckroom] = useState<Partial<CheckroomInterface>>({});
    const [damage, setDamage] = useState<DamageInterface[]>([]);
    const [status, setStatus] = useState<StatusInterface[]>([]);
    const [room, setRoom] = useState<RoomInterface[]>([]);
    const [Product, setProduct] = useState<ProductInterface[]>([]);
    const [date, setDate] = useState<Date | null>(new Date());
    const [employee, setEmployee] = React.useState<EmployeeInterface>();
    const [message, setAlertMessage] = useState("");
    
    const [success, setSuccess] = useState(false);
    const [error, setError] = useState(false);

// =========================(handleClose)====================================================

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

  // =========================(HandleChange)====================================================

  const handleChange = (event: SelectChangeEvent) => {
    const name = event.target.name as keyof typeof checkroom;
    console.log(event.target.name);
    console.log(event.target.value);
    
    setCheckroom({
      ...checkroom,
      [name]: event.target.value,
    });
    console.log(checkroom);
  };
  
  const handleInputChange = (
    event: React.ChangeEvent<{ id?: string; value: any }>
  ) => {
    const id = event.target.id as keyof typeof checkroom;

    const { value } = event.target;

    setCheckroom({ ...checkroom, [id]: value });
  };

  const getRoom = async () => {
    let res = await GetRoom();
    if (res) {
        setRoom(res);
    }
};

const getProduct = async () => {
    let res = await GetProduct();
    if (res) {
        setProduct(res);
    }
};

const getDamage = async() => {
    let res = await GetDamage();
    if (res) {
        setDamage(res);
    }
};

const getStatus = async() => {
  let res = await GetStatus();
  if (res) {
      setStatus(res);
  }
}
const apiUrl = "http://localhost:8080";

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
        checkroom.EmployeeID = res.data.ID
      } else {
        console.log("else");
      }
    });
};


useEffect(() => {
    getRoom();
    getProduct();
    getDamage();
    getStatus();
    getEmployee();
}, []);

  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  const convertTypeC = (data: string | number | null) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  async function submit() {
    let data = {
      RoomID: convertType(checkroom.RoomID),
      ProductID: convertType(checkroom.ProductID),
      DamageID: convertType(checkroom.DamageID),
      StatusID: convertType(checkroom.StatusID),
      Date: date,
      EmployeeID: convertTypeC(localStorage.getItem('id')),
    };
    let res = await Checkrooms(data);
    if (res.status) {
      setAlertMessage("บันทึกสำเร็จ");
      setSuccess(true);
    } else {
      console.log(res.message);
      setAlertMessage(res.message);
      setError(true);
    }
  }

return (
    <div>
      <Container maxWidth="sm" sx={{ marginTop: 6 }}>
        <Paper
          elevation={4}
          sx={{
            marginBottom: 2,
            marginTop: 2,
            padding: 1,
            paddingX: 2,
            display: "flex",
            justifyContent: "flex-start",
          }}
        >
          <h4 style={{ color: "#000000" }}>Checkroom</h4>
        </Paper>
        <form>
          <Paper
            variant="outlined"
            sx={{ padding: 2, paddingTop: 1, marginBottom: 2 }}

            
          >
          
            <Grid container spacing={1} sx={{ marginBottom: 2}}>
                    
                   
              {/*=======================================(select Room)===========================================================*/}
              <Grid
                xs={12}
                
                sx={{ display: "flex", alignItems: "center", margin: 1 }}
              >
                <FormLabel
                  id="demo-simple-select-helper-label"
                  sx={{ marginRight: 6.5, fontSize: 17, paddingBottom: 1}}
                >
                  Room:
                </FormLabel>
                <Select
                  required
                  id="Room_ID"
                  value={checkroom.RoomID + ""}
                  onChange={handleChange}
                  fullWidth
                  native
                  inputProps={{
                    name: "RoomID",
                  }}
                >
                  <option aria-label="None" value="">
                    กรุณาเลือกหมายเลขห้อง
                  </option>
                  {room.map((item: RoomInterface) => (
                    <option value={item.ID} key={item.ID}>
                      {item.Room_No}
                    </option>
                  ))}
                </Select>
                <FormHelperText disabled sx={{ width: 350, marginLeft: 2 }}>
                  กรุณาเลือกห้อง
                </FormHelperText>
              </Grid>

              {/*=======================================(Product)===========================================================*/}
              <Grid
                xs={12}
                
                sx={{ display: "flex", alignItems: "center", margin: 1 }}
              >
                <FormLabel
                  id="demo-simple-select-helper-label"
                  sx={{ marginRight: 4.5, fontSize: 17, paddingBottom: 2 }}
                >
                  Product:
                </FormLabel>
                <Select
                  // labelId="demo-simple-select-helper-label"
                  id="PrductID"
                  value={checkroom.ProductID + ""}
                  onChange={handleChange}
                  native
                  inputProps={{
                    name: "ProductID",
                  }}
                  fullWidth
                >
                  <option aria-label="None" value="">
                    กรุณาเลือกอุปกรณ์ภายในห้อง
                  </option>
                  {Product.map((item: ProductInterface) => (
                    <option value={item.ID} key={item.ID}>
                      {item.Name}
                    </option>
                  ))}
                </Select>
                <FormHelperText disabled sx={{ width: 350, marginLeft: 2 }}>
                  เลือกอุปกรณ์ภายในห้อง
                </FormHelperText>
              </Grid>

              {/*=======================================(Damage)===========================================================*/}
              <Grid
                xs={12}
              
                sx={{ display: "flex", alignItems: "center", margin: 1 }}
              >
                <FormLabel
                  id="demo-simple-select-helper-label"
                  sx={{ marginRight: 3.8, fontSize: 17, paddingBottom: 2 }}
                >
                  Damage:
                </FormLabel>
                <Select
                  // labelId="demo-simple-select-helper-label"
                  id="DamageID"
                  value={checkroom.DamageID + ""}
                  onChange={handleChange}
                  native
                  inputProps={{
                    name: "DamageID",
                  }}
                  fullWidth
                >
                  <option aria-label="None" value="">
                    กรุณาเลือกความเสียหาย
                  </option>
                  {damage.map((item: DamageInterface) => (
                    <option value={item.ID} key={item.ID}>
                      {item.Description}
                    </option>
                  ))}
                </Select>
                <FormHelperText disabled sx={{ width: 350, marginLeft: 2 }}>
                  เลือกอุปกรณ์ภายในห้อง
                </FormHelperText>
              </Grid>

              {/*=======================================(Status)===========================================================*/}
              <Grid
                xs={12}
                
                sx={{ display: "flex", alignItems: "center", margin: 1 }}
              >
                <FormLabel
                  id="demo-simple-select-helper-label"
                  sx={{ marginRight: 6, fontSize: 17, paddingBottom: 2 }}
                >
                  Status:
                </FormLabel>
                <Select
                  // labelId="demo-simple-select-helper-label"
                  id="StatusID"
                  value={checkroom.StatusID + ""}
                  onChange={handleChange}
                  native
                  inputProps={{
                    name: "StatusID",
                  }}
                  fullWidth
                >
                   <option aria-label="None" value="">
                    กรุณาเลือกสถานะของห้องพัก
                  </option>
                  {status.map((item: StatusInterface) => (
                    <option value={item.ID} key={item.ID}>
                      {item.S_Name}
                    </option>
                  ))}
                </Select>
                <FormHelperText disabled sx={{ width: 350, marginLeft: 2 }}>
                  เลือกสถานะของห้อง
                </FormHelperText>
              </Grid>


              
                         
              <Grid
                container
                xs={12}
                md={12}
                gap={2}
                sx={{ justifyContent: "center", margin: 1 }}
              >
                <Button variant="contained" size="large" onClick={submit}>
                  บันทึกการตรวจสอบ
                </Button>
                <Link to="/checkroom/list" style={{ textDecoration: "none" }}>
                  <Button
                    variant="contained"
                    size="large"
                    style={{ backgroundColor: "#fff", color: "#1976d2" }}
                  >
                    กลับ
                  </Button>
                </Link>
              </Grid>
            </Grid>
          </Paper>
        </form>
      </Container>
      <Snackbar
        open={success}
        autoHideDuration={5000}
        onClose={handleClose}
        anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
      >
        <Alert onClose={handleClose} severity="success">
           {/* บันทึกข้อมูลสำเร็จ */}
           {message}
        </Alert>
      </Snackbar>

      <Snackbar
        open={error}
        autoHideDuration={5000}
        onClose={handleClose}
        anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
      >
        <Alert onClose={handleClose} severity="error">
             {/* บันทึกข้อมูลไม่สำเร็จ */}
         {message}
        </Alert>
      </Snackbar>
    </div>
  );
}


export default Checkroom;