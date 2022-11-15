/*====================
28.Connect MongoDB with Node
CMD: [ node 28.ConnectNodeMongo/index.js ]
===================*/
const {MongoClient} = require('mongodb');
/* const MongoClient = require('mongodb').MongoClient; */

const url = 'mongodb://localhost:27017';
const database = 'e-com';
const client = new MongoClient(url);

async function getData(){
    let result = await client.connect();
    console.log('Connected successfully to server');
    let db = result.db(database);
    let collection = db.collection('product');
    let response = await collection.find({}).toArray();
    console.log(response); 
}

getData();