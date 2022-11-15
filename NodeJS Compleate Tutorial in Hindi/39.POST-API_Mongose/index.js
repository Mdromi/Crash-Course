/*====================
39.POST API With Mongose
CMD: [ node 39.POST-API_Mongose/index.js ]
===================*/
const express = require('express');
require('./config');
const product = require('./product');
const app = express();
const port = 5000;

app.use(express.json());

// create DB
app.post('/create', async (req, resp) => {
    let data = new product(req.body);
    let result = await data.save();
    console.log(req.body);
    resp.send(result);
});

// read DB
app.get('/list', async (req, resp) => {
    let data = await product.find();
    console.log(data);
    resp.send(data);
});

// delete DB
app.delete('/delete/:_id', async (req, resp) => {
    console.log(req.params);
    let data = await product.deleteOne(req.params);
    resp.send(data);
});

// update DB
app.put('/update/:_id', async (req, resp) => {
    console.log(req.params);
    let data = await product.updateOne(
        req.params,
        {
            $set: req.body
        }
    );
    resp.send(data);
});


app.listen(port, function() {
    console.log('Server listening on http://localhost:' + port);
});