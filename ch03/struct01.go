/*
Struct: Data Structure in Go

What is stuct ?
  Struct is a data structure which encapsulates data fields of data types. Being said, this data fields can be of different data types as well.
  We have seen array, slice and map separately. Sometimes data can be in format of string , arrays or map. In this case it is hard to reference this
  data separately. To solve this problem struct data structure can be used.
  Struct is kind of data structure which contains a data of type int , data of string, data of float, data of array , data of slice or data of map. All these
  data can be reference by using a single variable.

  Valuse stored in struct can be easily accessed or changed.

Syntax:

  type struct_name struct {
    data_field_1 data_type
    data_field_2 data_type
    ...
    ...
    ...
  }

 Example:
    An airline system should display a invoice showing name of passenger, data, age, flight number, date of travel.
    Since this example contains an single entity of airline and all data is related to one passenger so we can create a struct for this example.
    One struct is referenced as a single variable for an airline.
*/
package main

import (
  "fmt"
)

type airline struct {
  passengerName string
  age int
  flightNumber string
  travelDate string
  travelTime string
  airlineName string
}
// This is a template where nothing is initialized. This template can be used multiple time which sevaral airlines.
func main(){
  // consider airline (jetairways) wants to use this struct to inject its passenger data into the system. so we can create a variable specific to jetairways.
  jetairways := airline {             // Object for jet airways
                passengerName: "Alex Walker",
                age: 28,
                flightNumber: "JA1990",
                travelDate: "Aug 21, 2019",
                travelTime: "19.08",
                airlineName: "Jet Airways",
  }
  unitedAirline := airline {           // Object for united airline.
    passengerName: "Jim Bob",
    age: 34,
    flightNumber: "UA689",
    travelDate: "June 01, 2019",
    travelTime: "00.23",
    airlineName: "United Airline",
  }
  // Displaying jetairways customer details. Now all data has been stored into a variable called jetairways. All data has been stored into a slice (underneath array)
  // And a block of that slice is represented by jetairways variable.
  // Individual data can be retrived by using . notation.
  fmt.Println(jetairways)
  // output:
  //{Alex Walker 28 JA1990 Aug 21 2019 19.08 Jet Airways}

  fmt.Println(unitedAirline)
  // output:
  //{Jim Bob 34 UA689 June 01, 2019 00.23 United Airline}

  // printing individual data
  fmt.Println("Jet Airways Data")
  fmt.Println(jetairways.passengerName)
  fmt.Println(jetairways.age)
  fmt.Println(jetairways.flightNumber)
  fmt.Println(jetairways.travelDate)
  fmt.Println(jetairways.travelTime)
  fmt.Println(jetairways.airlineName)
  fmt.Println("United Airline Data")
  fmt.Println(jetairways.passengerName)
  fmt.Println(jetairways.age)
  fmt.Println(jetairways.flightNumber)
  fmt.Println(jetairways.travelDate)
  fmt.Println(jetairways.travelTime)
  fmt.Println(jetairways.airlineName)
}
/*
In this way struct template can be used to initialize any object for any airline.
*/
