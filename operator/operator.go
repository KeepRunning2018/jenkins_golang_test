package operator

import (
    "errors"
)

func Division(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("除数不能为0")
    }

    return a / b, nil
}

func Add(a, b float64) (float64, error) {

    return a + b, nil
}

func minus(a, b float64) (float64, error) {

    return a - b, nil
}

func multiply(a, b float64) (float64, error) {

    return a * b, nil
}