const dbConnect = require('./mongodb');

const updateData = async () => {
    let data = await dbConnect();
    let result = await data.updateOne(
        {name: 'nokia 1100 '},{
            $set:{name: 'nokia 1100', price: 550}
        }
    );
    console.table(result)
}
updateData();