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