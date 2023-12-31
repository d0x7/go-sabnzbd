
# go-sabnzbd

Golang API bindings for SABnzbd

For full API documentation, see: http://wiki.sabnzbd.org/api

See `cmd/example/example.go` for a usage example

# Disclaimer

This is ancient code from an almost decade-old fork.  
I'm currently using it for a project and updating routes as I go, and need.


Both the history and queue mode/routes are up to spec of SABnzbd 3.7.1.  
For the rest I really got no clue; haven't used them yet.  
Some fields may be missing, others may exist but don't exist in SABnzbd anymore and therefore are never set.

# Installation

```bash
go get github.com/d0x7/go-sabnzbd/v2
```

# Example

```go
func main() {
  // Either create a simple client, using a host *with* port and an API key
  client, err := sabnzbd.SimpleClient("localhost:8080", "key")
  // Or create a secure client, which both always uses HTTPS and port 443
  client, err =  sabnzbd.SecureClient("sabnzbd.domain.tld", "key")
  // If you need a more advanced client, you can create one yourself, see the example in cmd/example/example.go for that

  // Check for errors, then call Auth to check if the API key is valid, and check again for potential errors
  _, err := client.Auth()
  // return value is not important; it's just the method of authentication used (apikey in this case)

  // You can check the remote server's version with Version()
  version, err := client.Version()

  // Get the current queue and whether it's paused, the current download speed, et cetera
  // SimpleQueue takes on parameters, whereas the normal Queue method takes paiging parameters (start, limit)
  queue, err := client.SimpleQueue()

  fmt.Printf("SpeedLimit is set to %v%%, aka. %s/s absolute\n", queue.SpeedLimitPercentage, queue.SpeedLimit)
  fmt.Printf("There is %s of %s free space left\n", queue.DownloadDiskFreeSpace, queue.DownloadDiskTotalSpace)
  fmt.Printf("Of the %d downloads, there are %s of %s downloaded already\n", queue.NoOfSlotsTotal, queue.BytesMissing, queue.Bytes)
  fmt.Printf("The remaining %s will be downloaded ", queue.BytesLeft)

  if queue.Paused {
      fmt.Printf("once the downloads are remained.\n")
  } else {
      fmt.Printf("in ETA %s\n", queue.TimeLeft)
  }

  //	=> SpeedLimit is set to 40%, aka. 5.03 MB/s absolute
  //	=> There is 15.86 TB of 18.50 TB free space left
  //	=> Of the 2 downloads, there are 779.35 MB of 9.55 GB downloaded already
  //	=> The remaining 8.77 GB will be downloaded in ETA 30m25s
}
```

Check out cmd/example/example.go to find a more advanced client creation and how to send NZBs to SABnzbd from Go.
