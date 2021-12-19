# Day 3: Binary Diagnostic

[https://adventofcode.com/2021/day/3](https://adventofcode.com/2021/day/3)

## Description

### Part One

The submarine has been making some <span title="Turns out oceans are heavy.">odd creaking noises</span>, so you ask it to produce a diagnostic report just in case.

The diagnostic report (your puzzle input) consists of a list of binary numbers which, when decoded properly, can tell you many useful things about the conditions of the submarine. The first parameter to check is the _power consumption_.

You need to use the binary numbers in the diagnostic report to generate two new binary numbers (called the _gamma rate_ and the _epsilon rate_). The power consumption can then be found by multiplying the gamma rate by the epsilon rate.

Each bit in the gamma rate can be determined by finding the _most common bit in the corresponding position_ of all numbers in the diagnostic report. For example, given the following diagnostic report:

    00100
    11110
    10110
    10111
    10101
    01111
    00111
    11100
    10000
    11001
    00010
    01010
    

Considering only the first bit of each number, there are five `0` bits and seven `1` bits. Since the most common bit is `1`, the first bit of the gamma rate is `1`.

The most common second bit of the numbers in the diagnostic report is `0`, so the second bit of the gamma rate is `0`.

The most common value of the third, fourth, and fifth bits are `1`, `1`, and `0`, respectively, and so the final three bits of the gamma rate are `110`.

So, the gamma rate is the binary number `10110`, or _`22`_ in decimal.

The epsilon rate is calculated in a similar way; rather than use the most common bit, the least common bit from each position is used. So, the epsilon rate is `01001`, or _`9`_ in decimal. Multiplying the gamma rate (`22`) by the epsilon rate (`9`) produces the power consumption, _`198`_.

Use the binary numbers in your diagnostic report to calculate the gamma rate and epsilon rate, then multiply them together. _What is the power consumption of the submarine?_ (Be sure to represent your answer in decimal, not binary.)
