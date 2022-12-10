"""
Author: Logan Maupin
Date: 11/21/2022

The purpose of this script is to ask the user for inputs of what the a, b, and c values are of the function.
Then given those values, calculate what your function's intercepts would be given that information.
"""

import math


def setup_variables():
    a_b_and_c_values_are = input("""Please enter the A, B, C values of your function, separated by a comma.\n
So if your function is x^2 + 8x + 15, the A value is 3, the B value is 4, and the C value is 5.\n
So your input should look like '1, 8, 15'\n""")

    list_of_our_values = a_b_and_c_values_are.strip().split(",")

    if len(list_of_our_values) == 3:

        a = float(list_of_our_values[0])
        b = float(list_of_our_values[1])
        c = float(list_of_our_values[2])

    elif len(list_of_our_values) == 2:
        a = float(list_of_our_values[0])
        b = float(list_of_our_values[1])
        c = 0.0

    elif len(list_of_our_values) == 1:

        a = float(list_of_our_values[0])
        b = 0.0
        c = 0.0

    return a, b, c


def quad_formula(a: float, b: float, c: float):
    discriminate = (math.sqrt((b**2) - (4*a*c)))

    if discriminate < 0:
        print("Domain error, your discriminate value can't be a non-real number")
        return 0.0, 0.0

    else:

        x_1 = -b - discriminate / (2*a)
        x_2 = -b + discriminate / (2*a)

        return x_1, x_2


def main():
    a, b, c = setup_variables()
    first_x, second_x = quad_formula(a, b, c)

    print(f"\nYour function's intercept values are: \n{first_x} and {second_x}")


if __name__ == '__main__':
    main()
