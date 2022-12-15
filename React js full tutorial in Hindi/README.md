# [React js full tutorial in Hindi]()

## <a id="table-of-contents">Table of Contents</a>

- <a href="#1"> React js and course Introduction </a> 
- <a href="#2"> Installation for react js </a> 
- <a href="#3"> Write First Code with ReactJs </a> 
- <a href="#4"> React js project File and Folder Structure </a> 
- <a href="#5"> what is package.json in react js </a> 
- <a href="#6"> Component in react </a> 
- <a href="#7"> Class component in react js </a> 
- <a href="#8"> What is jsx in react </a> 
- <a href="#9"> Click event in react js </a> 
- <a href="#10"> State in react functional component </a> 
- <a href="#11"> State in class component in react </a> 
- <a href="#12"> Props in the functional component in react </a> 
- <a href="#13"> Props in class component </a> 
- <a href="#14"> Get input box value </a> 
- <a href="#15"> Show and hide element (toggle) </a> 
- <a href="#16"> Basic form tutorial </a> 
- <a href="#17"> if-else conditions </a> 
- <a href="#18"> Pass function as props </a> 
- <a href="#19"> Introduction to react js life cycle </a> 
- <a href="#20"> Constructor life cycle method </a> 
- <a href="#21"> Render life cycle in React js </a> 
- <a href="#22"> Component did mount react js life cycle </a> 
- <a href="#23"> Component did update react js life cycle </a> 
- <a href="#24"> Should Component update in react js life cycle </a> 
- <a href="#25"> ComponentWillUnmount in react js </a> 
- <a href="#26"> Hooks in react js </a> 
- <a href="#27"> useEffect hook in react js </a> 
- <a href="#28"> useEffect with state and props in react js </a> 
- <a href="#29"> Style type in react js </a> 
- <a href="#30"> install bootstrap in react js </a> 
- <a href="#31"> Handle array with list in react js </a> 
- <a href="#32"> List with bootstrap in react js </a> 
- <a href="#33"> Nested list with a nested array in react tutorial </a> 
- <a href="#34"> How to resue component in reactJs </a> 
- <a href="#35"> Nested list with a nested array in react tutorial </a> 
- <a href="#36"> React js send data child to parent </a> 
- <a href="#37"> Pure Component in React js </a> 
- <a href="#38"> Use memo in react functional component </a> 
- <a href="#39"> Ref in React js </a> 
- <a href="#40"> useRef in react functional component </a> 
- <a href="#41"> React js forward ref </a> 
- <a href="#42"> React js Controlled Component </a> 
- <a href="#43"> React js Uncontrolled Component </a> 
- <a href="#44"> React js High order component </a> 
- <a href="#45"> React js Routing basic </a> 
- <a href="#46"> React js Routing example </a> 
- <a href="#47"> 404 page in react routing </a> 
- <a href="#48"> React js Dynamic routing </a> 
- <a href="#49"> React js Call Get Method API- HTTP call </a> 
- <a href="#50"> How to call post method API in react js </a> 
- <a href="#51"> React js delete method API" </a> 
- <a href="#52"> React js pre-filled form </a>
- <a href="#53"> React js update API method </a>
- <a href="#54"> React js get the previous state in the functional component </a>
- <a href="#55"> get previous props in react js with hooks </a>
- <a href="#56"> Context API in react js </a>
- <a href="#57"> React js interview Questions </a>


<h3 id="1"> React js and course Introduction </h3>

__What is React__
- React is a javascript Library
- The main focus is building UI as fast possible
- So this is used to single page application
- Means complete website is a single page

__Create React App__
`npx create-react-app app-name`
or
`yarn create vite`

<a href="#table-of-contents">[↑] Back to top</a>

<h3 id="6"> Component in react </h3>

__Component__
- A piece of code that can reuse
- Such as Function
- But more powerful than function
- Header, Footer is best example

__Component Type__

- Functional Component
- Class Component
- HOC (Hight order component)
- Pure Component
- Control Component
- Uncontrolled Component

<a href="#table-of-contents">[↑] Back to top</a> 

