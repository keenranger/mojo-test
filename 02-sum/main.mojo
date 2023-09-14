import time
from utils.index import Index
from random import rand

fn main():
    var total : Float32 = 0
    let ready = time.now()
    let randoms = rand[DType.float32](1000000000)
    let start = time.now()
    for i in range(100000000):
        if randoms[i] > 0.5:
            total += randoms[i]
    let end = time.now()
    let elapsed = end - start
    print(total)
    print(elapsed/1e6) 