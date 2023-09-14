import random
import time

def main():
    total = 0
    randoms = [random.random() for i in range(1000000)]
    start = time.time()
    for i in range(1000000):
        if randoms[i] > 0.5:
            total += randoms[i]
    elapsed = time.time() - start
    print(total)
    print(elapsed*1000) # python time is in seconds, so we multiply by 1000 to get milliseconds

if __name__ == '__main__':
    main()
    