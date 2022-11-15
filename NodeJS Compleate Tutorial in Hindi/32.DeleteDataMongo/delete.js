const dbConnect = require('./mongodb');

const deleteData = async () => {
    let data = await dbConnect();
    let result = await data.deleteOne({name: 'max 2'});
    if(data.acknowledged){
        console.log('Record Deleted');
    }
    console.table(result)
}
deleteData();