<h3 id="8"> What is jsx in react </h3>

- JavaScript xml(JSx)
- Write HTML and javascript together
- Use Like HTML or XMLtags

<a href="#table-of-contents">[↑] Back to top</a> 

<h3 id="10"> State in react functional component  </h3>

__State:__ The state is a built-in React object that is used to contain data or information about the component. A component's state can change over time; whenever it changes, the component re-renders.
```js
import { useState } from 'react';
function User() {
    let [data, setData] = useState(0)
    function updateData() {
        setData(data+1)
    }
    return (
        <div>
            <h1>functional State Component</h1>
            <p>{data}</p>
            <button onClick={updateData}>Update Data</button>
        </div>
    )
}
export default User;
```
import
```js
import User from './User';
<User/>
```
<a href="#table-of-contents">[↑] Back to top</a> 
 
<h3 id="11"> State in class component in react </h3>

``` js
import React, {Component} from 'react'
export default class User extends Component
{
    constructor()
    {
        super();
        this.state = {
            data: 0
        }
    }
    updateData()
    {
       this.setState({data:this.state.data+1})
    }
    render() 
    {
        return (
            <div className='data'>
                <h1>state in class component in react</h1>
                <p>{this.state.data}</p>
                <button onClick={()=>this.updateData()}>Update Data</button>
            </div>
        )
    }
}
```
import
```js
import User_class from './User-Class'
<User_class/>
```

<a href="#table-of-contents">[↑] Back to top</a>

<h3 id="12"> Props in Functional component </h3>

__Props:__ Props are `used to store data that can be accessed by the children of a React component.` They are part of the concept of reusability. Props take the place of class attributes and allow you to create consistent interfaces across the component hierarchy. In this article, we'll learn about props in React.

Example `Student.js`
```js
function Student(props)
{
    return(
        <div>
            <h1>Student Component</h1>
            <p>Name: {props.details.name}</p>
            <p>Email: {props.details.email}</p>
            <p>Level: {props.level}</p>
        </div>
    )
}
export default Student
```
import `App.js`
```js
import { useState } from 'react';
import Student from './Student'

<div className="App">
    <h1>Props in react :)</h1>
    <Student details={{name, email:"mdromi@gmail.com"}} level="A"/>
    <button onClick={()=>{setName("Md Romi")}}>Update Name</button>
</div>
```
<a href="#table-of-contents">[↑] Back to top</a>

<h3 id="13"> Props in class component </h3>

Example `Student.js`
```js
import React, {Component} from 'react'
export default class Student extends Component
{
    render()
    {
        return(
            <div>
                <h1>Student Class Component</h1>
                <p>Name: {this.props.name}</p>
                <p>Email: {this.props.email}</p>
                <p>Level: {this.props.level}</p>
            </div>
        )
    }
}
```
import `App.js`
```js
import React, {Component} from 'react'
import Student from './Student'
class App extends Component {
  constructor()
  {
    super();
    this.state={
      name: 'Md',
      email:"mdromi@gmail.com",
      level:"A"
    }
  }
  render()
  {
    return (
      <div className="App">
      <h1> My First react App</h1>
      <h1>Props in react :)</h1>
      <Student name={this.state.name} email={this.state.email}  level={this.state.level} />
      <button onClick={()=>this.setState({name:"Md Romi"})}>Update Namer</button>
    </div>
    )
  }
}
export default App;
```

<a href="#table-of-contents">[↑] Back to top</a>

<h3 id="14"> Get input box value </h3>

- Make input filled 
- Make function and get value
- Make button and state
- Get value on button click

Example `app.js`
```js
import { useState } from 'react';
import './App.css';

function App() {
  const [data, setData] = useState(null)
  const [print, setPrint] = useState(false)
  
  function getData(val) {
    setData(val.target.value);
    setPrint(false)
  }
  return (
    <div className="App">
      <h1> Get Input box </h1>
      {
        print ? <h3>{data}</h3> : null
      }
      <input type="text" placeholder="anything..." onChange={getData}></input>
      <button onClick={()=>setPrint(true)}>Print Data</button>
    </div>
  );
}
export default App;
```

