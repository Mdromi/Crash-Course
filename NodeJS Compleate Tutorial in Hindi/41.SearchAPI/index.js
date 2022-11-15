/*====================
41.Search api with nodejs and mongoose
CMD: [ node 41.SearchAPI/index.js ]
===================*/
const express = require('express');
require('./config');
const product = require('./product');
const app = express();
const port = 5000;

app.use(express.json());

app.get('/search/:key', async (req, resp) => {
    console.log(req.params.key);
    let data = await product.find(
        {
            "$or": [
                {"name": {$regex: req.params.key}},
                {"brand": {$regex: req.params.key}},
                {"catagory": {$regex: req.params.key}}
            ]
        }
    );
    resp.send(data);
});


app.listen(port, function() {
    console.log('Server listening on http://localhost:' + port);
});