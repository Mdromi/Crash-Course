import { Routes, Route } from "react-router-dom";
import Nav from "./Nav";
import Home from "./Home";
import About from "./About"; 
import Users from "./Users"; 
import NotFound from "./NotFound"; 

function Routers() {
    return(
        <header>
        <Nav/>
        <Routes>
            <Route path="/about" element={<About/>} /> 
            <Route path="/users" element={<Users/>} /> 
            <Route path="/" exact={true}  element={<Home/>} />
            <Route path="*" exact={true}  element={<NotFound/>} />
        </Routes>
        </header>
    )
}
export default Routers