/*====================
23.Middleware
CMD: [ node 23.Middleware/index.js ]
===================*/

const express = require('express');
const app = express();
const port = 5000;

const reqFilter = (req, resp, next) => {
    console.log('reqFilter');
    if(!req.query.age){
        resp.send('Please Provide Age!');
    }
    else if(req.query.age < 18){
        resp.send('You can not access this page');
    }
    else{
        next();
    }
    
}

app.use(reqFilter);

app.get('/',(req, resp) =>{
    resp.send('Welcome To home Page');
})

app.get('/users',(req, resp) =>{
    resp.send('Welcome To Users Page');
})

app.listen(port);