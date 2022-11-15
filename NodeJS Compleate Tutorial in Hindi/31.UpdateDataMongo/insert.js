const dbConnect = require('./mongodb');

const insertData = async () => {
    const db = await dbConnect();
    const result = await db.insertMany(
        [
            {name: 'max 1', brand: 'micromax', price: 320, catagory: 'mobile'},
            {name: 'max 2', brand: 'micromax', price: 320, catagory: 'mobile'},
            {name: 'max 3', brand: 'micromax', price: 320, catagory: 'mobile'}
        ]
    );

    if(result.acknowledged){
        console.log("Data Inserted");
    }
}
insertData();
