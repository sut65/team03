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

import { StorageInterface, ProductTypeInterface, ProductInterface } from "../../models/IStorage";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import { DesktopDateTimePicker } from "@mui/x-date-pickers";
import { InputLabel, Stack } from "@mui/material";
import { GetEmployees, GetProductTypes, GetProducts, CreateStorage } from "./service/StorageHttpClientService";

import { grey } from '@mui/material/colors';
import { Message } from "@mui/icons-material";


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
  const [storage, setStorage] = useState<StorageInterface>({
    Time: new Date(),
    Quantity: 0,
  });
  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);
  const [message, setAlertMessage] = useState("");

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

  const handleInputChangenumber = (

    event: React.ChangeEvent<{ id?: string; value: any }>

  ) => {

    const id = event.target.id as keyof typeof storage;

    const { value } = event.target;

    setStorage({ ...storage, [id]: value  === "" ? "" : Number(value)  });
    console.log(`id: ${id} value:${value}`)

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

  // const getEmployee =  async () => {
  //   let res = await GetEmployees();
  //   if (res) {
  //     setEmployees(res);
  //   }
  // };

  useEffect(() => {
    getProductType();
    getProduct();
    //getEmployee();
  }, []);

  const convertType = (data: string | number | undefined | null) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  async function submit() {
    let data = {
      ProductTypeID: convertType(storage.ProductTypeID),
      ProductID: convertType(storage.ProductID),
      EmployeeID: convertType(localStorage.getItem("id")),
      Quantity: convertType(number)||0,
      Time: storage.Time,
    };

    console.log(data)

    let res = await CreateStorage(data);
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
              เพิ่มข้อมูลคลังสินค้าห้องพัก
            </Typography>
          </Box>
        </Box>
        <Divider />
        
        <Grid container spacing={3} sx={{ padding: 2 }}>
          {/*<Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
            <TextField
          id="outlined-number" label="รหัสสินค้า" type="number" InputLabelProps={{ shrink: true,}} value={storage?.ID} onChange={handleInputChangenumber}       
          />
            </FormControl>
        </Grid>*/}

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
            <InputLabel id="demo-simple-select-label">ประเภทของสินค้าห้องพัก</InputLabel>
              <Select
                labelId="demo-simple-select-label"
                native
                value={storage.ProductTypeID + ""}
                label="ประเภทของสินค้าห้องพัก."
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
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
                <LocalizationProvider dateAdapter={AdapterDayjs}>
                    <Stack spacing={3}>
                        <DesktopDateTimePicker
                        label="วันที่และเวลาที่บันทึกข้อมูล"
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