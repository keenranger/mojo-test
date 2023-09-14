import time

fn main():
    var total = 0
    let start = time.now()
    for i in range(10000000000000000000000000000000000):
        total += i
    let elapsed = time.now() - start
    print(total)
    print(elapsed)
    