const fs = require('fs');
const path = require('path');
const output = '<!doctype html><html><body><main id="app">培训报名审核系统</main></body></html>\n';
fs.writeFileSync(path.join(__dirname, 'index.html'), output);
