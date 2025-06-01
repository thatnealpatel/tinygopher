# tinygopher

Currently, a naive implementation exists for flashing
a ESP8266EX with some text and a basic collection of Go
gophers.

This repository is updated on a best-effort basis.

## Flashing

Since I use a developement version of Go pretty much exclusively,
I chose to build LLVM and TinyGo from scratch to elide problems
with the Go and TinyGo toolchains that I did not want to diagnose.

Details: https://tinygo.org/docs/guides/build/manual-llvm/

```sh
% git clone https://github.com/tinygo-org/tinygo.git
% cd tinygo
% git checkout dev
% git submodule update --init --recursive
% make llvm-source llvm-build   # build LLVM
% make                          # build TinyGo
```

In addition, some of the bindings per-chip are required
for flashing to work:

```sh
# todo(nealpatel) update this to include how I built this binary
% export PATH="$PATH:$HOME/esp/xtensa-esp32-elf/bin"
```

After set up, one-off flashes:

```sh
# set [port] and [target] accordingly to platform
% tinygo flash -opt=2 -port /dev/ttyUSB0 -target=nodemcu -monitor github.com/thatnealpatel/tinygopher
```

## License

See [`LICENSE`](LICENSE) for details.

## Disclaimer

This project is not an official Google project. It is not supported by
Google and Google specifically disclaims all warranties as to its quality,
merchantability, or fitness for a particular purpose.
