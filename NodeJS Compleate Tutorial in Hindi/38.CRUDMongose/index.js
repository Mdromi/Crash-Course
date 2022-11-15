/*====================
38.CRUD operations with Mongoose
CMD: [ node 38.CRUDMongose/index.js ]
===================*/
const mongoose = require('mongoose');

mongoose.connect('mongodb://localhost:27017/e-com');
const ProductSchema = new mongoose.Schema({
    name: String,
    price: Number,
    brand: String,
    catagory: String
});


const saveInDB = async () => {
    const ProductModel = mongoose.model('product', ProductSchema);
    let data = new ProductModel({
        name: "note2 pro", 
        price: 2000, 
        brand: "xamy", 
        catagory: "mobile"
    });
    let result = await data.save();
    console.log(result)
}
// saveInDB() // Create DataBase & save

const updateInDB = async () => {
    const Product = mongoose.model('product', ProductSchema);
    let data = await Product.updateMany(
        {name: 'note pro'},
        {
            $set: {price: 700, name: 'max 8'}
        }
    );
    console.log(data)
}
// updateInDB(); // updade Data Base

const deleteDB = async () => {
    const Product = mongoose.model('product', ProductSchema);
    let data = await Product.deleteOne({name: 'max 8'})
    console.log(data)
}
// deleteDB(); // delete Data Base

const findInDB = async () => {
    const Product = mongoose.model('product', ProductSchema);
    let data = await Product.find({name: 'max 8'})
    console.log(data)
}

findInDB();

console.log()