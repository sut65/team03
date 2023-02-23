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
  Paper,
  Grid,
  Dialog,
  DialogTitle,
  DialogContent,
  TextField,
  Rating,
  Snackbar,
  Divider,
  FormControl,
  FormLabel,
  Select,
  SelectChangeEvent,
} from "@mui/material";
import DeleteIcon from "@mui/icons-material/Delete";
import Review_Save from "./Review_Save";
import { ReviewInterface, SystemworkInterface } from "../../models/IReview";
import Moment from "moment";
import EditIcon from "@mui/icons-material/Edit";
import MuiAlert, { AlertProps } from "@mui/material/Alert";
import Carousel from "react-material-ui-carousel";
import im10 from "../../Image/im10.jpg";
import im7 from "../../Image/im7.jpg";
import im9 from "../../Image/im9.png";

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
  const imgSize = { width: "200px", height: "200px" };
  const [openImage, setOpenImage] = React.useState(false);
  const [img, setimg] = React.useState({});
  const [message, setAlertMessage] = React.useState("");


  const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
    props,
    ref
  ) {
    return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
  });

//=====================การอัพเดท======================
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
          window.location.reload();
          setSuccess(true);
        } else {          
          setError(true);
          setAlertMessage(res.error);
        }
      });
  };
  console.log(UpdateReview);

  // ================================ดึงข้อมูล==============================
  //---------------------Department------------------------
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
  //---------------------Systemwork------------------------
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
        }
      });
  };
  // ================================จบดึงข้อมูล==============================

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

  //ปิดรูปเมื่อกด 2 ครั้ง
  const closeImage = () => {
    setOpenImage(false);
  };
//เปิดรูปเมื่อกด 2 ครั้ง
  const OpenImageonCilck = (item: ReviewInterface) => {
    setOpenImage(true);
    setimg(item.Reviewimage);
  };
//check ว่ามีรูปไหม
  const CheckImage = (item: ReviewInterface) => {
    if (item.Reviewimage != "") {
      return (
        <img
          src={`${item.Reviewimage}`}
          style={imgSize}
          onDoubleClick={() => OpenImageonCilck(item)}
        />
      );
    }
  };

  //ถ้ามีการกด ADD Review จะมีหน้าต่างขึ้นในส่วนการเขียนรีวิว
  const handleClickOpenForCreate = () => {
    setOpenForCreate(true);
  };
  const handleCloseForCreate = () => {
    setOpenForCreate(false);
  };

//กดแล้วจะเก็บค่าเพื่อไปทำการแก้ไขรีวิว
  const handleClickOpenForEdit = (item: ReviewInterface) => {
    setReview1(item);
    setStart(item.Star);
    setOpenForEdit(true);
  };
  const handleCloseForEdit = () => {
    setOpenForEdit(false);
  };

  //เก็บภาพและแปลงภาพ
  const handleImageChange = (event: any) => {
    const image = event.target.files[0];

    const reader = new FileReader();
    reader.readAsDataURL(image);
    reader.onload = () => {
      const base64Data = reader.result;
      setImageString(base64Data);
    };
  };

// ส่วนลบ
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
          window.location.reload();
        } else {
        }
      });
  };

  console.log(review);

  useEffect(() => {
    getReview();
    getDepartment();
    getSystemwork();
  }, []);

  //เรียกใช้งานรูปภาพเป็นสไลด์ๆ
  function Item(props: any) {
    return <img src={props.item.Image} width="100%" height="600px" />;
  }

  var Slider = [
    {
      Image: im9,
    },
    {
      Image: im7,
    },
    {
      Image: im10,
    },
  ];

  function ImageC() {
    return (
      <Carousel>
        {Slider.map((item, i) => (
          <Item key={i} item={item} />
        ))}
      </Carousel>
    );
  }

  return (
    //การใส่สีโดยเรียกใช้แบบ function
    <ThemeProvider theme={themeshow}>
      <Container maxWidth="xl">
        {/* การเรียกใช้ Carousel */}
        {ImageC()}
        
        <Button
          color="primary"
          variant="contained"
          aria-label="add"
          sx={{ mt: 2, padding: 2 }}
          onClick={() => handleClickOpenForCreate()}
        >
          Add Review
        </Button>
{/* แสดงข้อมูล */}
        <div>
          <Grid container columns={6} spacing={3} sx={{ mt: 2 }}>
            {review.map((item: ReviewInterface) => (
              <Grid item xs={3} key={item.ID}>
                <Paper>
                  {/* อีเมล ถังขยะ แก้ไข */}
                  <Grid item xs={6}>

                    <TextField
                      sx={{ "& fieldset": { border: "none" } }}
                      value={item.Customer.Email}
                    />

                    <IconButton aria-label="edit" style={{ float: "right" }}>
                      <EditIcon onClick={() => handleClickOpenForEdit(item)} />
                    </IconButton>

                    <IconButton aria-label="delete" style={{ float: "right" }}>
                      <DeleteIcon
                        onClick={() => deleteReview(Number(item.ID))}
                      />
                    </IconButton>

                    <Grid>
                      <Grid></Grid>
                    </Grid>
                  </Grid>
                  {/* เวลากับดาว */}
                  <Grid container columns={12}>
                    <Grid
                      display={"flex"}
                      justifyContent={"center"}
                      item
                      xs={6}
                    >
                      <Rating
                        value={item.Star}
                        size="medium"
                        sx={{ mt: 1.7 }}
                      />
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
                  {/* comment */}
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
                    {CheckImage(item)}
                      {/* <img
                        src={`${item.Reviewimage}`}
                        width="100%"
                        height="250"
                      /> */}
                    </Grid>
                  </Grid>

                </Paper>
              </Grid>
            ))}
          </Grid>
        </div>
        {/* popup หน้าภาพ */}
        <Dialog fullWidth maxWidth="md" open={openImage} onClose={closeImage}>
          <DialogContent>
            <img src={`${img}`} width="100%" height="100%" />
          </DialogContent>
        </Dialog>

        {/* popup บันทึก */}
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

        {/* popup แก้ไข */}
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
                id="success" 
                open={success}
                autoHideDuration={8000}
                onClose={handleClose}
                anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
              >
                <Alert onClose={handleClose} severity="success">
                  บันทึกข้อมูลสำเร็จ
                </Alert>
              </Snackbar>

              <Snackbar
                id="error"
                open={error}
                autoHideDuration={6000}
                onClose={handleClose}
              >
                <Alert onClose={handleClose} severity="error">
                  {message}
                </Alert>
              </Snackbar>

              <Paper elevation={3}>
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
                        native
                        value={review1.DepartmentID}
                        onChange={handleChange}
                        inputProps={{
                          name: "DepartmentID",
                        }}
                      >
                        <option value={0} key={0}></option>
                        {department.map((item: DepartmentInterface) => (
                          <option value={item.ID}>{item.Name}</option>
                        ))}
                      </Select>
                    </FormControl>
                  </Grid>

                  {/* ComboboxPosition */}
                  <Grid item xs={3}>
                    <FormLabel>System Work</FormLabel>
                    <FormControl fullWidth variant="outlined">
                      <Select
                        native
                        value={review1.SystemworkID}
                        onChange={handleChange}
                        inputProps={{
                          name: "SystemworkID",
                        }}
                      >
                        <option value={0} key={0}></option>
                        {systemwork.map((item: SystemworkInterface) => (
                          <option value={item.ID}>{item.Name}</option>
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
                        disabled
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
    </ThemeProvider>
  );
}

export default Review_list;
