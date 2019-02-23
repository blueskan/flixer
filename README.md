
# flixer

  Automatically creates Web UI for CLI Prompts

[![Build Status](https://travis-ci.org/blueskan/flixer.svg?branch=master)](https://travis-ci.org/blueskan/flixer)

This project strongly inspired from Google Cloud Engine - Command Line Interface.
Purpose of the this project provide bridge between CLI and Web UI therefore ensure more user friendly CLI Prompts.

For example you build node application and ask some questions to user in CLI you can pipe your process to flixer and flixer automatically open browser (completely cross platform) and http server for you (actually serve your predefined html template) and after send post request to predefined post url you can send to user input directly stdout and redirect your main process or write file and read after (Currently support this 2 strategies). Also after this stage process die with gracefully.

**Node.JS Example**

```js
const { spawn } = require('child_process');
// Remember run without absolute path, flixer executable should be placed in your operating system $PATH
const flixerDoSomeAwesomeThink = spawn('flixer', ['run', 'stdout', '--template=./example/flixer.html']);

flixerDoSomeAwesomeThink.stdout.on('data', (data) => {
  // directly read stdout
  console.log(`stdout: ${data}`);
});
```

In the other programming languages this process same more or less..

If you want to using file mode (maybe you wanna do some scp trick in unix variants) just pass `--output-file` option to
flixer and you see created file with post content after request..

We are currently support just 2 Output Modes:  `stdout | file`
I will plan adds more output modes in days to come.

**All available options below:**

![cli flags etc](https://raw.githubusercontent.com/blueskan/flixer/master/example/cli.png)

**Quick Demo:**

![demo](https://raw.githubusercontent.com/blueskan/flixer/master/example/demo.gif)

**Installion:**
If you define your $GOPATH and $GOBIN variables and $GOBIN in your $PATH then very simle;
`go get -u github.com/blueskan/flixer`
and you can run anywhere with typing `flixer run ...`

Otherwise you can compile code yourself and build executable..
Example flixer template can be found in `example/` directory.

**TODOS**
- We need more testing.
- Different input support and more output support strategy
- Maybe automatic template generation?!

If you want to contribute with any idea or code feel free to send me an e-mail via me@batikansenemoglu.com or pull request.
