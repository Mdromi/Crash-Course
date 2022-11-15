/*====================
27.CRUD Opperation With Mongo
CMD: [ node 27.CRUDOpperationMongo/index.js ]
===================*/

const express = require('express');
const reqFilter = require('./middleware');
const app = express();
const route = express.Router();
const port = 5000;


// app.use(reqFilter);

route.use(reqFilter);
app.get('/',(req, resp) =>{
    resp.send('Welcome To home Page');
})

app.get('/users', reqFilter, (req, resp) =>{
    resp.send('Welcome To Users Page');
});

route.get('/about',(req, resp) =>{
    resp.send('Welcome To About Page');
});

route.get('/contact',(req, resp) =>{
    resp.send('Welcome To Contact Page')
});

app.use('/', route);

app.listen(port);