# Go Sample Examples

<p align="center">
    <img src="images/golang.png" alt="Main Information" width="600" height="400">
</p>

### 📖 Information

<ul style="list-style-type:disc">
  <li>This project showcases sample Go applications demonstrating various functionalities, including data processing, web server management, and concurrent programming techniques.</li>
  <li>It is designed to help developers understand Go's syntax, features, and best practices through practical examples.</li>
  <li>The repository includes:
    <ul>
      <li>Examples of basic Go data structures, functions, and interfaces.</li>
      <li>Web server setup using Go's standard library for HTTP handling.</li>
      <li>Concurrency examples using goroutines, channels, and the Go scheduler.</li>
      <li>File handling operations like reading from and writing to files.</li>
      <li>Working with arrays, slices, and maps to demonstrate Go's powerful collection types.</li>
      <li>JSON encoding and decoding examples, including working with custom data structures.</li>
      <li>Error handling following Go idioms to demonstrate best practices for robust programs.</li>
      <li>Usage of Go's testing framework to implement unit and integration tests.</li>
      <li>Practical examples of using Go's `net/http` package for building RESTful web services.</li>
      <li>Basic data parsing, string manipulation, arithmetic operations, and JSON handling in Go.</li>
    </ul>
  </li>
</ul>


### Explore Go Sample Examples

Examples Summary

<table style="width:100%">
  <tr>
      <th>Example</th>
      <th>Description</th>
      <th>Link</th>
  </tr>
  <tr>
      <td>Hello World</td>
      <td>Basic example of printing "Hello, World!" to the console.</td>
      <td><a href="/001_hello_world">001_hello_world</a></td>
  </tr>
  <tr>
      <td>Variables & Types</td>
      <td>Demonstrates variable declaration, initialization, and types.</td>
      <td><a href="/002_variables">002_variables</a></td>
  </tr>
  <tr>
      <td>Control Flow</td>
      <td>Shows usage of for loops, while loops, if-else statements, and switch statements.</td>
      <td><a href="/003_for_while_if_else_switch">003_for_while_if_else_switch</a></td>
  </tr>
  <tr>
      <td>Array, Slice, Range, Map</td>
      <td>Demonstrates array and slice operations, range iteration, and map manipulation.</td>
      <td><a href="/004_array_slice_range_map">004_array_slice_range_map</a></td>
  </tr>
  <tr>
      <td>Functions</td>
      <td>Illustrates basic function usage, variadic functions, closures, and recursion.</td>
      <td><a href="/005_function">005_function</a></td>
  </tr>
  <tr>
      <td>Pointers</td>
      <td>Shows how to use pointers to modify values, structs, slices, and maps. Includes pointer-to-pointer operations.</td>
      <td><a href="/006_pointer">006_pointer</a></td>
  </tr>
  <tr>
      <td>Strings & Runes</td>
      <td>Explains how to manipulate strings and runes, including encoding, slicing, and comparing.</td>
      <td><a href="/007_strings_runes">007_strings_runes</a></td>
  </tr>
  <tr>
      <td>Structs</td>
      <td>Demonstrates struct declaration, methods, embedded structs, JSON tags, and more.</td>
      <td><a href="/008_structs">008_structs</a></td>
  </tr>
  <tr>
      <td>Interfaces</td>
      <td>Covers defining and implementing interfaces, type assertions, type switches, and using methods with multiple interfaces.</td>
      <td><a href="/009_interfaces">009_interfaces</a></td>
  </tr>
  <tr>
      <td>Error Handling</td>
      <td>Shows basic error creation, custom error types, error wrapping, unwrapping, and checking for specific errors.</td>
      <td><a href="/010_error">010_error</a></td>
  </tr>
  <tr>
      <td>Goroutines & Channels</td>
      <td>Demonstrates using Goroutines for concurrency, WaitGroups for synchronization, channels for communication, and mutex for safe access to shared data.</td>
      <td><a href="/011_goroutine_channel">011_goroutine_channel</a></td>
  </tr>
  <tr>
      <td>Buffering</td>
      <td>Shows how to use buffered channels, select statements, and bytes.Buffer, along with a custom buffer implementation.</td>
      <td><a href="/012_buffering">012_buffering</a></td>
  </tr>
  <tr>
      <td>Channel Synchronization</td>
      <td>Demonstrates different ways of using channels for synchronization, including WaitGroups, one-way channels, and fan-out/fan-in patterns.</td>
      <td><a href="/013_channel_synchronization">013_channel_synchronization</a></td>
  </tr>
  <tr>
      <td>Channel Directions</td>
      <td>Shows how to use directional channels (send-only and receive-only) in Go, with examples including ping-pong, pipelining, and buffered channels.</td>
      <td><a href="/014_channel_directions">014_channel_directions</a></td>
  </tr>
  <tr>
      <td>Channel Select</td>
      <td>Demonstrates the use of select statement with multiple channels to handle concurrent events.</td>
      <td><a href="/015_channel_select">015_channel_select</a></td>
  </tr>
  <tr>
      <td>Timeouts</td>
      <td>Illustrates how to use timeouts with Goroutines and channels to control the execution flow.</td>
      <td><a href="/016_timeouts">016_timeouts</a></td>
  </tr>
  <tr>
      <td>Channel Non-Blocking</td>
      <td>Demonstrates non-blocking operations on channels using the select statement.</td>
      <td><a href="/017_channel_non_blocking">017_channel_non_blocking</a></td>
  </tr>
  <tr>
    <td rowspan="5">Channel Closing Examples</td>
    <td>Basic Channel Closing</td>
    <td><a href="/018_closing_channel/01_basic_channel_closing">01_basic_channel_closing</a></td>
  </tr>
  <tr>
    <td>Detecting Closed Channel</td>
    <td><a href="/018_closing_channel/02_detecting_closed_channel">02_detecting_closed_channel</a></td>
  </tr>
  <tr>
    <td>Closing Channels in Multiple Goroutines</td>
    <td><a href="/018_closing_channel/03_closing_channels_in_multiple_goroutines">03_closing_channels_in_multiple_goroutines</a></td>
  </tr>
  <tr>
    <td>Sending a Signal with a Closed Channel</td>
    <td><a href="/018_closing_channel/04_sending_a_signal_with_a_closed_channel">04_sending_a_signal_with_a_closed_channel</a></td>
  </tr>
  <tr>
    <td>Panic When Closing an Already Closed Channel</td>
    <td><a href="/018_closing_channel/05_panic_when_closing_an_already_closed_channel">05_panic_when_closing_an_already_closed_channel</a></td>
  </tr>
</table>


### Technologies

---
- Go 1.23.0


### Prerequisites

Ensure Go is installed on your system
```
    https://golang.org/dl/
```


### Contributors

- [Sercan Noyan Germiyanoğlu](https://github.com/Rapter1990)


