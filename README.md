# brainfuck cli interpreter


## insatll
``` bash
go install github.com/mreza0100/brainfuck-cli@latest
```

# usage

### interactive mode
```bash
brainfuck-cli i
```
<!-- use image -->
<img src="repl.png" width="500" height="400">


### read from args
```bash
brainfuck-cli r ">+++++++++[<++++++++>-]<.>+++++++[<++++>-]<+.+++++++..+++.[-]>++++++++[<++++>-] <.>+++++++++++[<++++++++>-]<-.--------.+++.------.--------.[-]>++++++++[<++++>- ]<+.[-]++++++++++."
```
### read from file
```bash
brainfuck-cli f ./test.bf
```
