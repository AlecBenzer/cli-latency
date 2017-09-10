Can you tell the difference between 10ms and 100ms on the command line?

```shell
$ go get github.com/AlecBenzer/cli-latency
$ cli-latency
This is 10ms (Press Enter):
Start...finish!
This is 100ms (Press Enter):
Start...finish!


Press Enter when ready...

Start...finish!

How long was that?
A) 10ms
B) 100ms
B

Press Enter when ready...

# ...

Press Enter when ready...

Start...finish!

How long was that?
A) 10ms
B) 100ms
B

Press Enter when ready...
^Ccorrect: 12
wrong: 0

100% accuracy
```

Inspired by: https://news.ycombinator.com/item?id=15059795
