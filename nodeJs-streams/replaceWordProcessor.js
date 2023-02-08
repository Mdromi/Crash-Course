const {Transform, pipeline} = require('stream');

const replaceWordProcessing = new Transform({
    transform(chunk, encoding, callback) {
        // replaceWordProcessing.emit('error', new Error('Something want wrong'));
        const finalString = chunk.toString().replaceAll(/ipsum/gi, 'Romi');
        callback(null, finalString)
        console.log('chunk', chunk.toString());
    }
});

module.exports = replaceWordProcessing;