#!/usr/bin/env python

"""Bubblesort.py: sorts a list of numbers"""

def bubblesort(sorting_list):
    """Performs bubble sort on a list and returns a list sorted numerically"""
    in_order = checklistorder(sorting_list)
    while not in_order:
        for idx in range(len(sorting_list)-1):
            if sorting_list[i] > sorting_list[i+1]:
                temp = sorting_list[i]
                sorting_list[i] = sorting_list[i+1]
                sorting_list[i+1] = temp
        in_order = checklistorder(sorting_list)
    return sorting_list
    

def checklistorder(check_list):
    """Checks whether a list is in numerical order. Returns boolean."""
    in_order = True
    for idx in range(len(check_list)-1):
        if check_list[i] > check_list[i+1]:
            in_order = False
            return in_order
    return in_order

if __name__ == '__main__':
    print('bubblesort has been imported!')

