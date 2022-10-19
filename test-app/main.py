#!/bin/env python
import time



def a():
    print("function a")
    b()

def b():
    print("function b")
    c()
def c():
    print("function c")
    time.sleep(6.4)
    d()


def d():
    print("function d")


while True:
    a()