<a href="#table-of-contents">[↑] Back to top</a>

<h3 id="15"> Show and hide element (toggle) </h3>

- Make a div and any other element
- Make state for hiding and show element
- Make button
- Update state on button click

Example `app.js`
```js
import { useState } from 'react';
function App() {
  const [status, setStatus] = useState(true)
  
  return (
    <div className="App">
      {
        status? <h1> Show or Hide </h1> : null
      }
      <button onClick={()=>setStatus(true)}>Show</button>
      <button onClick={()=>setStatus(false)}>Hide</button>
      <button onClick={()=>setStatus(!status)}>Toggle</button>
    </div>
  );
}
export default App;
```

<a href="#table-of-contents">[↑] Back to top</a>
<h3 id="16"> Basic form tutorial </h3>

- Make a html form
- Add some field
- Use state for field data
- Submit and control form

Example `app.js`
```js
import { useState } from 'react';

function App() {
  const [name, setName] = useState("")
  const [interest, setInterest] = useState("")
  const [tnc, setTnc] = useState(false)
  
  function getFormData(e) {
    console.log(name, interest, tnc);
    e.preventDefault()
  }
  
  return (
    <div className="App">
      <h1>Handle Form in react</h1>
      <form onSubmit={getFormData}>
        <input type="text" placeholder="Enter name" value={name} onChange={(e)=>setName(e.target.value)}/> <br/><br/>
        <select onChange={(e)=>setInterest(e.target.value)}>
            <option>Select Option</option>
            <option>Marvel</option>
            <option>Dc</option>
        </select><br/><br/>
        <input type="checkbox" onChange={(e)=>setTnc(e.target.checked)} /> 
        <span>Accept Terms and condition</span> <br/><br/>
        <button type="submit">Submit</button>
        <button type="reset">Clear</button>
      </form>
    </div>
  );
}

export default App;
```

<a href="#table-of-contents">[↑] Back to top</a>

<h3 id="17"> if-else conditions </h3>

- Make Component
- Use State 
- Not Recommended way for condition
- Correct condition

Example `app.js`
```js
import { useState } from 'react';

function App() {
  const [loggedIn, setLoggedIn] = useState(8)

  return (
    <div className="App">
      {
        loggedIn==1? <h1>Welcome User {loggedIn}</h1>
        :loggedIn==2? <h1>Welcome User {loggedIn}</h1>
        :<h1>Welcome User {loggedIn}</h1>
      }
      
    </div>
  );
}
export default App;
```
<a href="#table-of-contents">[↑] Back to top</a>

<h3 id="18"> Pass function as props </h3>

Example `User.js`
```js
import { useState } from 'react';
function User(props) {
    let [data, setData] = useState(0)

    function updateData() {
        setData(data+1)
    }
    return (
        <div>
            <h1>Pass function as props</h1>
            <p>{data}</p>
            <button onClick={props.data}>Call Data Function</button>
        </div>
    )
}
export default User;
```
import `App.js`
```js
function App() {
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
```
<a href="#table-of-contents">[↑] Back to top</a>

<h3 id="19"> 
    <a href="https://programmingwithmosh.com/javascript/react-lifecycle-methods/">Introduction to react js life cycle 
    </a>
 </h3>

__When Life Cycle method work__
- When Component is loaded
- When Component is updated with state and props
- When Component is removed

__Life Cycle Component 3 State:__ 
- Mounting – Birth of your component
- Update – Growth of your component
- Unmount – Death of your component
<div align="center">
    <a href="https://programmingwithmosh.com/javascript/react-lifecycle-methods/">
    	<img src="./react-life-cycle.webp" height="500" alt="react-life-cycle">
    </a>
</div>

<a href="#table-of-contents">[↑] Back to top</a>

<h3 id=""> </h3>
<h3 id=""> </h3>
<h3 id=""> </h3>
<h3 id=""> </h3>
<h3 id=""> </h3>
<h3 id=""> </h3>
<h3 id=""> </h3>
<h3 id=""> </h3>

<a href="#table-of-contents">[↑] Back to top</a> 

