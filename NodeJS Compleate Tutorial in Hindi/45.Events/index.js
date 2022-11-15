/*====================
45.Events and event emitter in nodejs
CMD: [ node 45.Events/index.js ]
===================*/

const express = require('express');
const EventEmitter = require('events');
const app = express();
const port = 5000;
const event = new EventEmitter();

let count = 0;

event.on('connectAPI', () => {
    count++;
    console.log('event called ', count);
})

app.get('/', (req, resp) => {
    resp.send('api called');
    event.emit('connectAPI');
});

app.get('/search', (req, resp) => {
    resp.send('search api called');
    event.emit('connectAPI');
});

app.get('/update', (req, resp) => {
    event.emit('connectAPI');
    resp.send('update api called');
});



app.listen(port, function() {
    console.log('Server listening on http://localhost:' + port);
});