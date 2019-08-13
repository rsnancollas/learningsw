#!/usr/bin/env python

import random
import timeit

"""Bubblesort.py: sorts a list of numbers"""

def bubblesort_check(sorting_list):
    """Performs bubble sort on a list, checks list, and returns a list sorted numerically"""
    in_order = checklistorder(sorting_list)
    while not in_order:
        for i in range(len(sorting_list)-1):
            if sorting_list[i] > sorting_list[i+1]:
                temp = sorting_list[i]
                sorting_list[i] = sorting_list[i+1]
                sorting_list[i+1] = temp
        in_order = checklistorder(sorting_list)
    return sorting_list
    

def checklistorder(check_list):
    """Checks whether a list is in numerical order. Returns boolean."""
    in_order = True
    for i in range(len(check_list)-1):
        if check_list[i] > check_list[i+1]:
            in_order = False
            return in_order
    return in_order

def bubblesort(sorting_list):
    """Performs bubble sort on a list and returns a list sorted numerically"""
    sorted = False
    while not sorted:
        sorted = True
        for i in range(len(sorting_list)-1):
            if sorting_list[i] > sorting_list[i+1]:
                temp = sorting_list[i]
                sorting_list[i] = sorting_list[i+1]
                sorting_list[i+1] = temp
                sorted = False
    return sorting_list

def sort_test():
    """Compare runtime of bubblesort programs"""
    check_sort_time = 0
    sort_time = 0
    for i in range(10000):
        test_list = random.sample(range(100), 100)
        test_list_copy = test_list.copy()
        check_sort_time += timeit.timeit(lambda: bubblesort_check(test_list_copy), number=1)
        sort_time += timeit.timeit(lambda: bubblesort(test_list), number=1)
    print("Bubble sort: %s" %sort_time)
    print("Bubble sort with check: %s" %check_sort_time)


if __name__ == '__main__':
    print('bubblesort has been imported!')

