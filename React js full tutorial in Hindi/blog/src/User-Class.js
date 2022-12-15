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