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

