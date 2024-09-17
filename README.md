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

<table>
  <tr>
      <th>ID</th>
      <th>Example</th>
      <th>Description</th>
      <th>Link</th>
  </tr>
  <tr>
      <td>1</td>
      <td>Hello World</td>
      <td>Basic example of printing "Hello, World!" to the console.</td>
      <td><a href="/001_hello_world">001_hello_world</a></td>
  </tr>
  <tr>
      <td>2</td>
      <td>Variables & Types</td>
      <td>Demonstrates variable declaration, initialization, and types.</td>
      <td><a href="/002_variables">002_variables</a></td>
  </tr>
  <tr>
      <td>3</td>
      <td>Control Flow</td>
      <td>Shows usage of for loops, while loops, if-else statements, and switch statements.</td>
      <td><a href="/003_for_while_if_else_switch">003_for_while_if_else_switch</a></td>
  </tr>
  <tr>
      <td>4</td>
      <td>Array, Slice, Range, Map</td>
      <td>Demonstrates array and slice operations, range iteration, and map manipulation.</td>
      <td><a href="/004_array_slice_range_map">004_array_slice_range_map</a></td>
  </tr>
  <tr>
      <td>5</td>
      <td>Functions</td>
      <td>Illustrates basic function usage, variadic functions, closures, and recursion.</td>
      <td><a href="/005_function">005_function</a></td>
  </tr>
  <tr>
      <td>6</td>
      <td>Pointers</td>
      <td>Shows how to use pointers to modify values, structs, slices, and maps. Includes pointer-to-pointer operations.</td>
      <td><a href="/006_pointer">006_pointer</a></td>
  </tr>
  <tr>
      <td>7</td>
      <td>Strings & Runes</td>
      <td>Explains how to manipulate strings and runes, including encoding, slicing, and comparing.</td>
      <td><a href="/007_strings_runes">007_strings_runes</a></td>
  </tr>
  <tr>
      <td>8</td>
      <td>Structs</td>
      <td>Demonstrates struct declaration, methods, embedded structs, JSON tags, and more.</td>
      <td><a href="/008_structs">008_structs</a></td>
  </tr>
  <tr>
      <td>9</td>
      <td>Interfaces</td>
      <td>Covers defining and implementing interfaces, type assertions, type switches, and using methods with multiple interfaces.</td>
      <td><a href="/009_interfaces">009_interfaces</a></td>
  </tr>
  <tr>
      <td>10</td>
      <td>Error Handling</td>
      <td>Shows basic error creation, custom error types, error wrapping, unwrapping, and checking for specific errors.</td>
      <td><a href="/010_error">010_error</a></td>
  </tr>
  <tr>
      <td>11</td>
      <td>Goroutines & Channels</td>
      <td>Demonstrates using Goroutines for concurrency, WaitGroups for synchronization, channels for communication, and mutex for safe access to shared data.</td>
      <td><a href="/011_goroutine_channel">011_goroutine_channel</a></td>
  </tr>
  <tr>
      <td>12</td>
      <td>Buffering</td>
      <td>Shows how to use buffered channels, select statements, and bytes.Buffer, along with a custom buffer implementation.</td>
      <td><a href="/012_buffering">012_buffering</a></td>
  </tr>
  <tr>
      <td>13</td>
      <td>Channel Synchronization</td>
      <td>Demonstrates different ways of using channels for synchronization, including WaitGroups, one-way channels, and fan-out/fan-in patterns.</td>
      <td><a href="/013_channel_synchronization">013_channel_synchronization</a></td>
  </tr>
  <tr>
      <td>14</td>
      <td>Channel Directions</td>
      <td>Shows how to use directional channels (send-only and receive-only) in Go, with examples including ping-pong, pipelining, and buffered channels.</td>
      <td><a href="/014_channel_directions">014_channel_directions</a></td>
  </tr>
  <tr>
      <td>15</td>
      <td>Channel Select</td>
      <td>Demonstrates the use of select statement with multiple channels to handle concurrent events.</td>
      <td><a href="/015_channel_select">015_channel_select</a></td>
  </tr>
  <tr>
      <td>16</td>
      <td>Timeouts</td>
      <td>Illustrates how to use timeouts with Goroutines and channels to control the execution flow.</td>
      <td><a href="/016_timeouts">016_timeouts</a></td>
  </tr>
  <tr>
      <td>17</td>
      <td>Channel Non-Blocking</td>
      <td>Demonstrates non-blocking operations on channels using the select statement.</td>
      <td><a href="/017_channel_non_blocking">017_channel_non_blocking</a></td>
  </tr>
  <tr>
      <td rowspan="5">18</td>
      <td>Basic Channel Closing</td>
      <td>Shows how to close a channel and handle its closure properly.</td>
      <td><a href="/018_closing_channel/01_basic_channel_closing">01_basic_channel_closing</a></td>
  </tr>
  <tr>
      <td>Detecting Closed Channel</td>
      <td>Explains how to detect if a channel has been closed and how to handle it.</td>
      <td><a href="/018_closing_channel/02_detecting_closed_channel">02_detecting_closed_channel</a></td>
  </tr>
  <tr>
      <td>Closing Channels in Multiple Goroutines</td>
      <td>Demonstrates how to safely close channels when multiple Goroutines are involved.</td>
      <td><a href="/018_closing_channel/03_closing_channels_in_multiple_goroutines">03_closing_channels_in_multiple_goroutines</a></td>
  </tr>
  <tr>
      <td>Sending a Signal with a Closed Channel</td>
      <td>Shows the behavior of sending signals through a closed channel and how to manage it.</td>
      <td><a href="/018_closing_channel/04_sending_a_signal_with_a_closed_channel">04_sending_a_signal_with_a_closed_channel</a></td>
  </tr>
  <tr>
      <td>Panic When Closing an Already Closed Channel</td>
      <td>Explains what happens when you attempt to close an already closed channel and how to handle it.</td>
      <td><a href="/018_closing_channel/05_panic_when_closing_an_already_closed_channel">05_panic_when_closing_an_already_closed_channel</a></td>
  </tr>
  <tr>
      <td rowspan="3">19</td>
      <td>Basic Example with Range Over Channel</td>
      <td>Shows how to use the range keyword to iterate over values received from a channel.</td>
      <td><a href="/019_range_over_channel/001_basic_example_range_over_channel">001_basic_example_range_over_channel</a></td>
  </tr>
  <tr>
      <td>Multiple Goroutines Sending to a Channel</td>
      <td>Demonstrates how multiple Goroutines can send data to a single channel and how to handle it.</td>
      <td><a href="/019_range_over_channel/002_multiple_goroutines_sending_to_a_channel">002_multiple_goroutines_sending_to_a_channel</a></td>
  </tr>
  <tr>
      <td>Buffered Channels with Range</td>
      <td>Shows how to use buffered channels in conjunction with range to manage concurrent operations.</td>
      <td><a href="/019_range_over_channel/003_buffered_channels_with_range">003_buffered_channels_with_range</a></td>
  </tr>
  <tr>
      <td rowspan="5">20</td>
      <td>Simple Timer</td>
      <td>Demonstrates how to create and use a basic timer in Go.</td>
      <td><a href="/020_timers/001_simple_timer">001_simple_timer</a></td>
  </tr>
  <tr>
      <td>Stop Timer</td>
      <td>Shows how to stop a timer before it triggers.</td>
      <td><a href="/020_timers/002_stop_timer">002_stop_timer</a></td>
  </tr>
  <tr>
      <td>Reset Timer</td>
      <td>Explains how to reset a timer to its initial state.</td>
      <td><a href="/020_timers/003_reset_timer">003_reset_timer</a></td>
  </tr>
  <tr>
      <td>Using time.After</td>
      <td>Demonstrates how to use the time.After function to create timers.</td>
      <td><a href="/020_timers/004_using_time_after">004_using_time_after</a></td>
  </tr>
  <tr>
      <td>Timer with Select</td>
      <td>Shows how to use timers with the select statement for time-based control flow.</td>
      <td><a href="/020_timers/005_timer_with_select">005_timer_with_select</a></td>
  </tr>
  <tr>
      <td rowspan="5">21</td>
      <td>Basic Ticker</td>
      <td>Demonstrates how to create and use a basic ticker in Go.</td>
      <td><a href="/021_tickers/01_basic_ticker">01_basic_ticker</a></td>
  </tr>
  <tr>
      <td>Stop Ticker</td>
      <td>Shows how to stop a ticker before it triggers the next tick.</td>
      <td><a href="/021_tickers/02_stop_ticker">02_stop_ticker</a></td>
  </tr>
  <tr>
      <td>Ticker with Select</td>
      <td>Illustrates how to use the select statement with a ticker for time-based control flow.</td>
      <td><a href="/021_tickers/03_ticker_with_select">03_ticker_with_select</a></td>
  </tr>
  <tr>
      <td>Reset Ticker</td>
      <td>Explains how to reset a ticker to its initial state.</td>
      <td><a href="/021_tickers/04_reset_ticker">04_reset_ticker</a></td>
  </tr>
  <tr>
      <td>Ticker with Limited Ticks</td>
      <td>Shows how to stop a ticker after a specific number of ticks.</td>
      <td><a href="/021_tickers/05_ticker_with_limited_ticks">05_ticker_with_limited_ticks</a></td>
  </tr>
  <tr>
    <td rowspan="5">22</td>
    <td>Basic Worker Pool</td>
    <td>Demonstrates how to implement a simple worker pool in Go.</td>
    <td><a href="/022_worker_pools/001_basic_worker_pool">001_basic_worker_pool</a></td>
  </tr>
  <tr>
    <td>Worker Pool with Buffered Channels</td>
    <td>Shows how to implement a worker pool with buffered channels.</td>
    <td><a href="/022_worker_pools/002_worker_pool_with_buffered_channels">002_worker_pool_with_buffered_channels</a></td>
  </tr>
  <tr>
    <td>Worker Pool with Error Handling</td>
    <td>Illustrates how to handle errors in a worker pool.</td>
    <td><a href="/022_worker_pools/003_worker_pool_with_error_handling">003_worker_pool_with_error_handling</a></td>
  </tr>
  <tr>
    <td>Dynamic Worker Pool</td>
    <td>Explains how to create a dynamic worker pool that adjusts based on workload.</td>
    <td><a href="/022_worker_pools/004_dynamic_worker_pool">004_dynamic_worker_pool</a></td>
  </tr>
  <tr>
    <td>Rate-Limited Worker Pool</td>
    <td>Shows how to limit the rate at which jobs are processed in a worker pool.</td>
    <td><a href="/022_worker_pools/005_rate_limited_worker_pool">005_rate_limited_worker_pool</a></td>
  </tr>
  <tr>
    <td rowspan="4">23</td>
    <td>Basic WaitGroup</td>
    <td>Demonstrates basic usage of `sync.WaitGroup` for synchronizing goroutines.</td>
    <td><a href="/023_waitgroup/001_basic_waitgroup">001_basic_waitgroup</a></td>
  </tr>
  <tr>
    <td>WaitGroup with Anonymous Functions</td>
    <td>Shows how to use `sync.WaitGroup` with anonymous functions in goroutines.</td>
    <td><a href="/023_waitgroup/002_waitGroup_with_anonymous_functions">002_waitGroup_with_anonymous_functions</a></td>
  </tr>
  <tr>
    <td>WaitGroup with Multiple Waits</td>
    <td>Illustrates using `sync.WaitGroup` to manage multiple stages of goroutines.</td>
    <td><a href="/023_waitgroup/003_waitgroup_with_multiple_waits">003_waitgroup_with_multiple_waits</a></td>
  </tr>
  <tr>
    <td>WaitGroup with Error Handling</td>
    <td>Demonstrates using `sync.WaitGroup` with error handling in workers.</td>
    <td><a href="/023_waitgroup/004_waitgroup_with_error_handling">004_waitgroup_with_error_handling</a></td>
  </tr>
  <tr>
    <td rowspan="5">24</td>
    <td>Basic Rate Limiter</td>
    <td>Demonstrates a basic implementation of a rate limiter.</td>
    <td><a href="/024_rate_limiter/001_basic_rate_limiter">001_basic_rate_limiter</a></td>
  </tr>
  <tr>
    <td>Rate Limiter with Burst Capacity</td>
    <td>Shows how to implement rate limiting with burst capacity.</td>
    <td><a href="/024_rate_limiter/002_rate_limiter_burst_capacity">002_rate_limiter_burst_capacity</a></td>
  </tr>
  <tr>
    <td>Custom Rate Limiter Using `time.After`</td>
    <td>Demonstrates custom rate limiting using `time.After` for more flexibility.</td>
    <td><a href="/024_rate_limiter/003_custom_rate_limiter_using_time_after">003_custom_rate_limiter_using_time_after</a></td>
  </tr>
  <tr>
    <td>Rate Limiter with `context.Context`</td>
    <td>Shows how to implement rate limiting using `context.Context` for request cancellation and control.</td>
    <td><a href="/024_rate_limiter/004_rate_limiter_with_context">004_rate_limiter_with_context</a></td>
  </tr>
  <tr>
    <td>Rate Limiter with `time.NewTicker`</td>
    <td>Demonstrates rate limiting using `time.NewTicker` for better control over the ticker's lifecycle.</td>
    <td><a href="/024_rate_limiter/005_rate_limiter_with_time_newticker">005_rate_limiter_with_time_newticker</a></td>
  </tr>
  <tr>
    <td rowspan="5">25</td>
    <td>Basic Atomic Counter Using `sync/atomic`</td>
    <td>Demonstrates a simple atomic counter using the `sync/atomic` package.</td>
    <td><a href="/025_atomic_counters/001_basic_atomic_counter_using_sync_atomic">001_basic_atomic_counter_using_sync_atomic</a></td>
  </tr>
  <tr>
    <td>Atomic Counter with Decrement and Compare-And-Swap</td>
    <td>Shows an atomic counter with decrement and compare-and-swap operations.</td>
    <td><a href="/025_atomic_counters/002_atomic_counter_with_decrement_and_compare_and_swap">002_atomic_counter_with_decrement_and_compare_and_swap</a></td>
  </tr>
  <tr>
    <td>Atomic Flag Using `sync/atomic`</td>
    <td>Demonstrates the use of an atomic flag for state management with `sync/atomic`.</td>
    <td><a href="/025_atomic_counters/003_atomic_flag_using_sync_atomic">003_atomic_flag_using_sync_atomic</a></td>
  </tr>
  <tr>
    <td>Atomic Counter with Load and Store Operations</td>
    <td>Shows how to use atomic load and store operations with counters.</td>
    <td><a href="/025_atomic_counters/004_atomic_counter_with_load_and_store_operations">004_atomic_counter_with_load_and_store_operations</a></td>
  </tr>
  <tr>
    <td>Atomic Pointer</td>
    <td>Demonstrates how to atomically load and store a pointer value.</td>
    <td><a href="/025_atomic_counters/005_atomic_pointer">005_atomic_pointer</a></td>
  </tr>
  <tr>
    <td rowspan="10">26</td>
    <td>Basic Sorting with Integers</td>
    <td>Demonstrates basic sorting operations with integer slices.</td>
    <td><a href="/026_sorting/001_basic_sorting_with_integers">001_basic_sorting_with_integers</a></td>
  </tr>
  <tr>
    <td>Sorting Strings</td>
    <td>Shows how to sort slices of strings.</td>
    <td><a href="/026_sorting/002_sorting_strings">002_sorting_strings</a></td>
  </tr>
  <tr>
    <td>Sorting by Custom Slice</td>
    <td>Demonstrates sorting a slice based on custom sorting logic.</td>
    <td><a href="/026_sorting/003_sorting_by_custom_slice">003_sorting_by_custom_slice</a></td>
  </tr>
  <tr>
    <td>Sorting by Multiple Fields</td>
    <td>Shows how to sort slices based on multiple fields.</td>
    <td><a href="/026_sorting/004_sorting_by_multiple_fields">004_sorting_by_multiple_fields</a></td>
  </tr>
  <tr>
    <td>Sorting a Custom Type Using `sort.Interface`</td>
    <td>Demonstrates sorting a custom type by implementing the `sort.Interface`.</td>
    <td><a href="/026_sorting/005_sorting_a_custom_type_using_sort_interface">005_sorting_a_custom_type_using_sort_interface</a></td>
  </tr>
  <tr>
    <td>Reverse Sorting</td>
    <td>Shows how to sort slices in reverse order.</td>
    <td><a href="/026_sorting/006_reverse_sorting">006_reverse_sorting</a></td>
  </tr>
  <tr>
    <td>Sorting with Custom Comparator Function</td>
    <td>Demonstrates sorting with a custom comparator function.</td>
    <td><a href="/026_sorting/007_sorting_with_custom_comparator_function">007_sorting_with_custom_comparator_function</a></td>
  </tr>
  <tr>
    <td>Sorting a Map by Keys</td>
    <td>Shows how to sort a map by its keys.</td>
    <td><a href="/026_sorting/008_sorting_a_map_by_keys">008_sorting_a_map_by_keys</a></td>
  </tr>
  <tr>
    <td>Sorting a Map by Values</td>
    <td>Demonstrates sorting a map by its values.</td>
    <td><a href="/026_sorting/009_sorting_a_map_by_values">009_sorting_a_map_by_values</a></td>
  </tr>
  <tr>
    <td>Concurrent Sorting with Goroutines</td>
    <td>Shows how to use goroutines to sort slices concurrently.</td>
    <td><a href="/026_sorting/010_concurrent_sorting_with_goroutines">010_concurrent_sorting_with_goroutines</a></td>
  </tr>
  <tr>
    <td rowspan="10">27</td>
    <td>Basic Example of Defer</td>
    <td>Demonstrates the basic usage of the `defer` keyword in Go.</td>
    <td><a href="/027_panic_and_defer/001_basic_example_of_defer">001_basic_example_of_defer</a></td>
  </tr>
  <tr>
    <td>Multiple Defer</td>
    <td>Shows how multiple `defer` statements work and their order of execution.</td>
    <td><a href="/027_panic_and_defer/002_multiple_defer">002_multiple_defer</a></td>
  </tr>
  <tr>
    <td>Defer with Function Call</td>
    <td>Demonstrates using `defer` with function calls.</td>
    <td><a href="/027_panic_and_defer/003_defer_with_function_call">003_defer_with_function_call</a></td>
  </tr>
  <tr>
    <td>Basic Example of Panic</td>
    <td>Shows a basic example of using the `panic` function in Go.</td>
    <td><a href="/027_panic_and_defer/004_basic_example_of_panic">004_basic_example_of_panic</a></td>
  </tr>
  <tr>
    <td>Panic with Recover</td>
    <td>Demonstrates how to use `recover` to handle panics and prevent program crashes.</td>
    <td><a href="/027_panic_and_defer/005_panic_with_recover">005_panic_with_recover</a></td>
  </tr>
  <tr>
    <td>Panic with Defer</td>
    <td>Shows how `defer` can be used alongside `panic` for resource cleanup.</td>
    <td><a href="/027_panic_and_defer/006_panic_with_defer">006_panic_with_defer</a></td>
  </tr>
  <tr>
    <td>Chain Panic and Defer</td>
    <td>Demonstrates chaining of `panic` and `defer` statements to understand their interactions.</td>
    <td><a href="/027_panic_and_defer/007_chain_panic_and_defer">007_chain_panic_and_defer</a></td>
  </tr>
  <tr>
    <td>Defer Log Functions Exit</td>
    <td>Shows how to use `defer` to log function exits and track function calls.</td>
    <td><a href="/027_panic_and_defer/008_defer_log_functions_exit">008_defer_log_functions_exit</a></td>
  </tr>
  <tr>
    <td>Closing Resources with Defer</td>
    <td>Demonstrates how to use `defer` to ensure resources like files are closed properly.</td>
    <td><a href="/027_panic_and_defer/009_closing_resources_with_defer">009_closing_resources_with_defer</a></td>
  </tr>
  <tr>
    <td>Defer with HTTPS Requests</td>
    <td>Shows how to use `defer` to handle HTTP response body cleanup after making requests.</td>
    <td><a href="/027_panic_and_defer/010_defer_with_https_requests">010_defer_with_https_requests</a></td>
  </tr>
  <tr>
    <td>28</td>
    <td>String Manipulations</td>
    <td>Demonstrates various string manipulation techniques in Go.</td>
    <td><a href="/028_string_manipulations/001_string_manipulations">001_string_manipulations</a></td>
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


