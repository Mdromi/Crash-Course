/*====================
31.Update Data From MongoDB
CMD: [ node 31.UpdateDataMongo/index.js ]
===================*/
const dbConnect = require('./mongodb')

// dbConnect().then((resp) => {
//     resp.find().toArray().then((data) => {
//         console.log(data)
//     })
// });
// console.log(dbConnect())

const main = async () => {
    let data = await dbConnect();
    data = await data.find().toArray();
    console.warn(data)
}
main();