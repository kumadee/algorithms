#!/usr/bin/env python3
def fibonacci():
    result = []
    while True:
        if len(result) < 2:
            result.append(1)
        else:
            result = [result[-1], result[-1] + result[-2]]
        if result[-1] % 17 == 0:
            return result[-1]
        if result[-1] > 10000:
            return

if __name__ == "__main__":
    res = fibonacci()
    if res != None:
        print(res)
