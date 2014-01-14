from threading import Thread

i = 0

def adder():
    global i
    for x in range(0, 1000000):     
        i += 1
def decr():
    global i
    for x in range(0, 1000000):
        i-=1


def main():
    adder_thr = Thread(target = adder)
    decr_thr = Thread(target = decr)
    adder_thr.start()
    decr_thr.start()
    for x in range(0, 50):
        print(i)
    adder_thr.join()
    decr.join()
    print("Done: " + str(i))


main()
