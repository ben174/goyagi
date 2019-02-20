package signals

import (
    "os"
    "os/signal"
    "syscall"
)

// onlyOneSignalHandler is a channel that is closed the first time we invoke
// trapSignals(). If trapSignals() is ever invoked again the application will
// panic trying to close a channel that is already closed. This is to ensure
// we're only trapping the singals in one place.
var onlyOneSignalHandler = make(chan struct{})

// Setup registers for SIGTERM and SIGINT. A channel is returned
// which is closed on one of these signals. If a second signal is caught,
// the program is terminated with exit code 1.
func Setup() <-chan struct{} {
    close(onlyOneSignalHandler) // panics when called twice

    stop := make(chan struct{}) // create the channel we will return to the callee

    c := make(chan os.Signal, 2)                      // create channel c that will receive the os.Signals
    signal.Notify(c, syscall.SIGINT, syscall.SIGTERM) // send messages to c if we receive SIGINT or SIGTERM (can be triggered throughCtrl+C)

    go func() {
        <-c
        close(stop) // we've received a SIGINT message, we'll close this channel to notify the other goroutine that we can start shutting down our server
        <-c
        os.Exit(1) // we've received a SIGTERM, terminate the program regardless of current execution
    }()

    return stop
}