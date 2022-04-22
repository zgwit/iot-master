const pack = require("./package.json")
const fs = require("fs");
const {execSync} = require("child_process");
const version = pack.version
const time = new Date().toLocaleString()

const ts = `
export var VERSION = "${version}"
export var BUILD_TIME = "${time}"
`
fs.writeFileSync("src/version.ts", ts)

execSync("ng build")
