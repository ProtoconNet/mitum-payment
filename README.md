### mitum-payment

*mitum-payment* is a payment contract model based on the second version of mitum(aka [mitum2](https://github.com/ProtoconNet/mitum2)).

#### Installation

Before you build `mitum-payment`, make sure to run mongodb for digest api.

```sh
$ git clone https://github.com/ProtoconNet/mitum-payment

$ cd mitum-payment

$ go build -o ./mitum-payment
```

#### Run

```sh
$ ./mitum-payment init --design=<config file> <genesis file>

$ ./mitum-payment run <config file> --dev.allow-consensus
```

[standalong.yml](standalone.yml) is a sample of `config file`.

[genesis-design.yml](genesis-design.yml) is a sample of `genesis design file`.
