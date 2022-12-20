import './App.css';
// import Student from './Student'
import Routers from "./Routers";


function App() {
  // const [data, setData] = useState([])
  
  function getData() {
    alert(`Hello from app`)
  }
  return (
    <div className="App">
     <Routers/>
    </div>
  );
}
export default App;



