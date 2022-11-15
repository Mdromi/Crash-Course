/*====================
42.Upload File with Node JS
CMD: [ node 42.UploadFileNodeJs/index.js ]
===================*/

const express = require('express');
const multer  = require('multer');
const app = express();
const port = 5000;

const upload = multer({
    storage: multer.diskStorage({
        destination: function(req, file, cb){
            cb(null, "uploads");
        },
        filename: function(req, file, cb){
            cb(null, file.fieldname + "-" + Date.now() + ".jpg");
        }
    })
}).single("file_name");



app.post('/upload', upload, (req, resp) => {
    console.log(upload);
    resp.send('File Uploded'); 
});

app.listen(port, function() {
    console.log('Server listening on http://localhost:' + port);
});