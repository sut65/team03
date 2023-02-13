import React, { useState, useEffect } from "react";
import Paper from "@mui/material/Paper";
import Grid from "@mui/material/Unstable_Grid2";
import TextField from "@mui/material/TextField";
import Container from "@mui/material/Container";
import FormLabel from "@mui/material/FormLabel";
import FormHelperText from "@mui/material/FormHelperText";
import Button from "@mui/material/Button";
import Snackbar from "@mui/material/Snackbar";
import MuiAlert, { AlertProps } from "@mui/material/Alert";
import { Link, Link as RouterLink } from "react-router-dom";
import { CustomerInterface } from "../../models/modelCustomer/ICustomer";
import { NametitleInterface } from "../../models/modelCustomer/INametitle";
import {  GetCustomerByUID,  UpdateCustomer } from "./service/servicecus";
import { FormControl, Select, SelectChangeEvent } from "@mui/material";
import { ProvinceInterface } from "../../models/modelCustomer/IProvince";
import { GenderInterface } from "../../models/modelCustomer/IGender";


const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
  props,

  ref
) {
  return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});


function EditCustomer() {
  
  const [customer, setCustomer] = useState<CustomerInterface>({});
  const [nametitle, setNametitle] = useState<NametitleInterface[]>([]);
  const [gender, setGender] = useState<GenderInterface[]>([]);
  const [province, setProvince] = useState<ProvinceInterface[]>([]);
  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);
  
  const handleChange = (event: SelectChangeEvent<number>) => { 
  const name = event.target.name as keyof typeof customer;
    setCustomer({
      ...customer,
      [name]: event.target.value,
    });
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

  const handleInputChange = (
    event: React.ChangeEvent<{ id?: string; value: unknown }>
) => {
    const id = event.target.id as keyof typeof customer;
    setCustomer({
        ...customer,
        [id]: event.target.value,
    });};

    const handleInputAge = (
      event: React.ChangeEvent<{ id?: string; value: unknown }>
  ) => {
      const id = event.target.id as keyof typeof customer;
      const { value } = event.target;
      setCustomer({
          ...customer,
          [id]: value  === "" ? "" : Number(value)  
      });};
  

    const handleSelectChange = (event: SelectChangeEvent<string>) => {
      const name = event.target.name as keyof typeof customer;
      setCustomer({
          ...customer,
          [name]: event.target.value,
      });

  };


  // =========================(Fetch API)====================================================

  const apiUrl = "http://localhost:8080";
  const requestOptionsGet = {
    method: "GET",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
    },
  };
  

  const fetchGender = async () => {
    fetch(`${apiUrl}/customers/genders`, requestOptionsGet)
      .then((response) => response.json())
      .then((result) => {
        setGender(result.data);
      });
  };
  const fetchNametitle = async () => {
    fetch(`${apiUrl}/nametitles`, requestOptionsGet)
      .then((response) => response.json())
      .then((result) => {
        setNametitle(result.data);
      });
  };
  const fetchProvince = async () => {
    fetch(`${apiUrl}/provinces`, requestOptionsGet)
      .then((response) => response.json())
      .then((result) => {
        setProvince(result.data);
      });
  };

  const fetchCustomer = async () => {
    let res = await GetCustomerByUID();
    res && setCustomer(res);
  };

  useEffect(() => {
    fetchGender();
    fetchNametitle();
    fetchProvince();
    fetchCustomer();
  }, []);

  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  async function update() {
    let newdata = {
      ID: convertType(customer.ID),
      Nametitle_ID: convertType(customer.Nametitle_ID),
      Gender_ID: convertType(customer.Gender_ID),
      Province_ID: convertType(customer.Province_ID),
      Firstname: customer.FirstName,
      Lastname: customer.LastName,
      Age: customer.Age,
      Phone: customer.Phone,
      Email: customer.Email,
      Password: customer.Password,
    };
    console.log(newdata);
    console.log(JSON.stringify(newdata))
    let res = await UpdateCustomer(newdata);
    if (res) {
      setSuccess(true);
      console.log(newdata)
    } else {
      setError(true);
    }
    console.log(JSON.stringify(newdata))
  };


  return (
    <div>
      <Snackbar
        open={success}
        autoHideDuration={5000}
        onClose={handleClose}
        anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
      >
        <Alert onClose={handleClose} severity="success">
          บันทึกข้อมูลสำเร็จ
        </Alert>
      </Snackbar>

      <Snackbar
        open={error}
        autoHideDuration={5000}
        onClose={handleClose}
        anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
      >
        <Alert onClose={handleClose} severity="error">
          บันทึกข้อมูลไม่สำเร็จ
        </Alert>
      </Snackbar>
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
          <h4 style={{ color: "#000000" }}>Profile</h4>
        </Paper>
        <form>
          <Paper
            variant="outlined"
            sx={{ padding: 2, paddingTop: 1, marginBottom: 2 }}

          >


            {/*=======================================(Title)===========================================================*/}
            <Grid container spacing={2} sx={{ marginBottom: 1.5 }}>
              <Grid
                xs={12}
                md={8}
                sx={{ display: "flex", alignItems: "center", margin: 1 }}
              >
                <FormLabel
                  id="demo-simple-select-helper-label"
                  sx={{ marginRight: 6.5, fontSize: 17, paddingBottom: 2 }}
                >
                  Title:
                </FormLabel>
                <FormControl fullWidth variant="outlined">
                  <Select
                    id="Nametitle_ID"
                    native
                    value={customer.Nametitle_ID}
                    onChange={handleChange}
                    inputProps={{
                      name: "Nametitle_ID",
                    }}
                  >
                    {nametitle.map((item: NametitleInterface) => (
                      <option value={item.ID}>{item.NT_Name}</option>
                    ))}
                  </Select>
                </FormControl>
                <FormHelperText disabled sx={{ width: 350, marginLeft: 2 }}>
                  คำนำหน้าชื่อ
                </FormHelperText>
              </Grid>

              {/*============================================(First name)======================================================*/}
              <Grid xs={6} md={6}>
                <p style={{ color: "grey", fontSize: 17 }}>First name</p>
                <FormControl fullWidth variant="outlined">
                  <TextField
                    id="FirstName"
                    variant="outlined"
                    type="string"
                    size="medium"
                    value={customer.FirstName || ""}
                    onChange={handleInputChange}
                  />
                </FormControl>

              </Grid>

              {/*=============================================(Last name)=====================================================*/}
              <Grid xs={6} md={6}>
                <p style={{ color: "grey", fontSize: 17 }}>Last name</p>
                <FormControl fullWidth variant="outlined">
                  <TextField
                    id="LastName"
                    variant="outlined"
                    type="string"
                    size="medium"
                    value={customer.LastName || ""}
                    onChange={handleInputChange}
                  />
                </FormControl>
              </Grid>
            </Grid>

            <Grid container spacing={2} sx={{ marginBottom: 1.5 }}>
              {/*============================================(Age)======================================================*/}
              <Grid xs={6} md={6}>
                <p style={{ color: "grey", fontSize: 17 }}>Age</p>
                <FormControl fullWidth variant="outlined">
                  <TextField
                    id="Age"
                    size="medium"
                    value={customer.Age}
                    onChange={handleInputAge}
                  />
                </FormControl>
              </Grid>
              {/*=============================================(Phone)=====================================================*/}
              <Grid xs={6} md={6}>
                <p style={{ color: "grey", fontSize: 17 }}>Phone number</p>
                <FormControl fullWidth variant="outlined">
                  <TextField
                    id="Phone"
                    variant="outlined"
                    type="string"
                    size="medium"
                    value={customer.Phone || ""}
                    onChange={handleInputChange}
                  />
                </FormControl>
              </Grid>
            </Grid>


            {/*===========================================(email)=======================================================*/}
            <Grid container spacing={1}>
              <Grid
                xs={12}
                md={12}
                sx={{ display: "flex", alignItems: "center", margin: 1 }}
              >
                <FormLabel sx={{ marginRight: 7, fontSize: 17 }}>
                  Email:
                </FormLabel>
                <FormControl fullWidth variant="outlined">
                  <TextField
                    id="Email"
                    variant="outlined"
                    type="string"
                    size="medium"
                    value={customer.Email || ""}
                    onChange={handleInputChange}
                  />
                </FormControl>
              </Grid>

              {/*=======================================(select Gender)===========================================================*/}
              <Grid
                xs={12}
                md={9}
                sx={{ display: "flex", alignItems: "center", margin: 1 }}
              >
                <FormLabel
                  id="demo-simple-select-helper-label"
                  sx={{ marginRight: 5.5, fontSize: 17, paddingBottom: 2 }}
                >
                  Gender:
                </FormLabel>
                <Select
                  id="Gender_ID"
                  native
                  value={customer.Gender_ID}
                  onChange={handleChange}
                  inputProps={{
                    name: "Gender_ID",
                  }}
                >
                  {gender.map((item: GenderInterface) => (
                    <option value={item.ID}>{item.G_Name}</option>
                  ))}
                </Select>

              </Grid>

              {/*=======================================(Province)===========================================================*/}
              <Grid
                xs={12}
                md={8}
                sx={{ display: "flex", alignItems: "center", margin: 1 }}
              >
                <FormLabel
                  id="demo-simple-select-helper-label"
                  sx={{ marginRight: 4.5, fontSize: 17, paddingBottom: 2 }}
                >
                  Province:
                </FormLabel>

                <Select
                  id="Province_ID"
                  native
                  value={customer.Province_ID}
                  onChange={handleChange}
                  inputProps={{
                    name: "Province_ID",
                  }}
                >
                  {province.map((item: ProvinceInterface) => (
                    <option value={item.ID}>{item.P_Name}</option>
                  ))}
                </Select>
              </Grid>

              <Grid
                container
                xs={12}
                md={12}
                gap={2}
                sx={{ justifyContent: "center", margin: 1 }}
              >
                <Button variant="contained" size="large" onClick={update}>
                  บันทึกการแก้ไข
                </Button>

                <Link to="/customer/profile" style={{ textDecoration: "none" }}>
                  <Button
                    variant="contained"
                    size="large"
                    style={{ backgroundColor: "#fff", color: "#1976d2" }}
                  >
                    กลับ
                  </Button>
                </Link>
              </Grid>


            </Grid></Paper></form></Container></div>
  )



}

export default EditCustomer;