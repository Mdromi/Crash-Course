/*====================
19.Render HTML and JSON
CMD: [ node 19.RenderHTMLJSON/index.js ]
===================*/
const express = require('express');
const app = express();
const port = 4000;

app.get('/', function(request, response) {
    // console.log('data sent by browser =>>>', request.query.name)
    response.send(`
        <h1>Welcome to home page</h1>
        <a href= "/about">About Page</a>
        <a href= "/">Home Page</a>
    `);
});

app.get('/about', function(request, response) {
    response.send(`
        <input type="text" placeholder="User name" value=" ${request.query.name}" />
        <button>Click Me </button>
    `);
});

app.get('/help', function(request, response) {
    response.send([
        {
            name: 'Md Romi',
            email: 'mdromi@gmail.com'
        },
        {
            name: 'Rakib Ahmed',
            email: 'rakibahmed@gmail.com'
        }
    ]);
});

app.listen(port, function() {
    console.log('Server listening on http://localhost:' + port);
});