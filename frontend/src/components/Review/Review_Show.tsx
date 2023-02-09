import React, { useEffect, useState } from "react";

import { Link as RouterLink } from "react-router-dom";

import Typography from "@mui/material/Typography";

import Button from "@mui/material/Button";

import Container from "@mui/material/Container";

import Box from "@mui/material/Box";

import { createTheme, ThemeProvider } from "@mui/material/styles";
import { grey } from "@mui/material/colors";
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
} from "@mui/material";
import moment from "moment";
import DeleteIcon from "@mui/icons-material/Delete";
import { Margin } from "@mui/icons-material";
import Review_Save from "./Review_Save";
import { ReviewInterface } from "../../models/IReview";
import Moment from "moment";

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

function Review_Show() {
  Moment.locale("th");
  const [openForCreate, setOpenForCreate] = React.useState(false);
  const [review, setReview] = React.useState<ReviewInterface[]>([]);
  // const c = { width: '200x', height: '200px' };
  // const d = { width: '0', height: '0' };
  const imgSize = { width: "200px", height: "200px" };
  //const [myImageStyle, setMyImage] = React.useState({width: '200px', height: '200px' })
  const [openImage, setOpenImage] = React.useState(false);
  const [img, setimg] = React.useState({});

  //ตรวจสอบภาพ
  // const CheckImage = (item: ReviewInterface) => {
  //   if (item.Reviewimage == "")
  //   {setMyImage({width: '200px', height: '200px' })}
  // };
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

  const OpenImageonCilck = (item: ReviewInterface) => {
    setOpenImage(true);
    setimg(item.Reviewimage);
  };

  const closeImage = () => {
    setOpenImage(false);
  };

  const getReview = async () => {
    const apiUrl = `http://localhost:8080/Reviews`;

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

  useEffect(() => {
    getReview();
  }, []);

  return (
    <div>
      <Container maxWidth="xl">
        <Box
          sx={{
            backgroundImage: `url(${"https://img.freepik.com/free-photo/modern-luxury-hotel-office-reception-lounge-with-meeting-room_105762-1772.jpg?w=1060&t=st=1674825855~exp=1674826455~hmac=e2e451703c84c01719b2d4ab952bfc88a4cc96670f85bfc1735a997d9f74efd3"})`,
            backgroundSize: "auto",
            padding: 0,
            height: 500,
            marginTop: 3,
            boxShadow: 5,
          }}
        ></Box>
        <div >
          <Grid
            container
            direction="row"
            columns={6}
            spacing={3}
            sx={{ mt: 2 }}
            
          >
            {review.map((item: ReviewInterface) => (
              <Grid item xs="auto" key={item.ID}>
                <Paper>
                  <Grid>
                    <TextField
                      sx={{ "& fieldset": { border: "none" } }}
                      value={item.Customer.Email}
                    />
                  </Grid>
                  <Grid container columns={12}>
                    <Grid
                      display={"flex"}
                      justifyContent={"center"}
                      item
                      xs={5}
                    >
                      <Rating
                        value={item.Star}
                        size="medium"
                        sx={{ mt: 1.7 }}
                      />
                    </Grid>
                    <Grid item xs={7}>
                      <TextField
                        sx={{ "& fieldset": { border: "none" } }}
                        value={`${Moment(item.Reviewdate).format(
                          "DD/MM/YYYY HH:mm A"
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
                      {CheckImage(item)}
                      {/* <img  src={`${item.Reviewimage}`} onError={() => CheckImage(item)} style={myImageStyle} onDoubleClick={() => OpenImageonCilck(item)}/> */}
                    </Grid>
                  </Grid>
                </Paper>
              </Grid>
            ))}
          </Grid>
        </div>
        <Dialog fullWidth maxWidth="md" open={openImage} onClose={closeImage}>
          <DialogContent>
            <img src={`${img}`} width="100%" height="100%" />
          </DialogContent>
        </Dialog>
      </Container>
    </div>
  );
}

export default Review_Show;
