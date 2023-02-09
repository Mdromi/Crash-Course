const xsrURL = "https://jsonplaceholder.typicode.com/users/";

function myAxios(method, url, body = null) {
  return new Promise((resolve, reject) => {
    const xhr = new XMLHttpRequest();

    xhr.open(method, url);
    xhr.responseType = "json";
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.onload = () => {
      if (xhr.status >= 400) reject(xhr.response);
      resolve(xhr.response);
    };
    xhr.onerror = () => {
        reject(xhr.response);
    };

    xhr.send(JSON.stringify(body));
  });
}

myAxios("GET", xsrURL)
.then((res) => console.log('Message from then block(GET)',res))
.catch((err) => console.log(err))

myAxios("POST", xsrURL, {
  name: "ROMI",
  job: "Student",
  age: 18,
})
.then((res) => console.log('Message from then block (POST)',res))
.catch((err) => console.log(err))
