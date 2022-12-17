import logo from './logo.svg';
import { Button } from 'react-bootstrap';
import { useState, useEffect } from 'react';
import React, {Component} from 'react'
import './App.css';
import User from './User';
import User_class from './User-Class'
// import Student from './Student'
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";

function App() {
  const [data, setData] = useState(0)
  
  function getData() {
    alert(`Hello from app`)
  }
  return (
    <div className="App">
      <h1>Render Method </h1>
      <User data="1234"/>
      {/* <button onClick={props.data}>Call Data Function</button> */}
    </div>
  );
}
export default App;



