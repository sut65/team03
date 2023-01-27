import * as React from "react";

import Box from "@mui/material/Box";

import Toolbar from "@mui/material/Toolbar";

import Typography from "@mui/material/Typography";

import IconButton from "@mui/material/IconButton";
import { Link as RouterLink } from "react-router-dom";

import MenuIcon from "@mui/icons-material/Menu";
import { createTheme, styled, useTheme } from "@mui/material/styles";
import { grey } from "@mui/material/colors";
import MuiAppBar, { AppBarProps as MuiAppBarProps } from '@mui/material/AppBar';
import CssBaseline from "@mui/material/CssBaseline";
import { Divider, Drawer, ListItem, ListItemIcon, ListItemText, MenuItem, ThemeProvider } from "@mui/material";

import LogoutIcon from '@mui/icons-material/Logout';
import ChevronLeftIcon from '@mui/icons-material/ChevronLeft';
import ChevronRightIcon from '@mui/icons-material/ChevronRight';
import HandymanIcon from '@mui/icons-material/Handyman';
import CheckroomIcon from '@mui/icons-material/Checkroom';
import ManageAccountsIcon from '@mui/icons-material/ManageAccounts';
import ReviewsIcon from '@mui/icons-material/Reviews';
import WarehouseIcon from '@mui/icons-material/Warehouse';
import ReceiptLongIcon from '@mui/icons-material/ReceiptLong';
import RoomServiceIcon from '@mui/icons-material/RoomService';
import PaymentIcon from '@mui/icons-material/Payment';
import PriceCheckIcon from '@mui/icons-material/PriceCheck';
import BedroomParentIcon from '@mui/icons-material/BedroomParent';
import PersonAddIcon from '@mui/icons-material/PersonAdd';
import FactCheckIcon from '@mui/icons-material/FactCheck';
import HomeIcon from '@mui/icons-material/Home';

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
  flexGrow: 1,
  padding: theme.spacing(3),
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

const AppBar = styled(MuiAppBar, {
  shouldForwardProp: (prop) => prop !== 'open',
})<AppBarProps> (({ theme, open }) => ({
  transition: theme.transitions.create(['margin', 'width'], {
    easing: theme.transitions.easing.sharp,
    duration: theme.transitions.duration.leavingScreen,
  }),
  ...(open && {
    width: `calc(100% - ${drawerWidth}px)`,
    marginLeft: `${drawerWidth}px`,
    transition: theme.transitions.create(['margin', 'width'], {
      easing: theme.transitions.easing.easeOut,
      duration: theme.transitions.duration.enteringScreen,
    }),
  }),
}));

const DrawerHeader = styled('div')(({ theme }) => ({
  display: 'flex',
  alignItems: 'center',
  padding: theme.spacing(0, 1),
  // necessary for content to be below app bar
  ...theme.mixins.toolbar,
  justifyContent: 'flex-end',
}));

function Navbar() {

  const themep = useTheme();
  const [open, setOpen] = React.useState(false);

  const handleDrawerOpen = () => {
    setOpen(true);
  };

  const handleDrawerClose = () => {
    setOpen(false);
  };

  const SignOut = () => {
    localStorage.clear();
    window.location.href = "/";
  }

  const menu = [
    { name: "หน้าหลัก", icon: <HomeIcon  />, path: "/" },
    { name: "Booking", icon: <ReceiptLongIcon  />, path: "/Book" },
    { name: "Repair", icon: <HandymanIcon   />, path: "/Rep" },
    { name: "Subscribe", icon: <PersonAddIcon  />, path: "/Sub" },
    { name: "Payment", icon: <PaymentIcon  />, path: "/PM" },
    { name: "Review", icon: <ReviewsIcon />, path: "/RW" },
    { name: "Service", icon: <RoomServiceIcon  />, path: "/ss" },
    
    { name: "Check IN - Check Out", icon: <FactCheckIcon  />, path: "/CNCO" },
    { name: "Check Payment", icon: <PriceCheckIcon  />, path: "/CPM" },
    { name: "Check The Room", icon: <CheckroomIcon  />, path: "/CTR" },
    { name: "Room Information", icon: <BedroomParentIcon />, path: "/RT" },
    { name: "Room Warehouse", icon: <WarehouseIcon  />, path: "/RoomW" },

    { name: "Manage Employee Information", icon: <ManageAccountsIcon  />, path: "/Manage-Show" },
  ]


 return (
  <ThemeProvider theme={bgnavbar}>
    <Box sx={{ display: 'flex' }} >
      <CssBaseline />
      <AppBar position="fixed" open={open}>
        <Toolbar>
          <IconButton
            color="secondary"
            aria-label="open drawer"
            onClick={handleDrawerOpen}
            edge="start"
            sx={{ mr: 2, ...(open && { display: 'none' }) }}
          >
            <MenuIcon />
          </IconButton>
          <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', width: '100%'}}>
            <Typography variant="h6" color="secondary" noWrap component="div">
              G03 Hotel
            </Typography>
            <MenuItem onClick={SignOut}><LogoutIcon style={{ marginRight: ".5rem" }}/>Log out</MenuItem>
          </Box>
          
        </Toolbar>

      </AppBar>

      <Drawer
        sx={{
          width: drawerWidth,
          flexShrink: 0,
          '& .MuiDrawer-paper': {
            width: drawerWidth,
            boxSizing: 'border-box',
          },
        }}
        variant="persistent"
        anchor="left"
        open={open}
      >
        <DrawerHeader>
          {/* ปุ่มกด < */}
          <IconButton onClick={handleDrawerClose}> 
            {themep.direction === 'ltr' ? <ChevronLeftIcon /> : <ChevronRightIcon />}
          </IconButton> {/* ปุ่มกด < */}
        </DrawerHeader>

        <Divider />

         {menu.map((item, index) => (
                <ListItem key={index} button component={RouterLink} onClick={handleDrawerClose}
                 to={item.path}>
                  <ListItemIcon>{item.icon}</ListItemIcon>
                  <ListItemText>{item.name}</ListItemText>
                  
                </ListItem>
              ))}

        <Divider />
      </Drawer>

      <Main open={open}>
        <DrawerHeader />
      </Main>

    </Box>
  </ThemeProvider>

 );

}

export default Navbar;