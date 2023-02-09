import React, { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import { grey } from "@mui/material/colors";
import { DepartmentInterface } from "../../models/IEmployee";
import {
  IconButton,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Paper,
  Grid,
  Dialog,
  DialogTitle,
  DialogContent,
  TextField,
  Rating,
  Snackbar,
  Alert,
  Divider,
  FormControl,
  FormLabel,
  MenuItem,
  Select,
  SelectChangeEvent,
} from "@mui/material";
import moment from "moment";
import DeleteIcon from "@mui/icons-material/Delete";
import { Margin } from "@mui/icons-material";
import Review_Save from "./Review_Save";
import { ReviewInterface, SystemworkInterface } from "../../models/IReview";
import Moment from "moment";
import EditIcon from "@mui/icons-material/Edit";

const themeshow = createTheme({
  palette: {
    primary: {
      // Purple and green play nicely together.
      main: grey[800],
    },
    secondary: {
      // This is green.A700 as hex.
      main: "#e8f5e9",
    },
  },
});

function Review_list() {
  Moment.locale("th");

  const [imageString, setImageString] = React.useState<
    string | ArrayBuffer | null
  >(null);
  const [openForCreate, setOpenForCreate] = React.useState(false);
  const [openForEdit, setOpenForEdit] = React.useState(false);
  const [review, setReview] = React.useState<ReviewInterface[]>([]);
  const [review1, setReview1] = React.useState<Partial<ReviewInterface>>({});
  const [department, setDepartment] = React.useState<DepartmentInterface[]>([]);
  const [systemwork, setSystemwork] = React.useState<SystemworkInterface[]>([]);
  const [start, setStart] = React.useState<number | null>();
  const [success, setSuccess] = React.useState(false);
  const [error, setError] = React.useState(false);

  const UpdateReview = () => {
    let UpdateData = {
      ID: review1.ID,
      Comment: review1.Comment,
      Star: start,
      Reviewimage: imageString,
      DepartmentID: review1.DepartmentID,
      SystemworkID: review1.SystemworkID,
    };
    const apiUrl = "http://localhost:8080/Reviews";
    const requestOptions = {
      method: "PATCH",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },
      body: JSON.stringify(UpdateData),
    };

    fetch(apiUrl, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          // setSuccess(true);
        } else {
          //setError(true);
        }
      });
  };
  console.log(UpdateReview);

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
  //---------------------Systemwork-------------------------------------
  const getSystemwork = async () => {
    const apiUrl = `http://localhost:8080/Systemworks`;

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
          setSystemwork(res.data);
        } else {
          console.log("else");
        }
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
    event: React.ChangeEvent<{ id?: string; value: any }>
  ) => {
    const id = event.target.id as keyof typeof Review_Save;

    const { value } = event.target;

    setReview1({ ...review1, [id]: value });
  };

  const handleChange = (event: SelectChangeEvent<number>) => {
    const name = event.target.name as keyof typeof review;
    setReview1({
      ...review1,
      [name]: event.target.value,
    });
  };
  const handleClickOpenForCreate = () => {
    setOpenForCreate(true);
  };
  const handleCloseForCreate = () => {
    setOpenForCreate(false);
  };

  const handleClickOpenForEdit = (item: ReviewInterface) => {
    setReview1(item);
    setStart(item.Star);
    setOpenForEdit(true);
  };
  const handleCloseForEdit = () => {
    setOpenForEdit(false);
  };
  const handleImageChange = (event: any) => {
    const image = event.target.files[0];

    const reader = new FileReader();
    reader.readAsDataURL(image);
    reader.onload = () => {
      const base64Data = reader.result;
      setImageString(base64Data);
    };
  };

  const deleteReview = (id: number) => {
    const apiUrl = "http://localhost:8080/Reviews/" + id;
    const requestOptions = {
      method: "DELETE",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },
    };

    fetch(apiUrl, requestOptions)
      .then((response) => response.json())
      .then(async (res) => {
        if (res.data) {
          //setSuccess(true);
          //await timeout(1000); //for 1 sec delay
          window.location.reload();
        } else {
          //setError(true);
        }
      });
  };
  const getReview = async () => {
    const apiUrl = `http://localhost:8080/Review/` + localStorage.getItem("id");

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
        //   console.log(res.data); //show ข้อมูล

        if (res.data) {
          setReview(res.data);
          console.log(res.data);
        } else {
          // console.log("else");
        }
      });
  };
  console.log(review)
  useEffect(() => {
    getReview();
    getDepartment();
    getSystemwork();
    // getOfficer();
  }, []);

  return (
    <Container>
      {" "}
      <Button
        color="secondary"
        aria-label="add"
        onClick={() => handleClickOpenForCreate()}
      >
        Add
      </Button>
      <div>
        <Grid container columns={6} spacing={3} sx={{ mt: 2 }}>
          {review.map((item: ReviewInterface) => (
            <Grid item xs={3} key={item.ID}>
              <Paper>
                <Grid item xs={6}>
                  <TextField
                    sx={{ "& fieldset": { border: "none" } }}
                    value={item.Customer.Email}
                  />
                  <IconButton aria-label="edit" style={{ float: "right" }}>
                    <EditIcon onClick={() => handleClickOpenForEdit(item)} />
                  </IconButton>
                  <IconButton aria-label="delete" style={{ float: "right" }}>
                    <DeleteIcon onClick={() => deleteReview(Number(item.ID))} />
                  </IconButton>
                  <Grid>
                    <Grid></Grid>
                  </Grid>
                </Grid>
                <Grid container columns={12}>
                  <Grid display={"flex"} justifyContent={"center"} item xs={6}>
                    <Rating value={item.Star} size="medium" sx={{ mt: 1.7 }} />
                  </Grid>
                  <Grid item xs={6}>
                    <TextField
                      sx={{ "& fieldset": { border: "none" } }}
                      value={`${Moment(item.Reviewdate).format(
                        "DD/MM/YYYY HH:mm:ss A"
                      )}`}
                    />
                  </Grid>
                </Grid>
                <Grid>
                  <TextField
                    fullWidth
                    multiline
                    sx={{ "& fieldset": { border: "none" } }}
                    value={item.Comment}
                  />
                </Grid>
                <Grid container display={"flex"} justifyContent={"center"}>
                  <Grid>
                    <img
                      src={`${item.Reviewimage}`}
                      width="100%"
                      height="250"
                    />
                  </Grid>
                </Grid>
              </Paper>
            </Grid>
          ))}
        </Grid>
      </div>
      <Dialog
        fullWidth
        maxWidth="md"
        open={openForCreate}
        onClose={handleCloseForCreate}
      >
        <DialogTitle>Write comment</DialogTitle>
        <DialogContent>
          <Review_Save />
        </DialogContent>
      </Dialog>
      <Dialog
        fullWidth
        maxWidth="md"
        open={openForEdit}
        onClose={handleCloseForEdit}
      >
        <DialogTitle>Edit comment</DialogTitle>
        <DialogContent>
          <Container maxWidth="md">
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
                    Review
                  </Typography>
                </Box>
              </Box>

              <Divider />

              <Grid
                container
                spacing={3}
                sx={{ padding: 2 }}
                style={{ marginLeft: "10.5%" }}
              >
                {/* ComboboxDepartment */}
                <Grid item xs={3}>
                  <FormLabel>Department</FormLabel>
                  <FormControl fullWidth variant="outlined">
                    <Select
                      value={review1.DepartmentID}
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
                  <FormLabel>System Work</FormLabel>
                  <FormControl fullWidth variant="outlined">
                    <Select
                      value={review1.SystemworkID}
                      onChange={handleChange}
                      inputProps={{
                        name: "SystemworkID",
                      }}
                    >
                      <MenuItem value={0} key={0}>
                        เลือกตำแหน่ง
                      </MenuItem>
                      {systemwork.map((item: SystemworkInterface) => (
                        <MenuItem value={item.ID}>{item.Name}</MenuItem>
                      ))}
                    </Select>
                  </FormControl>
                </Grid>
              </Grid>

              <Grid
                container
                spacing={3}
                sx={{ padding: 2 }}
                style={{ marginLeft: "14.5%" }}
              >
                <Grid item xs={6}>
                  <FormControl>
                    <FormLabel>Start</FormLabel>
                    <Rating
                      name="simple-controlled"
                      defaultValue={review1.Star}
                      onChange={(event, newValue) => {
                        setStart(newValue);
                      }}
                    />
                  </FormControl>
                </Grid>
              </Grid>

              <Grid container spacing={3} sx={{ padding: 2 }}>
                <Grid item xs={10}>
                  <FormControl fullWidth variant="outlined">
                    <FormLabel>Comment</FormLabel>

                    <TextField
                      id="Comment"
                      variant="outlined"
                      type="string"
                      size="medium"
                      value={review1.Comment || ""}
                      onChange={handleInputChange}
                    />
                  </FormControl>
                </Grid>
              </Grid>

              <Grid container spacing={3} sx={{ padding: 2 }}>
                <Grid item xs={10}>
                  <FormControl fullWidth variant="outlined">
                    <FormLabel>Image</FormLabel>
                    <img src={`${imageString}`} width="500" height="500" />
                    <input type="file" onChange={handleImageChange} />
                  </FormControl>
                </Grid>

                {/* วันที่ทำงาน */}
                <Grid item xs={12}>
                  <Button component={RouterLink} to="/RW" variant="contained">
                    Back
                  </Button>

                  <Button
                    style={{ float: "right" }}
                    onClick={UpdateReview}
                    variant="contained"
                    color="primary"
                  >
                    Submit
                  </Button>
                </Grid>
              </Grid>
            </Paper>
          </Container>
        </DialogContent>
      </Dialog>
    </Container>
  );
}

export default Review_list;
