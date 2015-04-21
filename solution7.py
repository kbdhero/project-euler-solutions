def my_prime_test(n):
    for num in xrange(2,n):
        if n % num == 0:
            return False
        else:
            return True

def euler7():
    for n in xrange(1,200000):
        if my_prime_test(n):
            yield n

for enum,n in enumerate(euler7()):
    if enum == 10001:
        print enum,n
