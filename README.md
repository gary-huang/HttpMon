# How to Run
- Compile by running `go build *.go`
- `./httpmon -stdin` for stdin input of log lines
- `./httpmon -file=${FILE_PATH}` for file input of log lines

# Proposed Structure of Program
The main function is in httpmon.go, a ring array buffer will be used to keep track of the traffic count for a section ("/api" for example) of a 10 second interval. It is meant to keep track of rolling window precision. The `Monitor` was supposed to keep track of the different sections encountered in the logs and have a map of section to its ring. The monitor would have depended on the ring of a URL section to add traffic counts to the ring for a URL section, stats would be returned at the end of a `ProcessLine` call on the `Monitor`. A go function would be called at the end of the `ProcessLine` call to scan through the sections and determine if an alert is needed to be raised or recovered. This can be inefficient since if theres a log line that comes in after 2 minutes of silence on a map of n sections, it can be an O(nm) time cost, where m is the size of the ring, to determine whether an alert is needed, since entries of the ring older than 2 minutes will need to be invalidated. 

# Optimizations
The scanner will consume 1 log line at a time, so the file is never entirely in memory for scalability. Obviously there needs to be concurrency control on each of the rings' traffic counts in order to determine precise alerting and recoveries needed, as well as traffic counting for a specfic section. To even better scale and optimize this, there can be a ring just used to keep track of total traffic for alerting purposes.

