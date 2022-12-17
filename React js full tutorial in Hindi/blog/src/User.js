import { useState } from 'react';
import React, {Component} from 'react'

class User extends Component {
    constructor() {
        super();
        this.state = {
            email: `mdr@gmail.com`
        }
    }
    render() {
        console.log(`Render Method`, this.state.email);
        return(
            <div>
                <h1>User Component</h1>
                <button onClick={()=>this.setState({email: `123@gmail.com`})}>Call Data Function</button>
            </div>
        )
    }
}
export default User;

// function User(props) {
//     let [data, setData] = useState(0)

//     function updateData() {
//         setData(data+1)
//     }
//     return (
//         <div>
//             <h1>Pass function as props</h1>
//             <p>{data}</p>
//             <button onClick={props.data}>Call Data Function</button>
//         </div>
//     )
// }

