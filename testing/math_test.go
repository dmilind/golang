/*
Testing Go Program/Package
Writing test program is an important aspect of any program. Test programs make sure that the go program is behaving exactly the way developer want.
It avoids unexpected failure of the program. Software test are driven by the user stories where each functionality can be verified.
Types of tests:
Unit test:
  Each module or function can be tested individually. Some raw data will be passed to the function and verification can be done once result is returned. Any change
  over the time period on same code. Regresion is a bug which occurs after the change is made on the existing code. Unit test suits can be fixed by regreesion
  testin.

Integration Testing:
  Once unit test is performed on the single function, integration testing can be perfored later. In integration testing, all components of the program are
  verified. It checks that all the components/functions are working appropriately.

Functional Testing:
  Functions testing always takes place on the user side, where user verfies that application is in intact and behaving as per expectations.

Testing in Go:
  Go provides a package testing which helps in writing the test suits for the written go program. There are some rules for writting test suits.
  1. test suits are saved as program_test.go
  2. test suits are placed where assosciated go code is placed.
  3. Functions in the test suites are begine with Test keyword

Command to run test suits:
  go test


*/

package math

import (
  "testing"
)

func TestAverage(t *testing.T) {
  got := Average(3.4,7.9)
  want := 5.65
  if got != want {
    t.Fatalf("unexpected average value %f from values %f", want,got)
  }
}

func TestSum(t *testing.T) {
  got := Sum(3.4,7.9)
  want := 11.3
  if got != want {
    t.Fatalf("unexpected sum value %f from values %f",want,got)
  }
}

func TestSubtract(t *testing.T) {
  got := Subtract(3.4,7.9)
  want := (3.4-7.9)
  if got != want {
    t.Fatalf("Unexpected subtact value of %f from %f", want, got)
  }
}

/*
go test result :
alculating Average:
Calculating Sum:
Calculating Subtract:
PASS
ok      practice/golang/testing 0.006s

go test -cover result

Calculating Average:
Calculating Sum:
Calculating Subtract:
PASS
coverage: 100.0% of statements
ok      practice/golang/testing 0.006s

*/
