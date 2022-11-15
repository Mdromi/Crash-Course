/*====================
22.Dynamic Page
CMD: [ node 22.DynamicPage/index.js ]
===================*/
const { response } = require('express');
const express = require('express');
const path = require('path'); 
const ejs = require('ejs');

const app = express();
const port = 3600; 
const publicPath = path.join(__dirname, 'public') ;

app.set('view engine', 'ejs');

app.get('', (_, resp) => {
    resp.sendFile(`${publicPath}/index.html`)
});

app.get('/profile', (_, resp) => {
    const user = {
        name: 'MD Romi',
        email: 'mdromi@gmail.com',
        city: 'Faridpur',
        skills: ['php', 'js', 'c++']
    }
    resp.render('profile', {user});
});

app.get('/login', (_, resp) => {
    resp.render('login')
})


app.get('*', (_, resp) => {
    resp.sendFile(`${publicPath}/404.html`)
});

app.listen(port, function() {
    console.log('Server listening on http://localhost:' + port);
});