# Trade and Quote serialization benchmark

This is a comparison of using Protobuf, VTProtobuf and custom Raw format of
marshaling and unmarshaling trade and quote data.

```text
benchstat result/proto.out result/protovt.out result/karmem.out result/raw.out result/rawunsafe.out
name \ time/op     result/proto.out  result/protovt.out  result/karmem.out  result/raw.out  result/rawunsafe.out
TradeMarshal-12          199ns ± 7%           77ns ± 1%          38ns ±27%       26ns ± 1%              5ns ±13%
TradeUnmarshal-12        250ns ±16%           91ns ± 0%          14ns ± 5%        4ns ± 2%              7ns ± 0%
QuoteMarshal-12          236ns ±11%           78ns ± 1%          38ns ± 4%       27ns ± 1%              5ns ± 0%
QuoteUnmarshal-12        222ns ± 1%           93ns ± 1%          16ns ± 1%        5ns ± 1%              6ns ± 0%

name \ B/obj       result/proto.out  result/protovt.out  result/karmem.out  result/raw.out  result/rawunsafe.out
TradeMarshal-12           63.7 ± 0%           63.9 ± 0%          64.0 ± 0%       54.0 ± 0%             54.0 ± 0%
TradeUnmarshal-12
QuoteMarshal-12           65.6 ± 0%           65.8 ± 0%          64.0 ± 0%       58.0 ± 0%             58.0 ± 0%
QuoteUnmarshal-12

name \ alloc/op    result/proto.out  result/protovt.out  result/karmem.out  result/raw.out  result/rawunsafe.out
TradeMarshal-12          64.0B ± 0%          64.0B ± 0%         64.0B ± 0%      64.0B ± 0%             0.0B     
TradeUnmarshal-12        24.0B ± 0%          16.0B ± 0%          0.0B            0.0B                  0.0B     
QuoteMarshal-12          79.0B ± 0%          79.0B ± 0%         64.0B ± 0%      64.0B ± 0%             0.0B     
QuoteUnmarshal-12        16.0B ± 0%          16.0B ± 0%          0.0B            0.0B                  0.0B     

name \ allocs/op   result/proto.out  result/protovt.out  result/karmem.out  result/raw.out  result/rawunsafe.out
TradeMarshal-12           1.00 ± 0%           1.00 ± 0%          1.00 ± 0%       1.00 ± 0%             0.00     
TradeUnmarshal-12         2.00 ± 0%           1.00 ± 0%          0.00            0.00                  0.00     
QuoteMarshal-12           1.00 ± 0%           1.00 ± 0%          1.00 ± 0%       1.00 ± 0%             0.00     
QuoteUnmarshal-12         1.00 ± 0%           1.00 ± 0%          0.00            0.00                  0.00   
```