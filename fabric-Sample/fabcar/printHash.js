'use strict';

const md5File = require('md5-file')

var commandLineArgs = process.argv.slice(2);

const hash = md5File.sync(commandLineArgs[0])
console.log(hash)