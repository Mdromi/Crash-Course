/*====================
21.Remove Extension from URL
CMD: [ node 21.RemoveExtension/index.js ]
===================*/

const express = require('express');
const path = require('path');

const app = express();
const port = 3000; 
const publicPath = path.join(__dirname, 'public') ;

// app.use(express.static(publicPath));
app.get('', (_, resp) => {
    resp.sendFile(`${publicPath}/index.html`)
});

app.get('/about', (_, resp) => {
    resp.sendFile(`${publicPath}/about.html`)
});

app.get('/help', (_, resp) => {
    resp.sendFile(`${publicPath}/help.html`)
});

app.get('*', (_, resp) => {
    resp.sendFile(`${publicPath}/404.html`)
});

app.listen(port, function() {
    console.log('Server listening on http://localhost:' + port);
});