# Trade and Quote serialization benchmark

This is a comparison of using Protobuf, VTProtobuf and custom Raw format of
marshaling and unmarshaling trade and quote data.

## Marshal

```text
BenchmarkKarmemTrade_Marshal-12         10188492               106.1 ns/op              64.00 B/obj           64 B/op          1 allocs/op
BenchmarkKarmemQuote_Marshal-12         10493108               104.5 ns/op              64.00 B/obj           64 B/op          1 allocs/op
BenchmarkProtoTrade_Marshal-12           4618500               245.1 ns/op              52.54 B/obj           64 B/op          1 allocs/op
BenchmarkProtoQuote_Marshal-12           3773437               315.2 ns/op              66.00 B/obj           80 B/op          1 allocs/op
BenchmarkProtoTrade_MarshalVT-12         7878373               145.2 ns/op              62.73 B/obj           64 B/op          1 allocs/op
BenchmarkProtoQuote_MarshalVT-12         6960794               164.7 ns/op              76.00 B/obj           80 B/op          1 allocs/op
BenchmarkRawTrade_Marshal-12            11121348               100.4 ns/op              53.00 B/obj           64 B/op          1 allocs/op
BenchmarkRawQuote_Marshal-12            10978698               100.1 ns/op              56.00 B/obj           64 B/op          1 allocs/op
```

## Unmarshal

```text
BenchmarkKarmemTrade_Unmarshal-12       88295451                12.87 ns/op            0 B/op          0 allocs/op
BenchmarkKarmemQuote_Unmarshal-12       89969697                11.88 ns/op            0 B/op          0 allocs/op
BenchmarkProtoTrade_Unmarshal-12         5372725               212.7 ns/op            24 B/op          2 allocs/op
BenchmarkProtoQuote_Unmarshal-12         4655619               244.6 ns/op            24 B/op          2 allocs/op
BenchmarkProtoTrade_UnmarshalVT-12      12671863                88.38 ns/op           16 B/op          1 allocs/op
BenchmarkProtoQuote_UnmarshalVT-12      12459482                88.84 ns/op           16 B/op          1 allocs/op
BenchmarkRawTrade_Unmarshal-12          236329374                5.137 ns/op           0 B/op          0 allocs/op
BenchmarkRawQuote_Unmarshal-12          197576809                5.987 ns/op           0 B/op          0 allocs/op
```

```sh
make generate
make bench
```
