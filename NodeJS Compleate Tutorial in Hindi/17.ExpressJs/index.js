/*====================
14.Asynchoronous Basics in NodeJs
CMD: [ node 17.ExpressJs/index.js ]
===================*/

// Import the top-level function of express
const express = require('express');

// Creates an Express application using the top-level function
const app = express();

// Define port number as 3000
const port = 4000;

// Routes HTTP GET requests to the specified path "/" with the specified callback function
app.get('/', function(request, response) {
    console.log('data sent by browser =>>>', request.query.name)
    response.send('Hello, World!' + request.query.name);
});

app.get('/about', function(request, response) {
    response.send('This is about us Page');
});

app.get('/help', function(request, response) {
    response.send('This is help us Page');
});

// Make the app listen on port 3000
app.listen(port, function() {
    console.log('Server listening on http://localhost:' + port);
});