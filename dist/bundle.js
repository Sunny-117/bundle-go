// Module: src/main.js
// src/main.js
const module1 = require('./module1');
const module2 = require('./module2');

console.log(module1);
console.log(module2);


// Module: src/module1.js
module.exports = 'This is module 1';


// Module: src/module2.js
module.exports = 'This is module 2';


