import math


def quadratic_formula(a, b, c):
    # Calculate the discriminant
    discriminant = b ** 2 - 4 * a * c

    # Check if the discriminant is less than 0
    if discriminant < 0:
        return "This equation has no real solutions"

    # Calculate the two roots of the quadratic equation
    root1 = (-b + math.sqrt(discriminant)) / (2 * a)
    root2 = (-b - math.sqrt(discriminant)) / (2 * a)

    return (root1, root2)


# Example
quadratic_formula(1, 2, -3)
# Output: (1.0, -3.0)
