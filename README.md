# Go Doomsday & Leap Year Checker

A lightweight CLI tool built with Golang to handle calendar-related logic, specifically determining leap years and calculating the day of the week for any date in history using the Doomsday Algorithm.

## Key Features
The project currently supports two core functionalities:

1. Leap Year Validator:

- Validates any user-provided year.
- Uses the Gregorian calendar logic to determine if the year is a leap year.

2. Weekday Calculator (Day Finder):

- Implements John Conway's Doomsday Algorithm.
- Accepts any date (Day/Month/Year) from the past or future.
- Returns the exact day of the week (Monday, Tuesday, etc.).

## The Mathematics Behind
This tool leverages precise astronomical and mathematical rules:

- Leap Year Logic: A year is a leap year if it is divisible by 4, except for century years, which must be divisible by 400.

``` year (mod 4) = 0 and (year (mod 100) != 0 or year (mod 400) = 0) ```

- Century Anchors Logic: Used as a reference point for every 100-year block.

```Anchor = (5 x (century (mod 4) + 2) % 7 ```

The Doomsday Rule: Based on the fact that certain dates (e.g., 4/4, 6/6, 8/8) always fall on the same day of the week within a given year.