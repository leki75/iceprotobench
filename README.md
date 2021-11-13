# Trade and Quote serialization benchmark

This is a comparison of using Protobuf, VTProtobuf and custom Raw format of
marshaling and unmarshaling trade and quote data.

## Marshal

```text
BenchmarkProtoTrade_Marshal-12           4719379               245.3 ns/op              52.55 B/obj           64 B/op          1 allocs/op
BenchmarkProtoTrade_MarshalVT-12         7992558               148.1 ns/op              62.74 B/obj           64 B/op          1 allocs/op
BenchmarkRawTrade_Marshal-12            11897222                99.33 ns/op             53.00 B/obj           64 B/op          1 allocs/op

BenchmarkProtoQuote_Marshal-12           3817172               314.2 ns/op              66.00 B/obj           80 B/op          1 allocs/op
BenchmarkProtoQuote_MarshalVT-12         7147515               164.5 ns/op              76.00 B/obj           80 B/op          1 allocs/op
BenchmarkRawQuote_Marshal-12            11824521                98.06 ns/op             56.00 B/obj           64 B/op          1 allocs/op
```

## Unmarshal

```text
BenchmarkProtoTrade_Unmarshal-12         5502446               216.1 ns/op            24 B/op          2 allocs/op
BenchmarkProtoTrade_UnmarshalVT-12      13182322                90.60 ns/op           16 B/op          1 allocs/op
BenchmarkRawTrade_Unmarshal-12          237809220                4.941 ns/op           0 B/op          0 allocs/op

BenchmarkProtoQuote_Unmarshal-12         4867123               243.3 ns/op            24 B/op          2 allocs/op
BenchmarkProtoQuote_UnmarshalVT-12      12689659                92.78 ns/op           16 B/op          1 allocs/op
BenchmarkRawQuote_Unmarshal-12          225513898                5.262 ns/op           0 B/op          0 allocs/op
```

# Running the benchmark suite

```sh
make generate
make bench
```