Divided the 100000000 to be computed over 10 threads, but each thread checks if a number is prime, if it's not it continues otherwise returns

The division is not fair because each thread is checking only one thing and incrementing

#Benchmark
Thread 8 completed in 7.561313125s
Thread 2 completed in 7.561211833s
Thread 7 completed in 7.561255167s
Thread 6 completed in 7.561154167s
Thread 3 completed in 7.561336042s
Thread 1 completed in 7.561283125s
Thread 5 completed in 7.561405958s
Thread 0 completed in 7.550763958s
Thread 4 completed in 7.561031583s
Checking till  100000000 found,  5761455  prime numbers and took  7.561576083s