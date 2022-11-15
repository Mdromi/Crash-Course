module.exports = (req, resp, next) => {
    console.log('reqFilter');
    if(!req.query.age){
        resp.send('Please Provide Age!');
    }
    else if(req.query.age < 18){
        resp.send('You can not access this page');
    }
    else{
        next();
    }
    
}