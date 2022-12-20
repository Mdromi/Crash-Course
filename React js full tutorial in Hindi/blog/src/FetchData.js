import { useEffect, useState } from "react";
import Table from "react-bootstrap/Table";

function FetchData() {
  const [data, setData] = useState([]);
  let isValid = true;
  useEffect(() => {
    const fetchData = async () => {
      const data = await fetch("https://jsonplaceholder.typicode.com/users/");
      const json = await data.json();
      console.log(json);
      if (isValid) setData(json);
    };
    fetchData().catch(console.error);
  }, []);
  return (
    <div>
      <h3>Users Fetch Data</h3>
      <Table>
        <thead>
          <tr>
            <th>#</th>
            <th>Name</th>
            <th>Username</th>
            <th>Email</th>
            <th>Phone</th>
            <th>Website</th>
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
            </tr>
          ))}
        </tbody>
      </Table>
    </div>
  );
}
export default FetchData;
