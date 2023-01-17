"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.Log = void 0;
var Log = /** @class */ (function () {
    function Log() {
    }
    Log.success = function (msg) {
        console.log("%c ".concat(msg), 'color: green');
    };
    Log.danger = function (msg) {
        console.log("%c ".concat(msg), 'color: red');
    };
    Log.info = function (msg) {
        console.log("%c ".concat(msg), 'color: black; background: yellow');
    };
    return Log;
}());
exports.Log = Log;
