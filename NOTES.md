- HTTP requests need to use PDK supplied methods instead of native methods
    - True for Go and JS so far
    - Hosts must be explicitly allowed. "*" also works. This is because of security concerns.
    - WASI must be enabled (also necessary for FS access?) (Maybe make a PR/Issue for pointing that out inside the Http section of READMEs)

- TS: CLI scaffolded boilerplate would benefit from `import "@extism/js-pdk"` being included at the top of the index.ts file (I should just make a PR)
    - SDK doesn't have strong typing for the first arg to `createPlugin`. In editor, we were not yelled at about a missing key in our manifest object

- VSCode
    - Activating extension 'eamodio.gitlens' failed: WebAssembly.compileStreaming is not a function. (caused by trying to load remote wasm files)

- Documenting things: Each surface at GitKraken might not have all the same events/hooks. eg: GKD and Gitlens both do git actions, GKDev and CLI do not. but launchpad and other features are crosscutting

- pivot to trying local filesystem writes/logging (eventually)

- extism cli scaffolding for plugins is nice

- OutputString will only work once per function. last one wins. If you want to output multiple lines, you will need to build the string as the function executes and send it at the end

- Filesystem path access is a bit unclear, an example would be really helpful
    - seems file system access in go needs a bit of extra work:
        https://github.com/extism/go-pdk/blob/main/example/reactor/tiny_main.go

- Host functions seem to be unable to be used when "runInWorker" is set to true in the JS SDK

//////////////////

Next steps:
    - Figure out filesystem stuff.
    - Start creating Slides
    - Discuss actual plugin functionality with Eric



//////////////////

Our Plugin API Surface Area

## Events (from Host to Extension)

- launchpad
    - loaded
    - pinned
    - snoozed
    - etc

## Functions (callable Host functions from within Extension)

- notification

## Hooks (from Extension to Host)

- init
- 






//////////////////

Project:

I got accepted for a talk to SquiggleConf which is a conference about dev tools in Boston in October. ( https://2024.squiggleconf.com/ ). The talk is about Extism using WASM as a compatibility layer for plugin systems. I am exploring adding plugin to our GitKraken CLI tool.

Today:

We got the http GET request working. Just needed to explicitly allow the host/domain. We only did that with the CLI testing tool. Not sure how to do that in the Host application though. I also need to figure out why the POST request is failing.
