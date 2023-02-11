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

//import { BookingsInterface } from "../../models/IBooking";
import { 
  CheckInOutInterface,
  CheckInOutStatusInterface, 
} from "../../models/ICheckInOut";
import { EmployeeInterface } from "../../models/IEmployee"; 

import { StorageInterface, ProductTypeInterface, ProductInterface } from "../../models/IStorage";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import { DesktopDateTimePicker } from "@mui/x-date-pickers";
import { InputLabel, Stack } from "@mui/material";
import { GetEmployees, GetProductTypes, GetProducts, GetStorage, GetStorages, UpdateStorage } from "./service/StorageHttpClientService";

import { createTheme, ThemeProvider } from "@mui/material/styles";
import { grey } from '@mui/material/colors';

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

const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
  props,
  ref
) {
  return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function StorageCreate() {
  const [producttypes, setProductTypes] = useState<ProductTypeInterface[]>([]);
  const [product, setProduct] = useState<ProductInterface[]>([]);
  const [employees, setEmployees] = useState<EmployeeInterface[]>([]);
  //const [storage, setStorage] = useState<StorageInterface[]>([]);

  const [number, setNumber] = useState("");
  const [st, setSt] = useState<StorageInterface[]>([]);
  const [storage, setStorage] = useState<StorageInterface>({});
  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);
  const [errorMSG, setErrorMSG] = useState("");

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
    const name = event.target.name as keyof typeof storage;
    setStorage({
      ...storage,
      [name]: event.target.value,
    });
  }; 

  const onChangeStorage = async (event: SelectChangeEvent) => {
    const id = event.target.value 
    let res = await GetStorage(id);
    if (res) {
      setStorage(res);
    }
  }; 

  const handleInputChangenumber = (

    event: React.ChangeEvent<{ id?: string; value: any }>

  ) => {

    const id = event.target.id as keyof typeof storage;

    const { value } = event.target;

    setStorage({ ...storage, [id]: value  === "" ? "" : Number(value)  });

  };

  const getStorage =  async () => {
    let res = await GetStorages();
    if (res) {
      setSt(res);
    }
  };

  const getProduct =  async () => {
    let res = await GetProducts();
    if (res) {
      setProduct(res);
    }
  };

  const getProductType =  async () => {
    let res = await GetProductTypes();
    if (res) {
      setProductTypes(res);
    }
  };

  const getEmployee =  async () => {
    let res = await GetEmployees();
    if (res) {
      setEmployees(res);
    }
  };

  useEffect(() => {
    getProductType();
    getProduct();
    getEmployee();
    getStorage();
  }, []);

  const convertType = (data: string | number | undefined | null) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  async function submit() {
    let data = {
      ID: convertType(storage.ID) || 0,
      ProductTypeID: convertType(storage.ProductTypeID),
      ProductID: convertType(storage.ProductID),
      //EmployeeID: convertType(checkinout.EmployeeID),
      EmployeeID: convertType(localStorage.getItem("id")+""),
      Quantity: convertType(number)||0,
      Time: storage.Time,
    };
    console.log(data)
    let res = await UpdateStorage(data);
    if (res.status) {
      setSuccess(true);
      setErrorMSG("")
    } else {
      setError(true);
      setErrorMSG(res.data)
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
        อัพเดตข้อมูลไม่สำเร็จ: {errorMSG}
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
              แก้ไขข้อมูลคลังสินค้าห้องพัก
            </Typography>
          </Box>
        </Box>
        <Divider />
        
        <Grid container spacing={3} sx={{ padding: 2 }}>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
            <InputLabel id="demo-simple-select-label">รหัสสินค้า</InputLabel>
              <Select
                labelId="demo-simple-select-label"
                native
                value={storage.ID + ""}
                label="รหัสสินค้า"
                onChange={onChangeStorage}
                inputProps={{
                  name: "ID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกรหัสสินค้า
                </option>
                {st.map((item: StorageInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.ID}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
            <InputLabel id="demo-simple-select-label">ชื่อสินค้า</InputLabel>
              <Select
                labelId="demo-simple-select-label"
                native
                value={storage.ProductID + ""}
                label="ชื่อสินค้า."
                onChange={handleChange}
                inputProps={{
                  name: "ProductID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกสินค้า
                </option>
                {product.map((item: ProductInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
            <InputLabel id="demo-simple-select-label">ประเภทของห้องพัก</InputLabel>
              <Select
                labelId="demo-simple-select-label"
                native
                value={storage.ProductTypeID + ""}
                label="ประเภทของห้องพัก."
                onChange={handleChange}
                inputProps={{
                  name: "ProductTypeID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกประเภทของสินค้า
                </option>
                {producttypes.map((item: ProductTypeInterface) => (
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
          id="Quantity" label="จำนวน" type="string" 
          InputLabelProps={{ shrink: true,}} 
          value={number} 
          onChange={(e)=>setNumber(e.target.value)}   
          inputProps={{name: "Quantity"}}    
          />
            </FormControl>
          </Grid>



          {/*<Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
            <InputLabel id="demo-simple-select-label">สถานะของห้องพัก</InputLabel>
              <Select
                labelId="demo-simple-select-label"
                native
                value={Store.StateID + ""}
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
                </Grid>*/}
          {/* <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>สถานะของห้อง</p>
              <Select
                native
                value={room.StateID + ""}
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
          {/* <Grid item xs={6}>
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
          </Grid> */}
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
                <LocalizationProvider dateAdapter={AdapterDayjs}>
                    <Stack spacing={3}>
                        <DesktopDateTimePicker
                        label="วันที่และเวลา"
                        value={storage.Time}
                        onChange={(newValue) => {
                            setStorage({
                                ...storage,
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
              to="/RoomW"
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
export default StorageCreate;