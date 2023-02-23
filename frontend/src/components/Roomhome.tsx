import * as React from "react";
import Carousel from "react-material-ui-carousel";
import im10 from "../Image/im10.jpg"
import im7 from "../Image/im7.jpg"
import im9 from "../Image/im9.png"
import Logo1 from "../Image/LOGO.png"

import { useEffect, useState } from 'react';

import Toolbar from "@mui/material/Toolbar";

import Typography from "@mui/material/Typography";

import { Link as RouterLink } from "react-router-dom";

import { createTheme, styled, useTheme } from "@mui/material/styles";
import { grey } from "@mui/material/colors";
import MuiAppBar, { AppBarProps as MuiAppBarProps } from '@mui/material/AppBar';
import CssBaseline from "@mui/material/CssBaseline";
import { AppBar, Box, Button, Grid, ThemeProvider } from "@mui/material";

const bgnavbar = createTheme({
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

const drawerWidth = 320; //ความยาวของ แถบเมนู

const Main = styled('main', { shouldForwardProp: (prop) => prop !== 'open' })<{
  open?: boolean;
}>(({ theme, open }) => ({
  // flexGrow: 1,
  // padding: theme.spacing(3),
  transition: theme.transitions.create('margin', {
    easing: theme.transitions.easing.sharp,
    duration: theme.transitions.duration.leavingScreen,
  }),
  marginLeft: `-${drawerWidth}px`,
  ...(open && {
    transition: theme.transitions.create('margin', {
      easing: theme.transitions.easing.easeOut,
      duration: theme.transitions.duration.enteringScreen,
    }),
    marginLeft: 0,
  }),
}));

interface AppBarProps extends MuiAppBarProps {
  open?: boolean;
}

// const AppBar = styled(MuiAppBar, {
//   shouldForwardProp: (prop) => prop !== 'open',
// })<AppBarProps> (({ theme, open }) => ({
//   transition: theme.transitions.create(['margin', 'width'], {
//     easing: theme.transitions.easing.sharp,
//     duration: theme.transitions.duration.leavingScreen,
//   }),
//   ...(open && {
//     width: `calc(100% - ${drawerWidth}px)`,
//     marginLeft: `${drawerWidth}px`,
//     transition: theme.transitions.create(['margin', 'width'], {
//       easing: theme.transitions.easing.easeOut,
//       duration: theme.transitions.duration.enteringScreen,
//     }),
//   }),
// }));


function Roomhome() {
  const themep = useTheme();
  const [open, setOpen] = React.useState(false);

  const handleDrawerOpen = () => {
    setOpen(true);
  };

  const handleDrawerClose = () => {
    setOpen(false);
  };
  
  function Item(props: any) {
    return (
         <img src={props.item.Image} width= "100%" height="600px"/>
      );
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
        {Slider.map((item, i) => (<Item key={i} item={item} />))}
      </Carousel>
    );
  }

  return (
  <ThemeProvider theme={bgnavbar}>
      <AppBar position="fixed">
        <Toolbar>
          <div>
            <img src={Logo1} width= "75px" height="75px"/>
          </div>
          <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', width: '100%'}}>
            <Typography variant="h6" color="secondary" noWrap component="div" marginLeft={2}>
              <div >
                G03 Hotel
              </div>
            </Typography>
          </Box>
          <Button component={RouterLink} to="/Homeshow"  color="secondary" >Home</Button>
          <Button component={RouterLink} to="/Roomhome"  color="secondary" >Room</Button>
          <Button component={RouterLink} to="/SignIn"  color="secondary" >ABOUT</Button>
          <Button component={RouterLink} to="/SignIn" variant="contained" color="secondary" >BOOK NOW</Button>
          
        </Toolbar>

      </AppBar>
  <div>
     <Grid>
      {ImageC()}
     </Grid >
     <div className="grid-con2">
        <div className="grid-item-21">
          <img src="https://cf.bstatic.com/xdata/images/hotel/max1280x900/19061811.jpg?k=d8542777976a7e63a981e6a0fef321472713bc27d6454055df45265eee200536&o=&hp=1" />
        </div>
     </div>
     <div className="grid-con">
        <div className="grid-item-1">
          <img src="https://content.r9cdn.net/himg/80/b3/46/expediav2-13266-f26277-844954.jpg" />
        </div>
        <div className="grid-item-2">
          <div className="grid-item-info">
            <div>
              <h3>Handmade piece</h3>
              <p>Welcome to Hotel G03</p>
            </div>
          </div>
          <div>
            <img src="https://files.guidedanmark.org/files/382/242733_Hotel_Ottilia__Brchner_Hotels_PR.jpg"></img>
          </div>
        </div>
        <div className="grid-item-3">
          <div>
            <img src="https://www.nzbusinesstraveller.co.nz/wp-content/uploads/2022/08/Copy-of-Chairman-Suite-Lounge-1024x683.jpg"></img>
          </div>
          <div className="grid-item-info2">
            <div>
              <h3>wow</h3>
              <p>Welcome to Hotel G03</p>
            </div>
          </div>
        </div>
        <div className="grid-item-4">
          <div>
            <img src="https://assets.langhamhotels.com/is/image/langhamhotelsstage/cdakl-rooms-chairman-suite-bedroom-1680-945:Medium?wid=1680&hei=944"></img>
          </div>
          <div className="grid-item-info">
            <div>
              <h3>wowhhghgh</h3>
              <p>Welcome to Hotel G03</p>
            </div>
          </div>
        </div>
        <div className="grid-item-5">
          <div>
            <img src="https://www.thedenizen.co.nz/wp-content/uploads/2021/01/Opera-Suite-Lounge-1.jpg"></img>
          </div>
        </div>
    </div>
  </div>
</ThemeProvider>
  );
}

export default Roomhome;