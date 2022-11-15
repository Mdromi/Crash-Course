const express = require('express');
const dbConnect = require('./mongodb');
const mongoDb = require('mongodb');
const app = express();
const port = 5000;

app.use(express.json());

// Get Data
app.get('/', async (req, resp) => {
    let data = await dbConnect();
    data = await data.find().toArray();
    // console.log(data)
    resp.send(data);
});

// Insert Data
app.post('/', async (req, resp) => {
    // console.log(req.body)
    let data = await dbConnect();
    let result = await data.insertOne(req.body)

    // resp.send(data)
    resp.send(result);
});

// Update Data
app.put('/:name', async (req, resp) => {
    // console.log(req.body)
    let data = await dbConnect();
    let result = await data.updateOne(
        {name: req.params.name},
        {$set: req.body}
    )
    resp.send({result: 'Update'})
});

// Delete Data
app.delete('/:id', async (req, resp) => {
    // console.log(req.params.id);
    let data = await dbConnect();
    const result = await data.deleteOne({__id: new mongoDb.ObjectId(req.params.id)});
    resp.send(result);
})

app.listen(port, function() {
    console.log('Server listening on http://localhost:' + port);
});