/*====================
37.Node Js Mongoose
CMD: [ node 37.NodeJsMongoose/index.js ]
===================*/
const mongoose = require('mongoose');

const main = async () => {
    await mongoose.connect('mongodb://localhost:27017/e-com');
    const ProductSchema = new mongoose.Schema({
        name: String,
        price: Number
    });
    const ProductModel = mongoose.model('product', ProductSchema);
    let data = new ProductModel({name: "n8", price: 1000});
    let result = await data.save();
    console.log(result)
}

main();