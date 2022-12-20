import { useEffect, useState } from "react";
import Table from "react-bootstrap/Table";

function FetchData() {
  const [data, setData] = useState([]);
  const [name, setName] = useState("");
  const [username, setUserName] = useState("");
  const [email, setEmail] = useState("");
  const [phone, setPhone] = useState("");
  const [website, setWebsite] = useState("");
  let isValid = true;
  useEffect(() => {
    getList()
  }, []);

  function getList() {
    const fetchData = async () => {
      const data = await fetch("https://jsonplaceholder.typicode.com/users/");
      const json = await data.json();
      if (isValid) setData(json);
    };
    fetchData().catch(console.error);
  }

  function selectUser(id) {
   let item = data[id-1]
   setName(item.name)
   setUserName(item.username)
   setEmail(item.email)
   setPhone(item.phone)
   setWebsite(item.website)
  }

  return (
    <div>
      <h3>Users Fetch Data</h3>
      <form>
        <input type="text" value={name} onChange={(e)=>setName(e.target.value)} ></input> <br /> <br />
        <input type="text" value={username} onChange={(e)=>setUserName(e.target.value)} ></input> <br /> <br />
        <input type="text" value={email} onChange={(e)=>setEmail(e.target.value)} ></input> <br /> <br />
        <input type="text" value={phone} onChange={(e)=>setPhone(e.target.value)} ></input> <br /> <br />
        <input type="text" value={website} onChange={(e)=>setWebsite(e.target.value)} ></input> <br /> <br />
      </form>
      <Table>
        <thead>
          <tr>
            <th>#</th>
            <th>Name</th>
            <th>Username</th>
            <th>Email</th>
            <th>Phone</th>
            <th>Website</th>
            <th>Operation</th>
          </tr>
        </thead>
        <tbody>
          {data.map((item) => (
            <tr key={item.id}>
              <td>{item.id}</td>
              <td>{item.name}</td>
              <td>{item.username}</td>
              <td>{item.email}</td>
              <td>{item.phone}</td>
              <td>{item.website}</td>
              <td><button onClick={()=> selectUser(item.id)}>Update</button></td>
            </tr>
          ))}
        </tbody>
      </Table>
    </div>
  );
}
export default FetchData;
