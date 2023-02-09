const URL = "https://jsonplaceholder.typicode.com/users/";

function myAxios(method, url, body = null) {
    const headers ={"Content-Type": "application/json"}
    return fetch(url, {
        method: method,
        headers: headers,
        body: JSON.stringify(body)
    }).then((res)=> res.json())
}

myAxios("GET", URL)
.then((res) => console.log('Message from Axios then block(GET)',res))
.catch((err) => console.log(err))

myAxios("POST", URL, {
    name: "ROMI",
    job: "Student",
    age: 18,
  })
  .then((res) => console.log('Message from Axios then block (POST)',res))
  .catch((err) => console.log(err))
  