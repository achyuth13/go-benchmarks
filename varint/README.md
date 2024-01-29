VarInt

Databases use VarInt to store and transmit int64 type data effeciently. It helps the application become more 
Storage Effecient, Network Effecient but there's a slight cost of CPU because of increase in computation.

go test -v varint/main_test.go   

Output
=== RUN   TestMain
    main_test.go:9: [164 2]
    main_test.go:10: [140 10]
    main_test.go:11: [132 96]
--- PASS: TestMain (0.00s)
PASS
ok      command-line-arguments  0.005s
