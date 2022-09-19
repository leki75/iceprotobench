# Trade and Quote serialization benchmark

This is a comparison of using Protobuf, VTProtobuf and custom Raw format of
marshaling and unmarshaling trade and quote data.

```text
$ make
benchstat result/proto.out result/protovt.out result/karmem.out result/raw.out
name \ time/op     result/proto.out  result/protovt.out  result/karmem.out  result/raw.out
TradeMarshal-12          258ns ± 6%          146ns ± 3%         106ns ± 2%      100ns ± 2%
TradeUnmarshal-12        216ns ± 2%           87ns ± 2%          12ns ± 1%        4ns ± 1%
QuoteMarshal-12          312ns ± 1%          158ns ± 3%         103ns ± 1%       98ns ± 1%
QuoteUnmarshal-12        205ns ± 2%           85ns ± 2%          12ns ± 1%        5ns ± 2%

name \ B/obj       result/proto.out  result/protovt.out  result/karmem.out  result/raw.out
TradeMarshal-12           62.5 ± 0%           62.7 ± 0%          64.0 ± 0%       53.0 ± 0%
TradeUnmarshal-12
QuoteMarshal-12           72.0 ± 0%           72.0 ± 0%          64.0 ± 0%       56.0 ± 0%
QuoteUnmarshal-12

name \ alloc/op    result/proto.out  result/protovt.out  result/karmem.out  result/raw.out
TradeMarshal-12          64.0B ± 0%          64.0B ± 0%         64.0B ± 0%      64.0B ± 0%
TradeUnmarshal-12        24.0B ± 0%          16.0B ± 0%          0.0B            0.0B     
QuoteMarshal-12          80.0B ± 0%          80.0B ± 0%         64.0B ± 0%      64.0B ± 0%
QuoteUnmarshal-12        16.0B ± 0%          16.0B ± 0%          0.0B            0.0B     

name \ allocs/op   result/proto.out  result/protovt.out  result/karmem.out  result/raw.out
TradeMarshal-12           1.00 ± 0%           1.00 ± 0%          1.00 ± 0%       1.00 ± 0%
TradeUnmarshal-12         2.00 ± 0%           1.00 ± 0%          0.00            0.00     
QuoteMarshal-12           1.00 ± 0%           1.00 ± 0%          1.00 ± 0%       1.00 ± 0%
QuoteUnmarshal-12         1.00 ± 0%           1.00 ± 0%          0.00            0.00 
```
