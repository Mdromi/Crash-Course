import logo from './logo.svg';
import { useState } from 'react';
import React, {Component} from 'react'
import './App.css';
import User from './User';
import User_class from './User-Class'
// import Student from './Student'

function App() {
  const [loggedIn, setLoggedIn] = useState(8)
  
  function getData() {
    alert(`Hello from app`)
  }
  return (
    <div className="App">
      <User data={getData}/>
    </div>
  );
}
export default App;

