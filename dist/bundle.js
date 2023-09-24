(function() {
var modules = {};
function require(modulePath) {
  var module = modules[modulePath];
  if (!module) {
    throw new Error('Cannot find module ' + modulePath);
  }
  if (!module.exports) {
    module.exports = {};
    module(module.exports, module, require);
  }
  return module.exports;
}
modules['src/main.js'] = (function() {
// src/main.js
const module1 = require('./module1');
const module2 = require('./module2');

console.log(module1);
console.log(module2);

return exports;
})();
modules['src/module1.js'] = (function() {
module.exports = 'This is module 1';

return exports;
})();
modules['src/module2.js'] = (function() {
module.exports = 'This is module 2';

return exports;
})();
require('main.js');
})();