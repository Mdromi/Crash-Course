const {MongoClient} = require('mongodb');

const url = 'mongodb://localhost:27017';
const database = 'e-com';
const client = new MongoClient(url);

async function dbConnect(){
    let result = await client.connect();
    console.log('Connected successfully to server');
    let db = result.db(database);
    return db.collection('product');
}

module.exports = dbConnect;